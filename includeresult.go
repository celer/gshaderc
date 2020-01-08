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

// #cgo LDFLAGS: -lshaderc_combined -lstdc++ -lm
/*
#include <shaderc/shaderc.h>
#include <stdlib.h>

static shaderc_include_result* new_shader_include_result(){
	return (shaderc_include_result*) malloc(sizeof(shaderc_include_result));
}

static void free_shader_include_result(void *user_data, shaderc_include_result *res){
	free(res);
}

shaderc_include_result* cbIncludeResolver(void* user_data, char* requested_source, int type,
    char* requesting_source, size_t include_depth);


static void _register_callback(shaderc_compile_options_t options,void *userData){
	shaderc_compile_options_set_include_callbacks(options, (shaderc_include_resolve_fn) cbIncludeResolver, free_shader_include_result, userData);
}


*/
import "C"
import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"unsafe"

	ptr "github.com/mattn/go-pointer"
)

type callback struct {
	resolver IncludeResolver
}

//export cbIncludeResolver
func cbIncludeResolver(userData unsafe.Pointer, requestedSource *C.char, itype C.int, requestingSource *C.char, includeDepth C.ulong) *C.shaderc_include_result {

	rds := C.GoString(requestedSource)
	ris := C.GoString(requestingSource)

	callback := ptr.Restore(userData).(*callback)

	sourceName, content, err := callback.resolver(rds, IncludeType(itype), ris, int(includeDepth))

	result := C.new_shader_include_result()

	if err == nil {
		result.source_name = C.CString(sourceName)
		result.source_name_length = C.ulong(len(sourceName))

		result.content = C.CString(content)
		result.content_length = C.ulong(len(content))

	} else {
		msg := err.Error()
		result.content = C.CString(msg)
		result.content_length = C.ulong(len(msg))
	}

	return result
}

type IncludeType int

const (
	// IncludeRelative E.g. #include "source"
	IncludeRelative IncludeType = IncludeType(C.shaderc_include_type_relative)
	// IncludeStandard E.g. #include <source>
	IncludeStandard = IncludeType(C.shaderc_include_type_standard)
)

// CreateDefaultIncludeResolver returns a basic include resolover which looks for files in a list of specified directories
func CreateDefaultIncludeResolver(dirs []string) IncludeResolver {
	return func(requestedSource string, itype IncludeType, requestingSource string, includeDepth int) (sourceName, content string, err error) {
		if itype == IncludeRelative {
			p := filepath.Join(".", requestedSource)
			absp, err := filepath.Abs(p)
			if err != nil {
				return "", "", fmt.Errorf("error opening file '%s': %w", requestedSource, err)
			}
			data, err := ioutil.ReadFile(absp)
			if err == nil {
				return absp, string(data), nil
			} else {
				return "", "", fmt.Errorf("error opening file '%s': %w", requestedSource, err)
			}
		} else {

			for _, d := range dirs {
				p := filepath.Join(d, requestedSource)
				absp, err := filepath.Abs(p)
				if err != nil {
					continue
				}
				data, err := ioutil.ReadFile(absp)
				if err == nil {
					return absp, string(data), nil
				}
			}
			return "", "", fmt.Errorf("unable to find file '%s'", requestedSource)
		}
	}
}

// IncludeResolver
// An includer resolver type for mapping an #include request to an include
// result. The requested_source parameter specifies the name of the source being
// requested. The type parameter specifies the kind of inclusion request being made.
// The requesting_source parameter specifies the name of the source containing
// the #include request. Returns the name of the source file, the contents and optionally
// an error
type IncludeResolver func(requestedSource string, itype IncludeType, requestingSource string, includeDepth int) (sourceName, content string, err error)

// SetIncludeCallback sets a include resolver
func (c *CompilerOptions) SetIncludeCallback(resolver IncludeResolver) {
	C._register_callback(c.options, ptr.Save(&callback{resolver: resolver}))
}
