package input

import (
	"errors"
	"strings"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/ymohl-cl/go-ui/objects"
	"github.com/ymohl-cl/go-ui/objects/block"
	"github.com/ymohl-cl/go-ui/objects/text"
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
	Txt   *text.Text
	block *block.Block

	// action click
	funcClick func(...interface{})
	dataClick []interface{}
}

/*
** Builder method
 */

// New create input object
func New(sizeText int, fontURL string, b *block.Block) (*Input, error) {
	var err error
	i := Input{status: objects.SBasic}

	x, y := b.GetPosition()
	w, h := b.GetSize()
	if i.Txt, err = text.New(caret, sizeText, fontURL, x+(w/2), y+(h/2)); err != nil {
		return nil, err
	}
	i.block = b

	return &i, nil
}

// Clone object and return a new
func (I Input) Clone(r *sdl.Renderer) (*Input, error) {
	var err error
	var prime *Input

	if prime.Txt, err = I.Txt.Clone(r); err != nil {
		return prime, err
	}
	if prime.block, err = I.block.Clone(r); err != nil {
		return prime, err
	}

	prime.status = I.status
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

// SetNewRune add a character on the text
func (I *Input) SetNewRune(key sdl.Keysym, r *sdl.Renderer) error {
	var s string
	var newStr string
	var err error

	if I.status == objects.SClick {
		s = I.Txt.GetTxt()

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
		if err = I.Txt.UpdateText(newStr, r); err != nil {
			panic(err)
		}
	}

	return nil
}

// Reset input with the caret only
func (I Input) Reset(r *sdl.Renderer) {
	var err error

	if err = I.Txt.UpdateText(caret, r); err != nil {
		panic(err)
	}
}

/*
** Getter method
 */

// GetTxt provide str on the actual input
func (I Input) GetTxt() string {
	s := I.Txt.GetTxt()

	idx := strings.Index(s, caret)
	return s[:idx] + s[idx+1:]
}

/*
** Updater method
 */
