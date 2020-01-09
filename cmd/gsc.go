// Copyright 2020 celer. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	gs "github.com/celer/gshaderc"
	"github.com/fsnotify/fsnotify"
)

var target = flag.String("target", gs.TargetVulkan11, "specify compilation target (vulkan_1_0, vulkan_1_1, opengl, opengl_compat, webgpu)")
var input = flag.String("input", "", "input shader to compile")
var entryPoint = flag.String("entry-point", "main", "entry point to the shader")
var forSize = flag.Bool("optimize-size", false, "optimize for size")
var forPerf = flag.Bool("optimize-performance", false, "optimize for performance")

type WatchDirs []string

func (w *WatchDirs) String() string {
	return "directory to watch for changes"
}

func (w *WatchDirs) Set(value string) error {
	*w = append(*w, value)
	return nil
}

var watchDirs WatchDirs

func main() {

	flag.Var(&watchDirs, "watch", "directory to watch for changes")

	flag.Parse()

	if len(watchDirs) > 0 {
		watcher, err := NewWatcher(watchDirs)
		if err != nil {
			log.Printf("%v", err)
			os.Exit(-7)
		}
		watcher.Run()
		os.Exit(0)
	}

	if *input == "" {
		flag.PrintDefaults()
		os.Exit(-1)
	}

	options := gs.NewCompilerOptions()
	defer options.Release()
	compiler := gs.NewCompiler()
	defer compiler.Release()

	shaderType := gs.GetShaderTypeByFilename(*input)

	err := options.SetTargetByName(*target)
	if err != nil {
		log.Printf("error: %v", err)
		os.Exit(-5)
	}

	log.Printf("Target: %s", *target)

	if *forSize {
		options.SetOptimizationLevel(gs.Size)
	} else if *forPerf {
		options.SetOptimizationLevel(gs.Performance)
	}

	data, err := ioutil.ReadFile(*input)
	if err != nil {
		log.Printf("error reading file: %v\n", err)
		os.Exit(-2)
	}

	result := compiler.CompileIntoSPV(string(data), shaderType, *input, *entryPoint, options)
	defer result.Release()

	if result.Error() == nil {
		err := ioutil.WriteFile(*input+".spv", result.Bytes(), 0644)
		if err != nil {
			log.Printf("error writing output: %v", err)
			os.Exit(-3)
		}
		os.Exit(0)
	} else {
		log.Printf("error compiling shader: %v\n", result.Error())
		fmt.Printf("%s\n", result.ErrorMessage())
		os.Exit(-4)
	}

}

type Watcher struct {
	watcher *fsnotify.Watcher
}

func NewWatcher(dirs []string) (*Watcher, error) {
	w := &Watcher{}

	var err error

	w.watcher, err = fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}
	for _, dir := range dirs {
		err := w.watcher.Add(dir)
		if err != nil {
			return nil, fmt.Errorf("error watching directory '%s': %w", dir, err)
		}
		log.Printf("watching directory %s for changes", dir)
	}
	return w, nil
}

func (w *Watcher) Run() {
	for {
		select {
		case event, ok := <-w.watcher.Events:
			if !ok {
				return
			}
			if event.Op&fsnotify.Write == fsnotify.Write {
				stype := gs.GetShaderTypeByFilename(event.Name)
				if stype != gs.InferFromSource {
					data, err := ioutil.ReadFile(event.Name)
					if err == nil {
						out, err := gs.CompileShader(string(data), event.Name, "")
						if err != nil {
							log.Printf("error compiling shader '%s': %v", event.Name, err)
						} else {
							err = ioutil.WriteFile(event.Name+".spv", out, 0644)
							if err != nil {
								log.Printf("error writing file '%s': %v", event.Name, err)
							} else {
								log.Printf("compiled %s -> %s", event.Name, event.Name+".spv")
							}
						}
					} else {
						log.Printf("error reading file '%s': %v", event.Name, err)
					}
				}
			}

		case err, ok := <-w.watcher.Errors:
			if !ok {
				log.Printf("error: %v", err)
				return
			}

		}
	}
	w.watcher.Close()
}
