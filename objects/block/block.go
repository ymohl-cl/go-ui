package block

import (
	"errors"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/ymohl-cl/go-ui/objects"
)

const (
	// Filled style
	Filled = 1
	// Border to draw only border
	Border = 2
)

// Block object implementation
type Block struct {
	// infos object
	status      uint8
	initialized bool

	// parameters objects
	rect   sdl.Rect
	colors map[uint8]sdl.Color

	// action click
	funcClick func(...interface{})
	dataClick []interface{}

	// style object
	style uint8
}

/*
** Builder method
 */

// New create block object
func New(styleBlock uint8) (*Block, error) {
	b := Block{status: objects.SFix}

	switch styleBlock {
	case Filled, Border:
		b.style = styleBlock
	default:
		return nil, errors.New(objects.ErrorStyle)
	}

	b.colors = make(map[uint8]sdl.Color)
	return &b, nil
}

// Clone object and return a new
func (B Block) Clone(r *sdl.Renderer) (*Block, error) {
	var err error
	var prime *Block

	if prime, err = New(B.style); err != nil {
		return prime, err
	}

	for s, c := range B.colors {
		if err = prime.SetVariantStyle(c.R, c.G, c.B, c.A, s); err != nil {
			return nil, err
		}
	}

	prime.UpdatePosition(B.rect.X, B.rect.Y)
	prime.UpdateSize(B.rect.W, B.rect.H)
	// Set functionClick
	prime.funcClick = B.funcClick
	prime.dataClick = B.dataClick
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

// SetVariantStyle define styles to interact with object.
func (B *Block) SetVariantStyle(red, green, blue, opacity uint8, status ...uint8) error {
	for _, s := range status {
		switch s {
		case objects.SFix, objects.SBasic, objects.SOver, objects.SClick:
			B.colors[s] = sdl.Color{
				R: red,
				G: green,
				B: blue,
				A: opacity,
			}
		default:
			return errors.New(objects.ErrorStatus)
		}
		if B.status == objects.SFix && s != objects.SFix {
			B.status = objects.SBasic
		}
	}
	return nil
}

/*
** Getter method
 */

/*
** Updater method
 */
