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
	"os"
	"path/filepath"
	"strings"

	gs "github.com/celer/gshaderc"
)

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

	ext := filepath.Ext(*input)

	shaderType := gs.InferFromSource

	prefix := ""

	switch ext {
	case ".frag":
		shaderType = gs.FragmentShader
		prefix = "frag."
	case ".vert":
		shaderType = gs.VertexShader
		prefix = "vert."
	case ".comp":
		shaderType = gs.ComputeShader
		prefix = "comp."
	case ".geom":
		shaderType = gs.GeometryShader
		prefix = "geom."
	}

	if *forSize {
		options.SetOptimizationLevel(gs.Size)
	} else if *forPerf {
		options.SetOptimizationLevel(gs.Performance)
	}

	data, err := ioutil.ReadFile(*input)
	if err != nil {
		fmt.Printf("error reading file: %v\n", err)
		os.Exit(-2)
	}

	result := compiler.CompileIntoSPV(string(data), shaderType, *input, *entryPoint, options)
	defer result.Release()

	if result.Error() == nil {
		name := strings.Split(*input, ".")
		err := ioutil.WriteFile(prefix+name[0], result.Bytes(), 0644)
		if err != nil {
			fmt.Printf("error writing output: %v", err)
			os.Exit(-3)
		}
	} else {
		fmt.Printf("error compiling shader: %v\n", result.Error())
		fmt.Printf("error compiling shader: %s\n", result.ErrorMessage())
	}

}
