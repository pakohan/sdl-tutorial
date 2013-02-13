package main

import (
	"github.com/0xe2-0x9a-0x9b/Go-SDL/sdl"
)

func loadImage(name string) *sdl.Surface {
	image := sdl.Load(name)

	if image == nil {
		panic(sdl.GetError())
	}

	return image
}

func main() {
	if sdl.Init(sdl.INIT_EVERYTHING) != 0 {
		panic(sdl.GetError())
	}

	defer sdl.Quit()

	screen := sdl.SetVideoMode(640, 480, 32, sdl.SWSURFACE)
	if screen == nil {
		panic(sdl.GetError())
	}


	hello := sdl.Load("hello.bmp")

	screen.Blit(nil, hello, nil)
	screen.Flip()

	sdl.Delay(2000)
	hello.Free()
}