package button

import (
	"errors"
	"sync"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/ymohl-cl/game-builder/objects"
)

type Button struct {
	// infos object
	status      uint8
	initialized bool

	// content object
	cBasic    Content
	cOver     Content
	cClick    Content
	cFix      Content
	funcClick func(...interface{})
	dataClick []interface{}

	// sdl objects
	renderer *sdl.Renderer
}

func (B *Button) Init(r *sdl.Renderer) error {
	if r == nil {
		return errors.New(objects.ErrorRenderer)
	}
	B.renderer = r

	if err := B.cFix.checkContent(); err != nil {
		return err
	}
	if err := B.cFix.initContent(r); err != nil {
		return err
	}
	if err := B.cBasic.checkContent(); err != nil {
		return err
	}
	if err := B.cBasic.initContent(r); err != nil {
		return err
	}
	if err := B.cOver.checkContent(); err != nil {
		return err
	}
	if err := B.cOver.initContent(r); err != nil {
		return err
	}
	if err := B.cClick.checkContent(); err != nil {
		return err
	}
	if err := B.cClick.initContent(r); err != nil {
		return err
	}

	if B.funcClick == nil {
		return errors.New(objects.ErrorTargetURL)
	}

	B.initialized = true
	return nil
}

// IsInit return status initialize
func (B *Button) IsInit() bool {
	return B.initialized
}

func (B *Button) Close() error {
	if err := B.cFix.closeContent(); err != nil {
		return err
	}
	if err := B.cBasic.closeContent(); err != nil {
		return err
	}
	if err := B.cOver.closeContent(); err != nil {
		return err
	}
	if err := B.cClick.closeContent(); err != nil {
		return err
	}

	B.initialized = false
	return nil
}

func (B *Button) GetStatus() uint8 {
	return B.status
}

func (B *Button) IsOver(x, y int32) bool {
	var pos *objects.Position
	var size *objects.Size
	var err error

	switch B.status {
	case objects.SFix:
		if pos, err = B.cFix.getPosition(); err != nil {
			panic(err)
		}
		if size, err = B.cFix.getSize(); err != nil {
			panic(err)
		}
	case objects.SBasic:
		if pos, err = B.cBasic.getPosition(); err != nil {
			panic(err)
		}
		if size, err = B.cBasic.getSize(); err != nil {
			panic(err)
		}
	case objects.SOver:
		if pos, err = B.cOver.getPosition(); err != nil {
			panic(err)
		}
		if size, err = B.cOver.getSize(); err != nil {
			panic(err)
		}
	case objects.SClick:
		if pos, err = B.cClick.getPosition(); err != nil {
			panic(err)
		}
		if size, err = B.cClick.getSize(); err != nil {
			panic(err)
		}
	}

	if x > pos.X && x < pos.X+size.W {
		if y > pos.Y && y < pos.Y+size.H {
			return true
		}
	}
	return false
}

func (B *Button) Click() {
	B.funcClick(B.dataClick)
}

func (B *Button) SetStatus(s uint8) {
	if B.status != objects.SFix {
		B.status = s
	}
}

func (B *Button) Draw(wg *sync.WaitGroup) {
	defer wg.Done()

	if B.initialized == false {
		panic(errors.New(objects.ErrorNotInit))
	}

	switch B.status {
	case objects.SFix:
		B.cFix.drawContent(wg)
	case objects.SBasic:
		B.cBasic.drawContent(wg)
	case objects.SOver:
		B.cOver.drawContent(wg)
	case objects.SClick:
		B.cClick.drawContent(wg)
	}
}
