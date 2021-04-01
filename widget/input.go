package widget

import "github.com/veandco/go-sdl2/sdl"

type Input struct {
	Text
	indexCursor int
}

func NewInput(font Font) (*Input, error) {
	var i Input
	var err error
	var w, h int

	i.font = font
	if w, h, err = i.font.driver.SizeUTF8(""); err != nil {
		return nil, err
	}

	i.SetSize(int32(w), int32(h))

	return &i, nil
}

func (i *Input) KeyboardEvent(key *sdl.KeyboardEvent) {
	if !i.focus {
		return
	}
	if key.State != sdl.RELEASED {
		return
	}
	if (key.Keysym.Scancode >= sdl.SCANCODE_A && key.Keysym.Scancode <= sdl.SCANCODE_Z) ||
		(key.Keysym.Scancode >= sdl.SCANCODE_1 && key.Keysym.Scancode <= sdl.SCANCODE_0) ||
		(key.Keysym.Scancode == sdl.SCANCODE_SPACE) {
		i.Set(i.str[:i.indexCursor] + string(key.Keysym.Sym) + i.str[i.indexCursor:])
		i.indexCursor++
	} else {
		switch key.Keysym.Scancode {
		case sdl.SCANCODE_RIGHT:
			if len(i.str) > i.indexCursor {
				i.indexCursor++
			}
		case sdl.SCANCODE_LEFT:
			if i.indexCursor > 0 {
				i.indexCursor--
			}
		case sdl.SCANCODE_BACKSPACE:
			if i.indexCursor > 0 {
				i.Set(i.str[:i.indexCursor-1] + i.str[i.indexCursor:])
				i.indexCursor--
			}
		case sdl.SCANCODE_DELETE:
			if i.indexCursor+1 < len(i.str) {
				i.Set(i.str[:i.indexCursor] + i.str[i.indexCursor+1:])
			}
		}
	}
}

func (i *Input) Click() {
	i.indexCursor = len(i.str)
	i.widget.Click()
}

func (i *Input) Render(r *sdl.Renderer) error {
	if i.focus {
		withoutCursor := i.str
		if i.indexCursor == len(i.str) {
			i.Set(withoutCursor + "|")
		} else {
			i.Set(i.str[:i.indexCursor] + "|" + i.str[i.indexCursor:])
		}
		defer i.Set(withoutCursor)
	}

	if err := i.Text.Render(r); err != nil {
		return err
	}
	return nil
}
