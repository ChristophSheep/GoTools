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
> go get -u github.com/go-gl/gl/v2.1/gl

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

func drawCube() {
	//gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
	gl.Scalef(0.1, 0.1, 0.1)
	//gl.Translatef(0, 0, -3.0)

	gl.Rotatef(30, 1, 1, 0)

	//gl.Rotatef(rotationY, 0, 1, 0)

	//rotationX += 0.5
	//rotationY += 0.5

	//gl.BindTexture(gl.TEXTURE_2D, texture)

	gl.Color4f(1, 1, 1, 1) // white

	gl.Color3ub(255, 255, 0) // yellow

	gl.Begin(gl.QUADS)

	gl.Normal3f(0, 0, 1)
	gl.TexCoord2f(0, 0)
	gl.Vertex3f(-1, -1, 1)
	gl.TexCoord2f(1, 0)
	gl.Vertex3f(1, -1, 1)
	gl.TexCoord2f(1, 1)
	gl.Vertex3f(1, 1, 1)
	gl.TexCoord2f(0, 1)
	gl.Vertex3f(-1, 1, 1)

	gl.Normal3f(0, 0, -1)
	gl.TexCoord2f(1, 0)
	gl.Vertex3f(-1, -1, -1)
	gl.TexCoord2f(1, 1)
	gl.Vertex3f(-1, 1, -1)
	gl.TexCoord2f(0, 1)
	gl.Vertex3f(1, 1, -1)
	gl.TexCoord2f(0, 0)
	gl.Vertex3f(1, -1, -1)

	gl.Normal3f(0, 1, 0)
	gl.TexCoord2f(0, 1)
	gl.Vertex3f(-1, 1, -1)
	gl.TexCoord2f(0, 0)
	gl.Vertex3f(-1, 1, 1)
	gl.TexCoord2f(1, 0)
	gl.Vertex3f(1, 1, 1)
	gl.TexCoord2f(1, 1)
	gl.Vertex3f(1, 1, -1)

	gl.Normal3f(0, -1, 0)
	gl.TexCoord2f(1, 1)
	gl.Vertex3f(-1, -1, -1)
	gl.TexCoord2f(0, 1)
	gl.Vertex3f(1, -1, -1)
	gl.TexCoord2f(0, 0)
	gl.Vertex3f(1, -1, 1)
	gl.TexCoord2f(1, 0)
	gl.Vertex3f(-1, -1, 1)

	gl.Normal3f(1, 0, 0)
	gl.TexCoord2f(1, 0)
	gl.Vertex3f(1, -1, -1)
	gl.TexCoord2f(1, 1)
	gl.Vertex3f(1, 1, -1)
	gl.TexCoord2f(0, 1)
	gl.Vertex3f(1, 1, 1)
	gl.TexCoord2f(0, 0)
	gl.Vertex3f(1, -1, 1)

	gl.Normal3f(-1, 0, 0)
	gl.TexCoord2f(0, 0)
	gl.Vertex3f(-1, -1, -1)
	gl.TexCoord2f(1, 0)
	gl.Vertex3f(-1, -1, 1)
	gl.TexCoord2f(1, 1)
	gl.Vertex3f(-1, 1, 1)
	gl.TexCoord2f(0, 1)
	gl.Vertex3f(-1, 1, -1)

	gl.End()
}

func drawLine(x1, y1, x2, y2 float32) {
	gl.Begin(gl.LINES)
	gl.Vertex2f(x1, y1)
	gl.Vertex2f(x2, y2)
	gl.End()
}

func clearScene() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func drawScene() {

	clearScene()

	gl.LineWidth(2)
	gl.Color3ub(255, 255, 0)

	y1 := float32(-1.0)
	step := float32(0.25)
	for y1 < 1 {
		drawLine(-1, y1, 1, 1)
		y1 += step
	}

	drawCube()
}

func main() {

	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	// Set Anti-Aliasing before window is created
	//
	glfw.WindowHint(glfw.Samples, 4)
	window, err := glfw.CreateWindow(640, 480, "Testing", nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()       // OK
	if err := gl.Init(); err != nil { // OK
		panic(err)
	}

	for !window.ShouldClose() {

		drawScene()

		window.SwapBuffers()
		glfw.PollEvents()
	}
}
