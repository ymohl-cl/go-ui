package text

import "github.com/veandco/go-sdl2/sdl"

func (T *Text) setSize(w, h int32) {
	T.rect.W = w
	T.rect.H = h

	if T.style.exist {
		T.style.rect.W = w
		T.style.rect.H = h
	}
}

func (T *Text) setUnderPosition(x, y int32) {
	if T.rect.X == T.style.rect.X-1 {
		T.style.rect.X = x - 1
	} else {
		T.style.rect.X = x + 1
	}

	if T.rect.Y == T.style.rect.Y-1 {
		T.style.rect.Y = y - 1
	} else {
		T.style.rect.Y = y + 1
	}
}

func (T *Text) cloneUnderStyle(source Styler) {
	T.style.color = sdl.Color{
		R: source.color.R,
		G: source.color.G,
		B: source.color.B,
		A: source.color.A,
	}

	T.style.rect = sdl.Rect{
		X: source.rect.X,
		Y: source.rect.Y,
	}
}
