//go:build (!android && !ios && !mobile) && ((darwin && arm64 && !gles) || (!gles && !arm && !arm64))
// +build darwin,arm64,!gles,!android,!ios,!mobile !gles,!arm,!arm64,!android,!ios,!mobile

package glfw

import (
	"runtime"

	"github.com/go-gl/glfw/v3.3/glfw"
)

func initWindowHints() {
	if runtime.GOOS == "darwin" {
		glfw.WindowHint(glfw.ContextVersionMajor, 3)
		glfw.WindowHint(glfw.ContextVersionMinor, 2)
		glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
		glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
		return
	}

	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 0)
}
