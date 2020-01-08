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

package gshaderc

import (
	"fmt"
	"testing"
)

func TestCompilerInclude(t *testing.T) {
	options := NewCompilerOptions()
	compiler := NewCompiler()

	includeRequests := 0
	options.SetIncludeCallback(func(requestedSource string, iType IncludeType, requestingSource string, includeDepth int) (sourceName, content string, err error) {
		includeRequests++
		fmt.Printf("Got an include request %s %d %s %d\n", requestedSource, iType, requestingSource, includeDepth)

		if requestedSource != "foo" {
			t.Fatal("Expected include request for 'foo'")
		}

		if requestingSource != "main.vert" {
			t.Fatal("Expected include request from 'main.vert'")
		}

		return requestedSource, "#version 450\n", nil
	})

	goodSource := "#version 450\n#include <foo>\nvoid main() {}"
	res := compiler.CompileIntoSPV(goodSource, VertexShader, "main.vert", "main", options)

	if includeRequests != 1 {
		t.Fatal("Expected 1 inclusion request")
	}

	res.Release()

	defer compiler.Release()
	defer options.Release()

}

func TestCompiler(t *testing.T) {
	options := NewCompilerOptions()
	compiler := NewCompiler()

	badSource := "void main(){}"
	res := compiler.CompileIntoSPV(badSource, VertexShader, "main.vert", "main", options)

	if res.Error() != CompilationError {
		t.Fatal("Expected compilation error")
	}
	if res.ErrorMessage() == "" {
		t.Fatal("Expected a specific error message")
	}
	res.Release()

	goodSource := "#version 450\nvoid main() {}"
	res = compiler.CompileIntoSPV(goodSource, VertexShader, "main.vert", "main", options)

	if res.Error() != nil {
		t.Fatal("Didn't expect a compilation error")
	}
	if res.ErrorMessage() != "" {
		t.Fatal("Didn't expect a specific error message")
	}

	if len(res.Bytes()) == 0 {
		t.Fatal("Expected a binary result")
	}

	res.Release()

	defer compiler.Release()
	defer options.Release()

}
