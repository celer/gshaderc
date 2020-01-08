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

// CommpilerOptions allows specific compiler options to be set
type CompilerOptions struct {
	options C.shaderc_compile_options_t
}

// NewCompilerOptions creates a new compiler options object
func NewCompilerOptions() *CompilerOptions {
	c := &CompilerOptions{}

	c.options = C.shaderc_compile_options_initialize()

	return c
}

// SetNanClamp
// Sets whether the compiler generates code for max and min builtins which,
// if given a NaN operand, will return the other operand. Similarly, the clamp
// builtin will favour the non-NaN operands, as if clamp were implemented
// as a composition of max and min.
func (c *CompilerOptions) SetNanClamp(enabled bool) {
	C.shaderc_compile_options_set_nan_clamp(c.options, C.bool(enabled))
}

// SetInvertY
// Sets whether the compiler should invert position.Y output in vertex shader.
func (c *CompilerOptions) SetInvertY(enabled bool) {
	C.shaderc_compile_options_set_invert_y(c.options, C.bool(enabled))
}

// SetBindingBase
// Sets the base binding number used for for a uniform resource type when
// automatically assigning bindings.  For GLSL compilation, sets the lowest
// automatically assigned number.  For HLSL compilation, the regsiter number
// assigned to the resource is added to this specified base.
func (c *CompilerOptions) SetBindingBase(kind UniformKind, base uint32) {
	C.shaderc_compile_options_set_binding_base(c.options, C.shaderc_uniform_kind(kind), C.uint(base))
}

// AddMacroDefinition
// Adds a predefined macro to the compilation options. This has the same
// effect as passing -Dname=value to the command-line compiler.  If value
// is NULL, it has the same effect as passing -Dname to the command-line
// compiler. If a macro definition with the same name has previously been
// added, the value is replaced with the new value. The macro name and
// value are passed in with char pointers, which point to their data, and
// the lengths of their data.
func (c *CompilerOptions) AddMacroDefinition(name, value string) {
	C.shaderc_compile_options_add_macro_definition(c.options, C.CString(name), C.ulong(len(name)), C.CString(value), C.ulong(len(value)))
}

// SetOptimizationLevel
// Sets the compiler optimization level to the given level. Only the last one
// takes effect if multiple calls of this function exist.
func (c *CompilerOptions) SetOptimizationLevel(level OptimizationLevel) {
	C.shaderc_compile_options_set_optimization_level(c.options, C.shaderc_optimization_level(level))
}

// SuppressWarnings
// Sets the compiler mode to suppress warnings, overriding warnings-as-errors
// mode. When both suppress-warnings and warnings-as-errors modes are
// turned on, warning messages will be inhibited, and will not be emitted
// as error messages.
func (c *CompilerOptions) SuppressWarnings() {
	C.shaderc_compile_options_set_suppress_warnings(c.options)
}

// Clone clones a copy of the compiler options
func (c *CompilerOptions) Clone() *CompilerOptions {
	n := &CompilerOptions{}
	n.options = C.shaderc_compile_options_clone(c.options)
	return n
}

// SetLimit sets a resource limit
func (c *CompilerOptions) SetLimit(limit ResourceLimit, value int) {
	C.shaderc_compile_options_set_limit(c.options, C.shaderc_limit(limit), C.int(value))
}

// SetTargetEnv
// Sets the target shader environment, affecting which warnings or errors will
// be issued.  The version will be for distinguishing between different versions
// of the target environment.  The version value should be either 0 or
// a value listed in shaderc_env_version.  The 0 value maps to Vulkan 1.0 if
// |target| is Vulkan, and it maps to OpenGL 4.5 if |target| is OpenGL.
func (c *CompilerOptions) SetTargetEnv(target Target, version EnvVersion) {
	C.shaderc_compile_options_set_target_env(c.options, C.shaderc_target_env(target), C.uint(version))
}

// SetSPIRVVersion
// Sets the target SPIR-V version. The generated module will use this version
// of SPIR-V.  Each target environment determines what versions of SPIR-V
// it can consume.  Defaults to the highest version of SPIR-V 1.0 which is
// required to be supported by the target environment.  E.g. Default to SPIR-V
// 1.0 for Vulkan 1.0 and SPIR-V 1.3 for Vulkan 1.1.
func (c *CompilerOptions) SetSPIRVVersion(version SPIRVVersion) {
	C.shaderc_compile_options_set_target_spirv(c.options, C.shaderc_spirv_version(version))
}

// SetWarningsAsErrors
// Sets the compiler mode to treat all warnings as errors. Note the
// suppress-warnings mode overrides this option, i.e. if both
// warning-as-errors and suppress-warnings modes are set, warnings will not
// be emitted as error messages.
func (c *CompilerOptions) SetWarningsAsErrors() {
	C.shaderc_compile_options_set_warnings_as_errors(c.options)
}

// SetAutoBindUniforms
// Sets whether the compiler should automatically assign bindings to uniforms
// that aren't already explicitly bound in the shader source.
func (c *CompilerOptions) SetAutoBindUniforms(auto bool) {
	C.shaderc_compile_options_set_auto_bind_uniforms(c.options, C.bool(auto))
}

// Releases the compiler options
func (c *CompilerOptions) Release() {
	C.shaderc_compile_options_release(c.options)
}
