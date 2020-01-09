package gshaderc

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	TargetVulkan11     string = "vulkan_1_1"
	TargetVulkan10            = "vulkan_1_0"
	TargetOpenGL              = "opengl"
	TargetOpenGLCompat        = "opengl_compat"
	TargetWebGPU              = "webgpu"
)

// CompileShader is a utility function for compiling a shader, it will
// determine the shader type based upon filename, and will default to vulkan_1_1
// if no target type is specified. It assumes a default entry point of "main"
func CompileShader(source, filename, target string) ([]byte, error) {

	options := NewCompilerOptions()
	defer options.Release()
	compiler := NewCompiler()
	defer compiler.Release()

	shaderType := GetShaderTypeByFilename(filename)
	if target != "" {
		err := options.SetTargetByName(target)
		if err != nil {
			return nil, err
		}
	} else {
		options.SetTargetByName(TargetVulkan11)
	}

	result := compiler.CompileIntoSPV(source, shaderType, filename, "main", options)
	defer result.Release()

	if result.Error() != nil {
		os.Stderr.Write([]byte(result.ErrorMessage()))
		return nil, result.Error()
	}
	return result.Bytes(), nil
}

func (c *CompilerOptions) SetTargetByName(target string) error {
	switch target {
	case TargetVulkan10:
		c.SetTargetEnv(Vulkan, Vulkan_1_0)
	case TargetVulkan11:
		c.SetTargetEnv(Vulkan, Vulkan_1_1)
	case TargetOpenGL:
		c.SetTargetEnv(OpenGL, OpenGL_4_5)
	case TargetOpenGLCompat:
		c.SetTargetEnv(OpenGLCompat, OpenGL_4_5)
	case TargetWebGPU:
		c.SetTargetEnv(WebGPU, WebGPUAll)
	default:
		return fmt.Errorf("unknown target: %s", target)
	}
	return nil
}

func GetShaderExtensionByType(stype ShaderType) string {
	switch stype {
	case FragmentShader:
		return "frag"
	case VertexShader:
		return "vert"
	case ComputeShader:
		return "comp"
	case GeometryShader:
		return "geom"
	case TessControlShader:
		return "tesc"
	case TessEvaluationShader:
		return "tese"
	}
	return ""
}

func GetShaderTypeByFilename(filename string) ShaderType {
	ext := filepath.Ext(filename)

	shaderType := InferFromSource

	switch ext {
	case ".frag":
		shaderType = FragmentShader
	case ".vert":
		shaderType = VertexShader
	case ".comp":
		shaderType = ComputeShader
	case ".geom":
		shaderType = GeometryShader
	case ".tesc":
		shaderType = TessControlShader
	case ".tese":
		shaderType = TessEvaluationShader
	}

	return shaderType

}
