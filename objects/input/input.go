package input

import (
	"errors"
	"strings"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/ymohl-cl/game-builder/objects"
	"github.com/ymohl-cl/game-builder/objects/block"
	"github.com/ymohl-cl/game-builder/objects/text"
)

const (
	caret = ")"

	// paddingSizeTxt is a secure space, count on character to keep space
	// on the right and on the left of word in input.
	paddingSizeTxt = 2
	errorSizeTxt   = "Text is too long"
)

// Input object implementation
type Input struct {
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
}

/*
** Builder method
 */

// New create input object, it's necessary to call SetParams before call Init
func New(sizeText int, fontURL string, styleBlock uint8) (*Input, error) {
	var err error
	i := Input{status: objects.SBasic}

	if i.txt, err = text.New(caret, sizeText, fontURL); err != nil {
		return nil, err
	}
	if i.block, err = block.New(styleBlock); err != nil {
		return nil, err
	}
	return &i, nil
}

// Clone object and return a new
func (I Input) Clone(r *sdl.Renderer) (*Input, error) {
	var err error
	var prime *Input

	if prime.txt, err = I.txt.Clone(r); err != nil {
		return prime, err
	}
	if prime.block, err = I.block.Clone(r); err != nil {
		return prime, err
	}

	I.fix.copy(&prime.fix)
	I.basic.copy(&prime.basic)
	I.over.copy(&prime.over)
	I.click.copy(&prime.click)

	if I.IsInit() {
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
func (I *Input) SetParams(x, y, w, h int32, r, g, b, a uint8) {
	I.fix.setPositionBlock(x, y)
	I.fix.setSizeBlock(w, h)
	I.fix.setColorBlock(r, g, b, a)

	I.basic.setPositionBlock(x, y)
	I.basic.setSizeBlock(w, h)
	I.basic.setColorBlock(r, g, b, a)

	I.over.setPositionBlock(x, y)
	I.over.setSizeBlock(w, h)
	I.over.setColorBlock(r, g, b, a)

	I.click.setPositionBlock(x, y)
	I.click.setSizeBlock(w, h)
	I.click.setColorBlock(r, g, b, a)
}

// SetColorTxtOnAll (status: /fix/basic/over/click)
func (I *Input) SetColorTxtOnAll(r, g, b, a uint8) {
	I.fix.setColorTXT(r, g, b, a)
	I.basic.setColorTXT(r, g, b, a)
	I.over.setColorTXT(r, g, b, a)
	I.click.setColorTXT(r, g, b, a)
}

// SetColorTxtOnFix (status: /fix)
func (I *Input) SetColorTxtOnFix(r, g, b, a uint8) {
	I.fix.setColorTXT(r, g, b, a)
}

// SetColorTxtOnBasic (status: /basic)
func (I *Input) SetColorTxtOnBasic(r, g, b, a uint8) {
	I.basic.setColorTXT(r, g, b, a)
}

// SetColorTxtOnOver (status: /Over)
func (I *Input) SetColorTxtOnOver(r, g, b, a uint8) {
	I.over.setColorTXT(r, g, b, a)
}

// SetColorTxtOnClick (status: /click)
func (I *Input) SetColorTxtOnClick(r, g, b, a uint8) {
	I.click.setColorTXT(r, g, b, a)
}

// SetColorBlockOnAll (status: /fix/basic/over/click)
func (I *Input) SetColorBlockOnAll(r, g, b, a uint8) {
	I.fix.setColorBlock(r, g, b, a)
	I.basic.setColorBlock(r, g, b, a)
	I.over.setColorBlock(r, g, b, a)
	I.click.setColorBlock(r, g, b, a)
}

// SetColorBlockOnFix (status: /fix)
func (I *Input) SetColorBlockOnFix(r, g, b, a uint8) {
	I.fix.setColorBlock(r, g, b, a)
}

// SetColorBlockOnBasic (status: /basic)
func (I *Input) SetColorBlockOnBasic(r, g, b, a uint8) {
	I.basic.setColorBlock(r, g, b, a)
}

// SetColorBlockOnOver (status: /over)
func (I *Input) SetColorBlockOnOver(r, g, b, a uint8) {
	I.over.setColorBlock(r, g, b, a)
}

// SetColorBlockOnClick (status: /click)
func (I *Input) SetColorBlockOnClick(r, g, b, a uint8) {
	I.click.setColorBlock(r, g, b, a)
}

// SetNewRune add a character on the text
func (I *Input) SetNewRune(key sdl.Keysym, r *sdl.Renderer) error {
	var s string
	var newStr string
	var err error

	if I.status == objects.SClick {
		s = I.txt.GetTxt()

		if (key.Scancode >= sdl.SCANCODE_A && key.Scancode <= sdl.SCANCODE_Z) || key.Scancode == sdl.SCANCODE_SPACE {
			newStr = I.addKeyCode(s, key.Sym)
		} else {
			switch key.Scancode {
			case sdl.SCANCODE_RIGHT:
				newStr = I.caretRight(s)
			case sdl.SCANCODE_LEFT:
				newStr = I.caretLeft(s)
			case sdl.SCANCODE_DOWN:
				newStr = I.caretEnd(s)
			case sdl.SCANCODE_UP:
				newStr = I.caretBegin(s)
			case sdl.SCANCODE_BACKSPACE:
				newStr = I.removeKeyBackspace(s)
			case sdl.SCANCODE_DELETE:
				newStr = I.removeKeyDelete(s)
			default:
				return nil // nothing todo
			}
		}

		if len(newStr) > len(s) && I.checkSizeTxt(newStr, s) == false {
			return errors.New(errorSizeTxt)
		}
		if err = I.txt.UpdateText(newStr, r); err != nil {
			panic(err)
		}
	}

	return nil
}

// Reset input with the caret only
func (I Input) Reset(r *sdl.Renderer) {
	var err error

	if err = I.txt.UpdateText(caret, r); err != nil {
		panic(err)
	}
}

/*
** Getter method
 */

// GetTxt provide str on the actual input
func (I Input) GetTxt() string {
	s := I.txt.GetTxt()

	idx := strings.Index(s, caret)
	return s[:idx] + s[idx+1:]
}

/*
** Updater method
 */
