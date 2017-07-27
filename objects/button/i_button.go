package button

import (
	"errors"
	"sync"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/ymohl-cl/game-builder/objects"
	"github.com/ymohl-cl/game-builder/objects/image"
)

// Init object to draw it. If error occurred, object can't be drawn
func (B *Button) Init(r *sdl.Renderer) error {
	var err error

	if r == nil {
		return errors.New(objects.ErrorRenderer)
	}

	if B.funcClick == nil || B.txt == nil || B.block == nil {
		return errors.New(objects.ErrorBuild)
	}

	B.setParamsByStatus()

	if B.txt != nil {
		if err = B.txt.Init(r); err != nil {
			return err
		}
	}
	if B.block != nil {
		if err = B.block.Init(r); err != nil {
			return err
		}
	}
	if err = B.initImages(r); err != nil {
		return err
	}
	B.initialized = true
	return nil
}

// IsInit return status initialize
func (B *Button) IsInit() bool {
	return B.initialized
}

// Close sdl ressources needest to object
func (B *Button) Close() error {
	var err error

	B.initialized = false

	if err = B.txt.Close(); err != nil {
		return err
	}
	if err = B.block.Close(); err != nil {
		return err
	}
	if err = B.closeImages(); err != nil {
		return err
	}
	return nil
}

// GetStatus object
func (B Button) GetStatus() uint8 {
	return B.status
}

// IsOver define if object and position parameters matches
func (B Button) IsOver(xRef, yRef int32) bool {
	var x, y int32
	var w, h int32

	x, y = B.block.GetPosition()
	w, h = B.block.GetSize()

	if xRef > x && xRef < x+w {
		if yRef > y && yRef < y+h {
			return true
		}
	}
	return false
}

// Click define a click on object
func (B Button) Click() {
	B.funcClick(B.dataClick...)
}

// SetStatus change object's status
func (B *Button) SetStatus(s uint8) {
	switch s {
	case objects.SFix, objects.SBasic, objects.SOver, objects.SClick:
		B.status = s
	default:
		panic(errors.New(objects.ErrorStatus))
	}

	B.updateParamsByStatus()
}

// UpdatePosition object
func (B *Button) UpdatePosition(x, y int32) {
	B.fix.setPositionBlock(x, y)
	B.basic.setPositionBlock(x, y)
	B.over.setPositionBlock(x, y)
	B.click.setPositionBlock(x, y)
	B.updatePositionByStatus()
}

// GetPosition object (x, y)
func (B Button) GetPosition() (int32, int32) {
	return B.block.GetPosition()
}

// GetSize object (width, height)
func (B Button) GetSize() (int32, int32) {
	return B.block.GetSize()
}

// MoveTo by increment position with x and y parameters
func (B *Button) MoveTo(x, y int32) {
	B.fix.moveTo(x, y)
	B.basic.moveTo(x, y)
	B.over.moveTo(x, y)
	B.click.moveTo(x, y)

	B.updatePositionByStatus()
}

// Update object after done modification
func (B *Button) Update(r *sdl.Renderer) error {
	return nil
}

// Draw object
func (B *Button) Draw(wg *sync.WaitGroup, r *sdl.Renderer) {
	var img *image.Image
	defer wg.Done()

	if B.initialized == false {
		return
	}
	if r == nil {
		panic(errors.New(objects.ErrorRenderer))
	}

	wg.Add(1)
	B.block.Draw(wg, r)
	switch B.status {
	case objects.SFix:
		img = B.fix.img
	case objects.SBasic:
		img = B.basic.img
	case objects.SOver:
		img = B.over.img
	case objects.SClick:
		img = B.click.img
	}
	if img != nil {
		wg.Add(1)
		img.Draw(wg, r)
	}
	wg.Add(1)
	go B.txt.Draw(wg, r)
}
