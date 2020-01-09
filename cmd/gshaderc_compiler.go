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
	"strings"

	gs "github.com/celer/gshaderc"
)

var target = flag.String("target", "vulkan_1_1", "specify compilation target (vulkan_1_0, vulkan_1_1, opengl, opengl_compat, webgpu)")
var input = flag.String("input", "", "input shader to compile")
var entryPoint = flag.String("entry-point", "main", "entry point to the shader")
var forSize = flag.Bool("optimize-size", false, "optimize for size")
var forPerf = flag.Bool("optimize-performance", false, "optimize for performance")

func main() {

	flag.Parse()

	if *input == "" {
		flag.PrintDefaults()
		os.Exit(-1)
	}

	options := gs.NewCompilerOptions()
	defer options.Release()
	compiler := gs.NewCompiler()
	defer compiler.Release()

	shaderType := gs.GetShaderTypeByFilename(*input)
	prefix := gs.GetShaderExtensionByType(shaderType)

	err := options.SetTargetByName(*target)
	if err != nil {
		log.Printf("error: %w", err)
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
		name := strings.Split(*input, ".")
		err := ioutil.WriteFile(prefix+name[0], result.Bytes(), 0644)
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
