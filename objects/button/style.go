package button

import (
	"image"

	"github.com/veandco/go-sdl2/sdl"
)

// Styler define parameters of text and block (position and color) and image
type Styler struct {
	txtColor   sdl.Color
	blockColor sdl.Color
	img        *image.Image
}

func (S Styler) copy(dest *Styler, r *sdl.Renderer) error {
	var err error

	dest.txtColor = sdl.Color{
		R: S.txtColor.R,
		G: S.txtColor.G,
		B: S.txtColor.B,
		A: S.txtColor.A,
	}
	dest.txtRect = sdl.Rect{
		W: S.txtRect.W,
		H: S.txtRect.H,
		X: S.txtRect.X,
		Y: S.txtRect.Y,
	}

	dest.blockColor = sdl.Color{
		R: S.blockColor.R,
		G: S.blockColor.G,
		B: S.blockColor.B,
		A: S.blockColor.A,
	}
	dest.blockRect = sdl.Rect{
		W: S.blockRect.W,
		H: S.blockRect.H,
		X: S.blockRect.X,
		Y: S.blockRect.Y,
	}

	if dest.img != nil {
		if dest.img, err = S.img.Clone(r); err != nil {
			return err
		}
	}
	return nil
}

func (S *Styler) setColorTXT(r, g, b, a uint8) {
	S.txtColor = sdl.Color{
		R: r,
		G: g,
		B: b,
		A: a,
	}
}

func (S *Styler) setColorBlock(r, g, b, a uint8) {
	S.blockColor = sdl.Color{
		R: r,
		G: g,
		B: b,
		A: a,
	}
}

func (S *Styler) setPositionBlock(x, y int32) {
	S.blockRect.X = x
	S.blockRect.Y = y

	S.txtRect.X = S.blockRect.X + (S.blockRect.W / 2)
	S.txtRect.Y = S.blockRect.Y + (S.blockRect.H / 2)
	if S.img != nil {
		S.img.UpdatePosition(x, y)
	}
}

func (S *Styler) setSizeBlock(w, h int32) {
	S.blockRect.W = w
	S.blockRect.H = h

	S.txtRect.X = S.blockRect.X + (S.blockRect.W / 2)
	S.txtRect.Y = S.blockRect.Y + (S.blockRect.H / 2)
}

func (S *Styler) moveTo(x, y int32) {
	S.blockRect.X += x
	S.blockRect.Y += y

	S.txtRect.X = S.blockRect.X + (S.blockRect.W / 2)
	S.txtRect.Y = S.blockRect.Y + (S.blockRect.H / 2)
	if S.img != nil {
		S.img.MoveTo(x, y)
	}
}
