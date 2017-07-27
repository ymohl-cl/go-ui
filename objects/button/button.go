package button

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/ymohl-cl/game-builder/objects"
	"github.com/ymohl-cl/game-builder/objects/block"
	"github.com/ymohl-cl/game-builder/objects/image"
	"github.com/ymohl-cl/game-builder/objects/text"
)

// Button object implementation
type Button struct {
	// infos object
	status      uint8
	initialized bool

	// content object
	txt   *text.Text
	block *block.Block

	// style object
	fix   Styler
	basic Styler
	over  Styler
	click Styler

	// action click
	funcClick func(...interface{})
	dataClick []interface{}
}

/*
** Builder method
 */

// New create button object, it's necessary to call SetParams before call Init
// Call set action to define action on click
func New(txt, fontURL string, sizeTXT int, styleBlock uint8) (*Button, error) {
	var err error
	b := Button{status: objects.SBasic}

	if b.txt, err = text.New(txt, sizeTXT, fontURL); err != nil {
		return nil, err
	}
	if b.block, err = block.New(styleBlock); err != nil {
		return nil, err
	}
	return &b, nil
}

// Clone object and return a new
func (B Button) Clone(r *sdl.Renderer) (*Button, error) {
	var err error
	var prime *Button

	if prime.txt, err = B.txt.Clone(r); err != nil {
		return prime, err
	}
	if prime.block, err = B.block.Clone(r); err != nil {
		return prime, err
	}

	B.fix.copy(&prime.fix, r)
	B.basic.copy(&prime.basic, r)
	B.over.copy(&prime.over, r)
	B.click.copy(&prime.click, r)

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
func (B *Button) SetParams(x, y, w, h int32, r, g, b, a uint8) {
	B.fix.setPositionBlock(x, y)
	B.fix.setSizeBlock(w, h)
	B.fix.setColorBlock(r, g, b, a)

	B.basic.setPositionBlock(x, y)
	B.basic.setSizeBlock(w, h)
	B.basic.setColorBlock(r, g, b, a)

	B.over.setPositionBlock(x, y)
	B.over.setSizeBlock(w, h)
	B.over.setColorBlock(r, g, b, a)

	B.click.setPositionBlock(x, y)
	B.click.setSizeBlock(w, h)
	B.click.setColorBlock(r, g, b, a)
}

// SetAction to get it on click button
func (B *Button) SetAction(f func(...interface{}), d ...interface{}) {
	B.funcClick = f

	for _, v := range d {
		B.dataClick = append(B.dataClick, v)
	}
}

// SetColorTxtOnAll (status: /fix/basic/over/click)
func (B *Button) SetColorTxtOnAll(r, g, b, a uint8) {
	B.fix.setColorTXT(r, g, b, a)
	B.basic.setColorTXT(r, g, b, a)
	B.over.setColorTXT(r, g, b, a)
	B.click.setColorTXT(r, g, b, a)
}

// SetColorTxtOnFix (status: /fix)
func (B *Button) SetColorTxtOnFix(r, g, b, a uint8) {
	B.fix.setColorTXT(r, g, b, a)
}

// SetColorTxtOnBasic (status: /basic)
func (B *Button) SetColorTxtOnBasic(r, g, b, a uint8) {
	B.basic.setColorTXT(r, g, b, a)
}

// SetColorTxtOnOver (status: /Over)
func (B *Button) SetColorTxtOnOver(r, g, b, a uint8) {
	B.over.setColorTXT(r, g, b, a)
}

// SetColorTxtOnClick (status: /click)
func (B *Button) SetColorTxtOnClick(r, g, b, a uint8) {
	B.click.setColorTXT(r, g, b, a)
}

// SetColorBlockOnAll (status: /fix/basic/over/click)
func (B *Button) SetColorBlockOnAll(r, g, b, a uint8) {
	B.fix.setColorBlock(r, g, b, a)
	B.basic.setColorBlock(r, g, b, a)
	B.over.setColorBlock(r, g, b, a)
	B.click.setColorBlock(r, g, b, a)
}

// SetColorBlockOnFix (status: /fix)
func (B *Button) SetColorBlockOnFix(r, g, b, a uint8) {
	B.fix.setColorBlock(r, g, b, a)
}

// SetColorBlockOnBasic (status: /basic)
func (B *Button) SetColorBlockOnBasic(r, g, b, a uint8) {
	B.basic.setColorBlock(r, g, b, a)
}

// SetColorBlockOnOver (status: /over)
func (B *Button) SetColorBlockOnOver(r, g, b, a uint8) {
	B.over.setColorBlock(r, g, b, a)
}

// SetColorBlockOnClick (status: /click)
func (B *Button) SetColorBlockOnClick(r, g, b, a uint8) {
	B.click.setColorBlock(r, g, b, a)
}

// SetImageOnAll (status: /fix/basic/over/click)
func (B *Button) SetImageOnAll(url string) error {
	var err error

	if B.fix.img, err = image.New(url); err != nil {
		return err
	}
	if B.basic.img, err = image.New(url); err != nil {
		return err
	}
	if B.over.img, err = image.New(url); err != nil {
		return err
	}
	if B.click.img, err = image.New(url); err != nil {
		return err
	}
	return nil
}

// SetImageOnFix (status: /fix)
func (B *Button) SetImageOnFix(url string) error {
	var err error

	if B.fix.img, err = image.New(url); err != nil {
		return err
	}
	return nil
}

// SetImageOnBasic (status: /basic)
func (B *Button) SetImageOnBasic(url string) error {
	var err error

	if B.basic.img, err = image.New(url); err != nil {
		return err
	}
	return nil
}

// SetImageOnOver (status: /over)
func (B *Button) SetImageOnOver(url string) error {
	var err error

	if B.over.img, err = image.New(url); err != nil {
		return err
	}
	return nil
}

// SetImageOnClick (status: /click)
func (B *Button) SetImageOnClick(url string) error {
	var err error

	if B.click.img, err = image.New(url); err != nil {
		return err
	}
	return nil
}

/*
** Getter method
 */

/*
** Updater method
 */
