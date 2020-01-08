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
import (
	"fmt"
)

var InvalidStageError = fmt.Errorf("invalid stage")
var CompilationError = fmt.Errorf("compilation error")
var InternalError = fmt.Errorf("internal error")
var NullResultObjectError = fmt.Errorf("null result object error")
var InvalidAssemblyError = fmt.Errorf("invalid assembly error")
var ValidationError = fmt.Errorf("validation error")
var TransformationError = fmt.Errorf("transformation error")
var ConfigurationError = fmt.Errorf("configuration error")
