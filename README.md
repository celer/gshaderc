# Introduction

Gshaderc is a golang wrapper to for https://github.com/google/shaderc. It's a pretty straight forward mapping of the C API

The goal in providing this wrapper primarly is for allowing golang vulkan applications to compile shaders as need.

# Getting started

 * You'll need to install and compile https://github.com/google/shaderc [1]
 * go get github.com/celer/gsharderc

# Examples

Here is the simplist example:

```go
	source := "#version 450\nvoid main() {}"
	// This will assume you're targeting vulkan, with an entry point of 'main' and infers the shader type based upon filename
	data, err := CompileShader(source, "main.vert", "")

```

Here is a more complex example:

```go
	options := gs.NewCompilerOptions()
	defer options.Release()
	compiler := gs.NewCompiler()
	defer compiler.Release()
	data, err := ioutil.ReadFile(*input)

	if err != nil {
		panic(err)
	}

	options.SetOptimizationLevel(gs.Performance)

	result := compiler.CompileIntoSPV(string(data), shaderType, filename, entryPoint, options)
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

See cmd/gshaderc_compiler.go for a basic example

# Foot notes

[1] Tested against commit 0b9a2992c73d41debe4924d9f39260f773b5840a

