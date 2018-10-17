/*

GLFW is a lightweight utility library for use with OpenGL. GLFW stands for Graphics Library Framework.
It provides programmers with the ability to create and manage windows and OpenGL contexts,
as well as handle joystick, keyboard and mouse input.

https://en.wikipedia.org/wiki/GLFW
https://www.glfw.org/
https://github.com/go-gl/glfw

Examples:

https://kylewbanks.com/blog/tutorial-opengl-with-golang-part-1-hello-opengl
https://github.com/sappharx/go-life/blob/master/main.go

https://github.com/raedatoui/learn-opengl-golang/blob/master/tutorial.go
https://github.com/raedatoui/learn-opengl-golang

Installation:

First get "glfw" with:

On Windows first install https://mingw-w64.org/
and choose "x86_64" during installation

> go get -u github.com/go-gl/glfw/v3.2/glfw

*/

package main

import (
	"log"
	"runtime" // OR: github.com/go-gl/gl/v2.1/gl

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

func initOpenGL() uint32 {
	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	prog := gl.CreateProgram()
	gl.LinkProgram(prog)
	return prog
}

func draw(window *glfw.Window, program uint32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)

	// Draw a line
	gl.Color3ub(255, 0, 0)
	gl.Begin(gl.LINES)
	gl.Vertex2f(10, 10)
	gl.Vertex2f(20, 20)
	gl.End()

	glfw.PollEvents()
	window.SwapBuffers()
}

func main() {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	window, err := glfw.CreateWindow(640, 480, "Testing", nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()

	program := initOpenGL()

	for !window.ShouldClose() {
		draw(window, program)
	}
}
