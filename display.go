package chip8

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	PIXELS_MONOCHROME = iota
)

type Display struct {
	window *sdl.Window
}

func NewDisplay() (*Display, error) {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		return nil, err
	}

	width := int32(64)
	height := int32(32)
	window, err := sdl.CreateWindow("chip8", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, width, height, sdl.WINDOW_SHOWN)
	if err != nil {
		return nil, err
	}

	display := &Display{
		window: window,
	}

	return display, nil
}

func (d *Display) Draw(vram []byte, dataType int) {
	switch dataType {
	case PIXELS_MONOCHROME:
		d.drawMonochrome(vram)
	}
}

func (d *Display) drawMonochrome(vram []byte) {
	r, err := d.window.GetRenderer()
	r.Clear()

	if err != nil {
		return
	}

	pixelData := decodeColourFromMonochromeBitmap(vram)
	for x := 0; x < 64; x++ {
		for y := 0; y < 32; y++ {
			colour := pixelData[(y*64)+(32*x)]
			r.SetDrawColor(colour.R, colour.G, colour.B, colour.A)
			r.DrawPoint(int32(x), int32(y))
		}
	}
	d.window.UpdateSurface()
}

func decodeColourFromMonochromeBitmap(vram []byte) []sdl.Color {

	pixels := make([]sdl.Color, 64*32)
	counter := 0
	for row := 0; row < 32; row++ {
		rowInt := vram[row]
		rowBinary := fmt.Sprintf("%064b", rowInt)

		for _, val := range []rune(rowBinary) {
			switch string(val) {
			case "0":
				pixels[counter] = sdl.Color{R: 0, G: 0, B: 0, A: 255}
				break
			case "1":
				pixels[counter] = sdl.Color{R: 255, G: 255, B: 255, A: 255}
				break
			}

			counter += 1
		}
	}

	return pixels
}

func (d *Display) Close() {

	d.window.Destroy()
	sdl.Quit()

}
