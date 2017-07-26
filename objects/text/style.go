package text

import "github.com/veandco/go-sdl2/sdl"

const (
	// PositionTopLeft define position to the underStyle
	PositionTopLeft = 0
	// PositionTopRight define position to the underStyle
	PositionTopRight = 1
	// PositionBotLeft define position to the underStyle
	PositionBotLeft = 2
	// PositionBotRight define position to the underStyle
	PositionBotRight = 3
)

// Styler define personalisable style of text
type Styler struct {
	exist   bool
	color   sdl.Color
	rect    sdl.Rect
	texture *sdl.Texture
}

// SetUnderStyle to the object instance
func (T *Text) SetUnderStyle(red, green, blue, opacity uint8, position uint8) {
	T.style.exist = true
	T.style.color = sdl.Color{
		R: red,
		G: green,
		B: blue,
		A: opacity,
	}

	if position == PositionTopLeft || position == PositionBotRight {
		T.style.rect.Y = T.rect.Y - 1
	} else {
		T.style.rect.Y = T.rect.Y + 1
	}
	if position == PositionTopRight || position == PositionBotRight {
		T.style.rect.X = T.rect.X + 1
	} else {
		T.style.rect.X = T.rect.X - 1
	}
}

// SetUnderColor to the object instance
func (T *Text) SetUnderColor(red, green, blue, opacity uint8) {
	T.style.color = sdl.Color{
		R: red,
		G: green,
		B: blue,
		A: opacity,
	}
}

// GetUnderColor provide color of underStyle object
func (T Text) GetUnderColor() (r, g, b, a uint8) {
	return T.style.color.R, T.style.color.G, T.style.color.B, T.style.color.A
}
