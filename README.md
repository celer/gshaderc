[![GoDoc](https://godoc.org/github.com/celer/gshaderc?status.svg)](https://godoc.org/github.com/celer/gshaderc) [![Go Report Card](https://goreportcard.com/badge/github.com/celer/gshaderc)](https://goreportcard.com/report/github.com/celer/gshaderc)

# Introduction

Gshaderc is a golang wrapper to for https://github.com/google/shaderc. It's a pretty straight forward mapping of the C API

The goal in providing this wrapper primarily is for allowing golang Vulkan applications to compile shaders as need.

# Getting started

 * You'll need to install and compile https://github.com/google/shaderc [1]
 * go get -u github.com/celer/gshaderc


# Examples

Here is the simplest example:

```go
source := "#version 450\nvoid main() {}"
// This will assume you're targeting Vulkan, with an entry point of 'main' and infers the shader type based upon filename
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

# Tools

There cmd/gsc.go is a tool to either manually or automatically compile shaders based off of changes. The default output name is to 
append .spv to compile files, and it will look for the extensions .vert, .frag, .comp, .tesc, .geom and .tese and automatically compile
these files into shaders for the given target when they change.

```console
celer@bear:~/go/src/github.com/celer/vkg/examples/sdf$ gsc -watch shaders/
2020/01/08 18:35:23 watching directory shaders/ for changes
shaders/sdf.comp:336: error: '' :  syntax error, unexpected INT, expecting COMMA or SEMICOLON
2020/01/08 18:35:26 error compiling shader 'shaders/sdf.comp': compilation error
shaders/sdf.comp:336: error: '' :  syntax error, unexpected INT, expecting COMMA or SEMICOLON
2020/01/08 18:35:26 error compiling shader 'shaders/sdf.comp': compilation error
shaders/sdf.comp:336: error: '' :  syntax error, unexpected INT, expecting COMMA or SEMICOLON
2020/01/08 18:35:26 error compiling shader 'shaders/sdf.comp': compilation error
shaders/sdf.comp:341: error: '' :  syntax error, unexpected SEMICOLON, expecting LEFT_PAREN
2020/01/08 18:35:27 error compiling shader 'shaders/sdf.comp': compilation error
2020/01/08 18:35:27 compiled shaders/sdf.comp -> shaders/sdf.comp.spv
```

See cmd/gsc.go for a basic example

# Foot notes

[1] Tested against commit 0b9a2992c73d41debe4924d9f39260f773b5840a

