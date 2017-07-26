package block

import (
	"errors"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/ymohl-cl/game-builder/objects"
)

// Block object implementation
type Block struct {
	// infos object
	status      uint8
	initialized bool

	// parameters objects
	rect  sdl.Rect
	color sdl.Color

	// style object
	style Styler
}

/*
** Builder method
 */

// New create block object, it's necessary to call SetParams before call Init
func New(style uint8) (*Block, error) {
	b := Block{status: objects.SFix}

	switch style {
	case Filled:
		b.style.block = Filled
	case Border:
		b.style.block = Border
	default:
		return nil, errors.New(objects.ErrorStyle)
	}

	return &b, nil
}

// Clone object and return a new
func (B Block) Clone(r *sdl.Renderer) (*Block, error) {
	var err error
	var prime *Block

	if prime, err = New(B.style.block); err != nil {
		return prime, err
	}
	prime.SetParams(B.rect.X, B.rect.Y, B.rect.W, B.rect.H, B.color.R, B.color.G, B.color.B, B.color.A)

	if B.IsInit() {
		if err = prime.Init(r); err != nil {
			return nil, err
		}
	}
	return prime, nil
}

/*
** Setter method
 */

// SetParams define object's position and color
func (B *Block) SetParams(x, y, w, h int32, red, green, blue, opacity uint8) {
	B.rect = sdl.Rect{
		X: x,
		Y: y,
		W: w,
		H: h,
	}
	B.color = sdl.Color{
		R: red,
		G: green,
		B: blue,
		A: opacity,
	}
}

/*
** Getter method
 */

// GetColor provide color object
func (B Block) GetColor() (r, g, b, a uint8) {
	return B.color.R, B.color.G, B.color.B, B.color.A
}

/*
** Updater method
 */

// UpdateColor to change color of initialized object
func (B *Block) UpdateColor(red, green, blue, opacity uint8) {
	B.color = sdl.Color{
		R: red,
		G: green,
		B: blue,
		A: opacity,
	}
}
