package main

import (
	// Game packages
	"github.com/plutial/Loop/window"
)

func main() {
	// Create a window
	window.Init(800, 450, "Game")
	window.Run()
}
