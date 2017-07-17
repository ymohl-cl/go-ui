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
	funcClick func(...interface{}) string
	dataClick []interface{}

	// sdl objects
	renderer *sdl.Renderer
}

// IsInit return status initialize
func (B *Button) IsInit() bool {
	return B.initialized
}

func (B *Button) Init(r *sdl.Renderer) error {
	if B.status == objects.SFix {
		if err := B.cFix.checkContent(); err != nil {
			return err
		}
		if err := B.cFix.initContent(r); err != nil {
			return err
		}
	} else {
		if B.funcClick == nil {
			return errors.New("Function not define to the button")
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
	}

	B.initialized = true
}

func (B *Button) Close() error {
	if B.status == objects.SFix {
		if err := B.cFix.closeContent(); err != nil {
			return err
		}
	} else {
		if err := B.cBasic.closeContent(); err != nil {
			return err
		}
		if err := B.cBasic.closeContent(r); err != nil {
			return err
		}

		if err := B.cOver.closeContent(); err != nil {
			return err
		}
		if err := B.cOver.closeContent(r); err != nil {
			return err
		}

		if err := B.cClick.closeContent(); err != nil {
			return err
		}
		if err := B.cClick.closeContent(r); err != nil {
			return err
		}
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

	switch B.status {
	case objects.SFix:
		pos = B.cFix.block.GetPosition()
		size = B.cFix.block.GetSize()
	case objects.SBasic:
		pos = B.cBasic.block.GetPosition()
		size = B.cBasic.block.GetSize()
	case objects.SOver:
		pos = B.cOver.block.GetPosition()
		size = B.cOver.block.GetSize()
	case objects.Click:
		pos = B.cClick.block.GetPosition()
		size = B.cClick.block.GetSize()
	}

	if !pos || size {
		return false
	}
	if x > pos.X && x < pos.X+size.W {
		if y > pos.Y && y < pos.Y+size.H {
			return true
		}
	}
	return false
}

func (B *Button) Click() {
	B.SetStatus(objects.SClick)
	B.funcClick(B.dataClick)
}

func (B *Button) SetStatus(s uint8) {
	if B.status != objects.SFix {
		B.status = s
	}
}

func (B *Button) Draw(r *sdl.Renderer, wg *sync.WaitGroup) error {
	if r == nil {
		return errors.New("Can't draw buttun because renderer is nil")
	}
	if wg == nil {
		return errors.New("Can't draw buttun because sync WaitGroup not define")
	}
	if B.initialized == false {
		return errors.New("Can't draw because object is not initialized")
	}

	wg.Add(1)
	defer wg.Done()

	switch B.status {
	case objects.SFix:
		if err := B.cFix.drawContent(); err != nil {
			return err
		}
	case objects.Sbasic:
		if err := B.cBasic.drawContent(); err != nil {
			return err
		}
	case objects.SOver:
		if err := B.cOver.drawContent(); err != nil {
			return err
		}
	case objects.SClick:
		if err := B.cClick.drawContent(); err != nil {
			return err
		}
	}

	return nil
}
