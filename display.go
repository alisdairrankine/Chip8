package chip8

const (
	PIXELS_MONOCHROME = iota
)

type Display interface {
	Draw(vram []byte, dataType int)
}
