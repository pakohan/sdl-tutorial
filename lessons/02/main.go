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

	sdl.WM_SetCaption("Hello World", "test.png")

	message := load_image("hello.bmp")
	background := load_image("background.bmp")

	apply_surface(0, 0, background, screen)
    apply_surface(320, 0, background, screen)
    apply_surface(0, 240, background, screen)
    apply_surface(320, 240, background, screen)

    apply_surface(180, 140, message, screen)

	message.Free()
	background.Free()

	screen.Flip()
	sdl.Delay(2000)

}

func load_image(filename string) *sdl.Surface {
	loadedImage := sdl.Load(filename)
	if loadedImage == nil {
		panic(sdl.GetError())
	}

	optimizedImage := loadedImage.DisplayFormat()
	loadedImage.Free()

	return optimizedImage
}

func apply_surface(x, y int16, src, dest *sdl.Surface) {
	var offset sdl.Rect

	offset.X = x
	offset.Y = y

	dest.Blit(&offset, src, nil)
}