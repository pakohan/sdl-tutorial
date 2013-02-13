package main

import (
	"github.com/0xe2-0x9a-0x9b/Go-SDL/sdl"
)

func main() {
	if sdl.Init(sdl.INIT_EVERYTHING) != 0 {
		panic(sdl.GetError())
	}

	defer sdl.Quit()

	screen := sdl.SetVideoMode(640, 480, 32, sdl.SWSURFACE)
	if screen == nil {
		panic(sdl.GetError())
	}

}
