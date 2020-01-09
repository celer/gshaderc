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
// #include <shaderc/shaderc.h>
import "C"
import ()

type Compiler struct {
	compiler C.shaderc_compiler_t
}

// NewCompiler creates a new compiler
func NewCompiler() *Compiler {
	c := &Compiler{}

	c.compiler = C.shaderc_compiler_initialize()

	return c
}

// CompileIntoSPV
// Takes a GLSL source string and the associated shader kind, input file
// name, compiles it according to the given additional_options. If the shader
// kind is not set to a specified kind, but shaderc_glslc_infer_from_source,
// the compiler will try to deduce the shader kind from the source
// string and a failure in deducing will generate an error. Currently only
// #pragma annotation is supported. If the shader kind is set to one of the
// default shader kinds, the compiler will fall back to the default shader
// kind in case it failed to deduce the shader kind from source string.
// The input_file_name is a null-termintated string. It is used as a tag to
// identify the source string in cases like emitting error messages. It
// doesn't have to be a 'file name'.
// The source string will be compiled into SPIR-V binary and a
// shaderc_compilation_result will be returned to hold the results.
// The entry_point_name null-terminated string defines the name of the entry
// point to associate with this GLSL source. If the additional_options
// parameter is not null, then the compilation is modified by any options
// present.  May be safely called from multiple threads without explicit
// synchronization. If there was failure in allocating the compiler object,
// null will be returned.
func (c *Compiler) CompileIntoSPV(source string, shaderType ShaderType, inputFilename string, entryPoint string, options *CompilerOptions) *CompilationResult {
	cr := &CompilationResult{}
	cr.result = C.shaderc_compile_into_spv(c.compiler,
		C.CString(source),
		C.ulong(len(source)),
		C.shaderc_shader_kind(shaderType),
		C.CString(inputFilename),
		C.CString(entryPoint), options.options)
	return cr
}

// CompileIntoPreProcessedText
// Like shaderc_compile_into_spv, but the result contains SPIR-V assembly text
// instead of a SPIR-V binary module.  The SPIR-V assembly syntax is as defined
// by the SPIRV-Tools open source project.
func (c *Compiler) CompileIntoPreProcessedText(source string, shaderType ShaderType, inputFilename string, entryPoint string, options *CompilerOptions) *CompilationResult {
	cr := &CompilationResult{}
	cr.result = C.shaderc_compile_into_preprocessed_text(c.compiler,
		C.CString(source),
		C.ulong(len(source)),
		C.shaderc_shader_kind(shaderType),
		C.CString(inputFilename),
		C.CString(entryPoint), options.options)
	return cr
}

// AssembleIntoSPV
// Takes an assembly string of the format defined in the SPIRV-Tools project
// (https://github.com/KhronosGroup/SPIRV-Tools/blob/master/syntax.md),
// assembles it into SPIR-V binary and a shaderc_compilation_result will be
// returned to hold the results.
// The assembling will pick options suitable for assembling specified in the
// additional_options parameter.
// May be safely called from multiple threads without explicit synchronization.
// If there was failure in allocating the compiler object, null will be
// returned.
func (c *Compiler) AssembleIntoSPV(source string, options *CompilerOptions) *CompilationResult {
	cr := &CompilationResult{}
	cr.result = C.shaderc_assemble_into_spv(c.compiler,
		C.CString(source),
		C.ulong(len(source)),
		options.options)

	return cr
}

// Release the compiler instance
func (c *Compiler) Release() {
	C.shaderc_compiler_release(c.compiler)
}
