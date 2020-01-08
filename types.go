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
package shaderc

// #cgo LDFLAGS: -lshaderc_combined -lstdc++
// #include <shaderc/shaderc.h>
import "C"
import ()

type ShaderType int

const (
	// Forced shader kinds. These shader kinds force the compiler to compile the
	// source code as the specified kind of shader.
	VertexShader         ShaderType = ShaderType(C.shaderc_vertex_shader)
	FragmentShader                  = ShaderType(C.shaderc_fragment_shader)
	ComputeShader                   = ShaderType(C.shaderc_compute_shader)
	GeometryShader                  = ShaderType(C.shaderc_geometry_shader)
	TessControlShader               = ShaderType(C.shaderc_tess_control_shader)
	TessEcaluationShader            = ShaderType(C.shaderc_tess_evaluation_shader)
	// Deduce the shader kind from #pragma annotation in the source code. Compiler
	// will emit error if #pragma annotation is not found.
	InferFromSource = ShaderType(C.shaderc_glsl_infer_from_source)
	// Default shader kinds. Compiler will fall back to compile the source code as
	// the specified kind of shader when #pragma annotation is not found in the
	// source code.
	DefaultVertexShader         = ShaderType(C.shaderc_glsl_default_vertex_shader)
	DefaultFragmentShader       = ShaderType(C.shaderc_glsl_default_fragment_shader)
	DefaultComputeShader        = ShaderType(C.shaderc_glsl_default_compute_shader)
	DefaultGeometryShader       = ShaderType(C.shaderc_glsl_default_geometry_shader)
	DefaultTessControlShader    = ShaderType(C.shaderc_glsl_default_tess_control_shader)
	DefaultTessEcaluationShader = ShaderType(C.shaderc_glsl_default_tess_evaluation_shader)

	SPIRVAssembly = ShaderType(C.shaderc_spirv_assembly)
)

type Target int

const (
	Vulkan Target = Target(C.shaderc_target_env_vulkan) // SPIR-V under Vulkan semantics
	OpenGL        = Target(C.shaderc_target_env_opengl) // SPIR-V under OpenGL semantics
	// NOTE: SPIR-V code generation is not supported for shaders under OpenGL
	// compatibility profile.
	OpenGLCompat = Target(C.shaderc_target_env_opengl_compat) // SPIR-V under OpenGL semantics,
	// including compatibility profile
	// functions
	WebGPU  = Target(C.shaderc_target_env_webgpu) // SPIR-V under WebGPU semantics
	Default = Vulkan
)

type EnvVersion int

const (
	Vulkan_1_0 EnvVersion = EnvVersion(C.shaderc_env_version_vulkan_1_0)
	Vulkan_1_1            = EnvVersion(C.shaderc_env_version_vulkan_1_1)
	OpenGL_4_5            = EnvVersion(C.shaderc_env_version_opengl_4_5)
	WebGPUEnv             = EnvVersion(C.shaderc_env_version_webgpu)
)

type SPIRVVersion int

const (
	SPIRV_1_0 SPIRVVersion = SPIRVVersion(C.shaderc_spirv_version_1_0)
	SPIRV_1_1 SPIRVVersion = SPIRVVersion(C.shaderc_spirv_version_1_1)
	SPIRV_1_2 SPIRVVersion = SPIRVVersion(C.shaderc_spirv_version_1_2)
	SPIRV_1_3 SPIRVVersion = SPIRVVersion(C.shaderc_spirv_version_1_3)
	SPIRV_1_4 SPIRVVersion = SPIRVVersion(C.shaderc_spirv_version_1_4)
	SPIRV_1_5 SPIRVVersion = SPIRVVersion(C.shaderc_spirv_version_1_5)
)

type OptimizationLevel int

const (
	Zero        OptimizationLevel = OptimizationLevel(C.shaderc_optimization_level_zero) // no optimization
	Size                          = OptimizationLevel(C.shaderc_optimization_level_size)
	Performance                   = OptimizationLevel(C.shaderc_optimization_level_performance)
)

