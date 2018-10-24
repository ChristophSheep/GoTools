package controller

import "github.com/go-gl/glfw/v3.0/glfw"

func keyCallback(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {

	// ZOOM - in, out

	//
	if mods == glfw.ModShift {
		if key == glfw.KeyUp {
			scene.camera.scale *= 0.5
		}
		if key == glfw.KeyDown {
			scene.camera.scale /= 0.5
		}
	}

	// MOVE CAMERA - left, right, up, down
	//
	if key == glfw.KeyUp {
		scene.camera.position[1] += 1.0
	}
	if key == glfw.KeyDown {
		scene.camera.position[1] -= 1.0
	}
	if key == glfw.KeyLeft {
		scene.camera.position[0] -= 1.0
	}
	if key == glfw.KeyRight {
		scene.camera.position[0] += 1.0
	}

}
