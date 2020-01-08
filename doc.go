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

// Package gshaderc provides an API wrapper around https://github.com/google/shaderc
// the primary goal being to allow opengl and vulkan based go applications to
// have the ability to compile shaders into SPIRV (https://en.wikipedia.org/wiki/Standard_Portable_Intermediate_Representation)
//
// With the release of Vulkan ( https://www.khronos.org/vulkan/ ) the ability to compile
// shaders is no longer directly provided with the Vulkan API as it was previously provided
// as part of the OpenGL APIs.
//
package gshaderc