// Uniform resource kinds.
// In Vulkan, uniform resources are bound to the pipeline via descriptors
// with numbered bindings and sets.
type UniformKind int

const (
	// Image and image buffer.
	UniformKindImage UniformKind = UniformKind(C.shaderc_uniform_kind_image)
	// Pure sampler.
	UniformKindSampler = UniformKind(C.shaderc_uniform_kind_sampler)
	// Sampled texture in GLSL, and Shader Resource View in HLSL.
	UniformKindTexture = UniformKind(C.shaderc_uniform_kind_texture)
	// Uniform Buffer Object (UBO) in GLSL.  Cbuffer in HLSL.
	UniformKindBuffer = UniformKind(C.shaderc_uniform_kind_buffer)
	// Shader Storage Buffer Object (SSBO) in GLSL.
	UniformKindStorageBuffer = UniformKind(C.shaderc_uniform_kind_storage_buffer)
	// Unordered Access View, in HLSL.  (Writable storage image or storage
	// buffer.)
	UniformKindUnorderedAccessView = UniformKind(C.shaderc_uniform_kind_unordered_access_view)
)

type ResourceLimit int

/*
	Some vim regex foo, incase this list ever needs to be regenerated

	s/\(shaderc_\(.*\)\),/\2=ResourceLimit(C.\1)/g
	s/^\s\+limit_max_\([a-z]\+\)/Max\u\1/g
	s/^\s\+limit_max_\([a-z]\+\)_\([a-z]\+\)_\([a-z]\+\)_\([a-z]\+\)_\([a-z]\+\)/Max\u\1\u\2\u\3\u\4\u\5/g
*/
const (
	MaxLights                                 ResourceLimit = ResourceLimit(C.shaderc_limit_max_lights)
	MaxClipPlanes                                           = ResourceLimit(C.shaderc_limit_max_clip_planes)
	MaxTextureUnits                                         = ResourceLimit(C.shaderc_limit_max_texture_units)
	MaxTextureCoords                                        = ResourceLimit(C.shaderc_limit_max_texture_coords)
	MaxVertexAttribs                                        = ResourceLimit(C.shaderc_limit_max_vertex_attribs)
	MaxVertexUniformComponents                              = ResourceLimit(C.shaderc_limit_max_vertex_uniform_components)
	MaxVaryingFloats                                        = ResourceLimit(C.shaderc_limit_max_varying_floats)
	MaxVertexTextureImageUnits                              = ResourceLimit(C.shaderc_limit_max_vertex_texture_image_units)
	MaxCombinedTextureImageUnits                            = ResourceLimit(C.shaderc_limit_max_combined_texture_image_units)
	MaxTextureImageUnits                                    = ResourceLimit(C.shaderc_limit_max_texture_image_units)
	MaxFragmentUniformComponents                            = ResourceLimit(C.shaderc_limit_max_fragment_uniform_components)
	MaxDrawBuffers                                          = ResourceLimit(C.shaderc_limit_max_draw_buffers)
	MaxVertexUniformVectors                                 = ResourceLimit(C.shaderc_limit_max_vertex_uniform_vectors)
	MaxVaryingVectors                                       = ResourceLimit(C.shaderc_limit_max_varying_vectors)
	MaxFragmentUniformVectors                               = ResourceLimit(C.shaderc_limit_max_fragment_uniform_vectors)
	MaxVertexOutputVectors                                  = ResourceLimit(C.shaderc_limit_max_vertex_output_vectors)
	MaxFragmentInputVectors                                 = ResourceLimit(C.shaderc_limit_max_fragment_input_vectors)
	MinProgramTexelOffset                                   = ResourceLimit(C.shaderc_limit_min_program_texel_offset)
	MaxProgramTexelOffset                                   = ResourceLimit(C.shaderc_limit_max_program_texel_offset)
	MaxClipDistances                                        = ResourceLimit(C.shaderc_limit_max_clip_distances)
	MaxComputeWorkGroupCountX                               = ResourceLimit(C.shaderc_limit_max_compute_work_group_count_x)
	MaxComputeWorkGroupCountY                               = ResourceLimit(C.shaderc_limit_max_compute_work_group_count_y)
	MaxComputeWorkGroupCountZ                               = ResourceLimit(C.shaderc_limit_max_compute_work_group_count_z)
	MaxComputeWorkGroupSizeX                                = ResourceLimit(C.shaderc_limit_max_compute_work_group_size_x)
	MaxComputeWorkGroupSizeY                                = ResourceLimit(C.shaderc_limit_max_compute_work_group_size_y)
	MaxComputeWorkGroupSizeZ                                = ResourceLimit(C.shaderc_limit_max_compute_work_group_size_z)
	MaxComputeUniformComponents                             = ResourceLimit(C.shaderc_limit_max_compute_uniform_components)
	MaxComputeTextureImageUnits                             = ResourceLimit(C.shaderc_limit_max_compute_texture_image_units)
	MaxComputeImageUniforms                                 = ResourceLimit(C.shaderc_limit_max_compute_image_uniforms)
	MaxComputeAtomicCounters                                = ResourceLimit(C.shaderc_limit_max_compute_atomic_counters)
	MaxComputeAtomicCounterBuffers                          = ResourceLimit(C.shaderc_limit_max_compute_atomic_counter_buffers)
	MaxVaryingComponents                                    = ResourceLimit(C.shaderc_limit_max_varying_components)
	MaxVertexOutputComponents                               = ResourceLimit(C.shaderc_limit_max_vertex_output_components)
	MaxGeometryInputComponents                              = ResourceLimit(C.shaderc_limit_max_geometry_input_components)
	MaxGeometryOutputComponents                             = ResourceLimit(C.shaderc_limit_max_geometry_output_components)
	MaxFragmentInputComponents                              = ResourceLimit(C.shaderc_limit_max_fragment_input_components)
	MaxImageUnits                                           = ResourceLimit(C.shaderc_limit_max_image_units)
	MaxCombinedImageUnitsAndFragment_outputs                = ResourceLimit(C.shaderc_limit_max_combined_image_units_and_fragment_outputs)
	MaxCombinedShaderOutputResources                        = ResourceLimit(C.shaderc_limit_max_combined_shader_output_resources)
	MaxImageSamples                                         = ResourceLimit(C.shaderc_limit_max_image_samples)
	MaxVertexImageUniforms                                  = ResourceLimit(C.shaderc_limit_max_vertex_image_uniforms)
	MaxTessControlImageUniforms                             = ResourceLimit(C.shaderc_limit_max_tess_control_image_uniforms)
	MaxTessEvaluationImageUniforms                          = ResourceLimit(C.shaderc_limit_max_tess_evaluation_image_uniforms)
	MaxGeometryImageUniforms                                = ResourceLimit(C.shaderc_limit_max_geometry_image_uniforms)
	MaxFragmentImageUniforms                                = ResourceLimit(C.shaderc_limit_max_fragment_image_uniforms)
	MaxCombinedImageUniforms                                = ResourceLimit(C.shaderc_limit_max_combined_image_uniforms)
	MaxGeometryTextureImageUnits                            = ResourceLimit(C.shaderc_limit_max_geometry_texture_image_units)
	MaxGeometryOutputVertices                               = ResourceLimit(C.shaderc_limit_max_geometry_output_vertices)
	MaxGeometryTotalOutputComponents                        = ResourceLimit(C.shaderc_limit_max_geometry_total_output_components)
	MaxGeometryUniformComponents                            = ResourceLimit(C.shaderc_limit_max_geometry_uniform_components)
	MaxGeometryVaryingComponents                            = ResourceLimit(C.shaderc_limit_max_geometry_varying_components)
	MaxTessControlInputComponents                           = ResourceLimit(C.shaderc_limit_max_tess_control_input_components)
	MaxTessControlOutputComponents                          = ResourceLimit(C.shaderc_limit_max_tess_control_output_components)
	MaxTessControlTextureImageUnits                         = ResourceLimit(C.shaderc_limit_max_tess_control_texture_image_units)
	MaxTessControlUniformComponents                         = ResourceLimit(C.shaderc_limit_max_tess_control_uniform_components)
	MaxTessControlTotalOutputComponents                     = ResourceLimit(C.shaderc_limit_max_tess_control_total_output_components)
	MaxTessEvaluationInputComponents                        = ResourceLimit(C.shaderc_limit_max_tess_evaluation_input_components)
	MaxTessEvaluationOutputComponents                       = ResourceLimit(C.shaderc_limit_max_tess_evaluation_output_components)
	MaxTessEvaluationTextureImageUnits                      = ResourceLimit(C.shaderc_limit_max_tess_evaluation_texture_image_units)
	MaxTessEvaluationUniformComponents                      = ResourceLimit(C.shaderc_limit_max_tess_evaluation_uniform_components)
	MaxTessPatchComponents                                  = ResourceLimit(C.shaderc_limit_max_tess_patch_components)
	MaxPatchVertices                                        = ResourceLimit(C.shaderc_limit_max_patch_vertices)
	MaxTessGenLevel                                         = ResourceLimit(C.shaderc_limit_max_tess_gen_level)
	MaxViewports                                            = ResourceLimit(C.shaderc_limit_max_viewports)
	MaxVertexAtomicCounters                                 = ResourceLimit(C.shaderc_limit_max_vertex_atomic_counters)
	MaxTessControlAtomicCounters                            = ResourceLimit(C.shaderc_limit_max_tess_control_atomic_counters)
	MaxTessEvaluationAtomicCounters                         = ResourceLimit(C.shaderc_limit_max_tess_evaluation_atomic_counters)
	MaxGeometryAtomicCounters                               = ResourceLimit(C.shaderc_limit_max_geometry_atomic_counters)
	MaxFragmentAtomicCounters                               = ResourceLimit(C.shaderc_limit_max_fragment_atomic_counters)
	MaxCombinedAtomicCounters                               = ResourceLimit(C.shaderc_limit_max_combined_atomic_counters)
	MaxAtomicCounterBindings                                = ResourceLimit(C.shaderc_limit_max_atomic_counter_bindings)
	MaxVertexAtomicCounterBuffers                           = ResourceLimit(C.shaderc_limit_max_vertex_atomic_counter_buffers)
	MaxTessControlAtomicCounterBuffers                      = ResourceLimit(C.shaderc_limit_max_tess_control_atomic_counter_buffers)
	MaxTessEvaluationAtomicCounterBuffers                   = ResourceLimit(C.shaderc_limit_max_tess_evaluation_atomic_counter_buffers)
	MaxGeometryAtomicCounterBuffers                         = ResourceLimit(C.shaderc_limit_max_geometry_atomic_counter_buffers)
	MaxFragmentAtomicCounterBuffers                         = ResourceLimit(C.shaderc_limit_max_fragment_atomic_counter_buffers)
	MaxCombinedAtomicCounterBuffers                         = ResourceLimit(C.shaderc_limit_max_combined_atomic_counter_buffers)
	MaxAtomicCounterBufferSize                              = ResourceLimit(C.shaderc_limit_max_atomic_counter_buffer_size)
	MaxTransformFeedbackBuffers                             = ResourceLimit(C.shaderc_limit_max_transform_feedback_buffers)
	MaxTransformFeedbackInterleavedComponents               = ResourceLimit(C.shaderc_limit_max_transform_feedback_interleaved_components)
	MaxCullDistances                                        = ResourceLimit(C.shaderc_limit_max_cull_distances)
	MaxCombinedClipAndCullDistances                         = ResourceLimit(C.shaderc_limit_max_combined_clip_and_cull_distances)
	MaxSamples                                              = ResourceLimit(C.shaderc_limit_max_samples)
)
