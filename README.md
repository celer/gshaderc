# Introduction

Gshaderc is a golang wrapper to for https://github.com/google/shaderc. It's a pretty straight forward mapping of the C API

The goal in providing this wrapper primarly is for allowing golang vulkan applications to compile shaders as need.

# Getting started

 * You'll need to install and compile https://github.com/google/shaderc *
 * go get github.com/celer/gsharderc

# Example

```go
	options := gs.NewCompilerOptions()
	defer options.Release()
	compiler := gs.NewCompiler()
	defer compiler.Release()
	data, err := ioutil.ReadFile(*input)

	if err != nil {
		panic(err)
	}

	result := compiler.CompileIntoSPV(string(data), shaderType, *input, *entryPoint, options)
	defer result.Release()

	if result.Error() == nil {
		err := ioutil.WriteFile("output", result.Bytes(), 0644)
		if err != nil {
			panic(err)
		}
	} else {
		panic(result.Error())
	}

```

See cmd/glslc.go for a basic example

# Foot notes

Tested against commit 0b9a2992c73d41debe4924d9f39260f773b5840a

