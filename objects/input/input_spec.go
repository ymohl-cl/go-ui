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
	caret        = ")"
	errorSizeTxt = "Text is too long"
)

// New create a new input object
func New(status uint8, sizeTXT int, fontURL string) (*Input, error) {
	var err error

	i := new(Input)

	i.status = status
	if i.txt, err = text.New(caret, sizeTXT, fontURL); err != nil {
		return nil, err
	}
	return i, nil
}

func (I *Input) SetColorTxt(red, green, blue, opacity uint8) error {
	var err error

	c := new(objects.Color)
	c.SetColor(red, green, blue, opacity)
	I.txt.SetColor(c)
	return err
}

func (I *Input) SetParams(bf, bb, bo, bc *block.Block) {
	I.bFix = bf
	I.bBasic = bb
	I.bOver = bo
	I.bClick = bc
}

func (I *Input) SetBlockFix(b *block.Block) {
	I.bFix = b
}

func (I *Input) SetBlockBasic(b *block.Block) {
	I.bBasic = b
}

func (I *Input) SetBlockOver(b *block.Block) {
	I.bOver = b
}

func (I *Input) SetBlockClick(b *block.Block) {
	I.bClick = b
}

func (I *Input) SetNewRune(kCode sdl.Keysym, renderer *sdl.Renderer) error {
	var s string
	var newStr string
	var sizeTXT *objects.Size
	var sizeBlock *objects.Size
	var err error

	if I.status == objects.SClick {
		s = I.txt.GetTxt()

		if (kCode.Scancode >= sdl.SCANCODE_A && kCode.Scancode <= sdl.SCANCODE_Z) || kCode.Scancode == sdl.SCANCODE_SPACE {
			newStr = I.addKeyCode(s, kCode.Sym)
		} else {
			switch kCode.Scancode {
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

		I.txt.SetText(newStr)
		if err = I.txt.Init(renderer); err != nil {
			panic(err)
		}
	}
	if sizeTXT, err = I.txt.GetSize(); err != nil {
		panic(err)
	}
	if sizeBlock, err = I.bBasic.GetSize(); err != nil {
		panic(err)
	}
	if sizeTXT.W >= sizeBlock.W {
		I.txt.SetText(s)
		if err := I.txt.Init(renderer); err != nil {
			panic(err)
		}
		return errors.New(errorSizeTxt)
	}
	return nil
}

func (I *Input) addKeyCode(s string, kCode sdl.Keycode) string {
	idx := strings.Index(s, caret)
	newStr := s[:idx] + string(kCode) + s[idx:]
	return newStr
}

func (I *Input) caretRight(s string) string {
	idx := strings.Index(s, caret)
	if idx == len(s)-1 {
		return s
	}
	newStr := s[:idx] + string(s[idx+1]) + string(caret) + s[idx+2:]
	return newStr
}

func (I *Input) caretLeft(s string) string {
	idx := strings.Index(s, caret)
	if idx == 0 {
		return s
	}
	newStr := s[:idx-1] + string(caret) + s[idx-1:idx] + s[idx+1:]
	return newStr
}

func (I *Input) caretEnd(s string) string {
	idx := strings.Index(s, caret)

	if idx == len(s)-1 {
		return s
	}
	newStr := s[:idx] + s[idx+1:] + string(caret)
	return newStr
}

func (I *Input) caretBegin(s string) string {
	idx := strings.Index(s, caret)
	if idx == 0 {
		return s
	}
	newStr := string(caret) + s[:idx] + s[idx+1:]
	return newStr
}

func (I *Input) removeKeyBackspace(s string) string {
	idx := strings.Index(s, caret)
	if idx == 0 {
		return s
	}
	newStr := s[:idx-1] + s[idx:]
	return newStr
}

func (I *Input) removeKeyDelete(s string) string {
	idx := strings.Index(s, caret)
	if idx == len(s)-1 {
		return s
	}
	newStr := s[:idx+1] + s[idx+2:]
	return newStr
}
