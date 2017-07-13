package button

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/ymohl-cl/game-builder/objects/block"
	"github.com/ymohl-cl/game-builder/objects/image"
	"github.com/ymohl-cl/game-builder/objects/text"
)

// Button object implementation
type Button struct {
	// infos object
	initialized bool

	// content object
	block *block.Block
	img   *image.Image
	txt   *text.Text

	// action click
	funcClick func(...interface{})
	dataClick []interface{}
}

/*
** Builder method
 */

// New create button object
func New(bl *block.Block, i *image.Image, t *text.Text) *Button {
	b := Button{block: bl, img: i, txt: t}

	return &b
}

// Clone object and return a new
func (B Button) Clone(r *sdl.Renderer) (*Button, error) {
	var err error
	var prime *Button
	var t *text.Text
	var bl *block.Block
	var i *image.Image

	if B.block != nil {
		if bl, err = B.block.Clone(r); err != nil {
			return nil, err
		}
	}
	if B.img != nil {
		if i, err = B.img.Clone(r); err != nil {
			return nil, err
		}
	}
	if B.txt != nil {
		if t, err = B.txt.Clone(r); err != nil {
			return nil, err
		}
	}
	prime = New(bl, i, t)

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

/*
** Getter method
 */

/*
** Updater method
 */
