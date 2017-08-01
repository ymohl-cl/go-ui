package input

import (
	"errors"
	"sync"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/ymohl-cl/game-builder/objects"
)

// Init object to draw it. If error occurred, object can't be drawn
func (I *Input) Init(r *sdl.Renderer) error {
	var err error

	if r == nil {
		return errors.New(objects.ErrorRenderer)
	}

	if I.Txt == nil || I.block == nil {
		return errors.New(objects.ErrorBuild)
	}

	if !I.Txt.IsInit() {
		if err = I.Txt.Init(r); err != nil {
			return err
		}
	}
	if !I.block.IsInit() {
		if err = I.block.Init(r); err != nil {
			return err
		}
	}

	I.initialized = true
	return nil
}

// IsInit return status initialize
func (I Input) IsInit() bool {
	return I.initialized
}

// Close sdl ressources needest to object
func (I *Input) Close() error {
	I.initialized = false

	if err := I.Txt.Close(); err != nil {
		return err
	}
	if err := I.block.Close(); err != nil {
		return err
	}
	return nil
}

// SetAction to get it on click button
func (I *Input) SetAction(f func(...interface{}), d ...interface{}) {
	I.funcClick = f

	for _, v := range d {
		I.dataClick = append(I.dataClick, v)
	}
}

// SetStatus change object's status
func (I *Input) SetStatus(s uint8) {
	if I.block != nil {
		I.block.SetStatus(s)
	}
	if I.Txt != nil {
		I.Txt.SetStatus(s)
	}
	I.status = s
}

// MoveTo by increment position with x and y parameters
func (I *Input) MoveTo(x, y int32) {
	if I.block != nil {
		I.block.MoveTo(x, y)
	}
	if I.Txt != nil {
		I.Txt.MoveTo(x, y)
	}
}

// UpdatePosition object
func (I *Input) UpdatePosition(x, y int32) {
	var diffX, diffY int32

	if I.block != nil {
		blX, blY := I.block.GetPosition()
		diffX = blX - x
		diffY = blY - y
		I.block.UpdatePosition(x, y)
	}
	if I.Txt != nil {
		if diffX == 0 && diffY == 0 {
			I.Txt.UpdatePosition(x, y)
		} else {
			I.Txt.MoveTo(x, y)
		}
	}

}

// UpdateSize to change size of initialized object
func (I *Input) UpdateSize(w, h int32) {
	if I.block != nil {
		I.block.UpdateSize(w, h)
	}
	if I.Txt != nil {
		I.Txt.UpdateSize(w, h)
	}
	return
}

// UpdateColor to change color of initialized object
func (I *Input) UpdateColor(red, green, blue, opacity uint8, r *sdl.Renderer) {
	return
}

// IsOver define if object and position parameters matches
func (I Input) IsOver(xRef, yRef int32) bool {
	if I.block != nil {
		return I.block.IsOver(xRef, yRef)
	} else if I.Txt != nil {
		return I.Txt.IsOver(xRef, yRef)
	}
	return false
}

// GetStatus object
func (I Input) GetStatus() uint8 {
	return I.status
}

// GetPosition object (x, y)
func (I Input) GetPosition() (int32, int32) {
	if I.block != nil {
		return I.block.GetPosition()
	} else if I.Txt != nil {
		return I.Txt.GetPosition()
	}
	return -1, -1
}

// GetColor object (current color by status)
func (I Input) GetColor() (r, g, b, a uint8) {
	return
}

// GetSize object (width, height)
func (I Input) GetSize() (int32, int32) {
	if I.block != nil {
		return I.block.GetSize()
	} else if I.Txt != nil {
		return I.Txt.GetSize()
	}
	return -1, -1
}

// Click define a click on object
func (I *Input) Click() {
	if I.status != objects.SFix {
		if I.funcClick != nil {
			I.funcClick(I.dataClick)
		}
		I.status = objects.SClick

		if I.block != nil {
			I.block.Click()
		}
		if I.Txt != nil {
			I.Txt.Click()
		}
	}
}

// Draw object
func (I *Input) Draw(wg *sync.WaitGroup, r *sdl.Renderer) {
	defer wg.Done()

	if I.initialized == false {
		return
	}
	if r == nil {
		panic(errors.New(objects.ErrorRenderer))
	}

	if I.block != nil {
		wg.Add(1)
		I.block.Draw(wg, r)
	}
	if I.Txt != nil {
		wg.Add(1)
		go I.Txt.Draw(wg, r)
	}
}
