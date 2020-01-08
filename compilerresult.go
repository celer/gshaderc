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

// #cgo LDFLAGS: -lshaderc_combined -lstdc++
// #include <shaderc/shaderc.h>
import "C"
import (
	"fmt"
	"unsafe"
)

// CompilationResult the result of compiling stuff
type CompilationResult struct {
	result C.shaderc_compilation_result_t
}

func compilationStatusToError(status C.shaderc_compilation_status) error {
	switch status {
	case C.shaderc_compilation_status_success:
		return nil
	case C.shaderc_compilation_status_invalid_stage:
		return InvalidStageError
	case C.shaderc_compilation_status_compilation_error:
		return CompilationError
	case C.shaderc_compilation_status_internal_error:
		return InternalError
	case C.shaderc_compilation_status_null_result_object:
		return NullResultObjectError
	case C.shaderc_compilation_status_invalid_assembly:
		return InvalidAssemblyError
	case C.shaderc_compilation_status_validation_error:
		return ValidationError
	case C.shaderc_compilation_status_transformation_error:
		return TransformationError
	case C.shaderc_compilation_status_configuration_error:
		return ConfigurationError
	}
	return fmt.Errorf("unknown error %v", status)
}

// ErrorMessage returns a specific error message
func (c *CompilationResult) ErrorMessage() string {
	em := C.shaderc_result_get_error_message(c.result)
	return C.GoString(em)
}

// Error returns a generic error message
func (c *CompilationResult) Error() error {
	status := C.shaderc_result_get_compilation_status(c.result)
	return compilationStatusToError(status)
}

// NumErrors returns the number of errors
func (c *CompilationResult) NumErrors() int {
	return int(C.shaderc_result_get_num_errors(c.result))
}

// NumWarnings returns the number of warnings
func (c *CompilationResult) NumWarnings() int {
	return int(C.shaderc_result_get_num_warnings(c.result))
}

// Bytes returns the resulting compiled item
func (c *CompilationResult) Bytes() []byte {
	dataLen := C.shaderc_result_get_length(c.result)
	data := C.shaderc_result_get_bytes(c.result)

	b := C.GoBytes(unsafe.Pointer(data), C.int(dataLen))
	return b
}

// Release releases the compilation results
func (c *CompilationResult) Release() {
	C.shaderc_result_release(c.result)
}
