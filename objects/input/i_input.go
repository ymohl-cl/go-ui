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

	if I.txt == nil || I.block == nil {
		return errors.New(objects.ErrorBuild)
	}

	I.setParamsByStatus()

	if err = I.txt.Init(r); err != nil {
		return err
	}
	if err = I.block.Init(r); err != nil {
		return err
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

	if err := I.txt.Close(); err != nil {
		return err
	}
	if err := I.block.Close(); err != nil {
		return err
	}
	return nil
}

// GetStatus object
func (I Input) GetStatus() uint8 {
	return I.status
}

// IsOver define if object and position parameters matches
func (I Input) IsOver(xRef, yRef int32) bool {
	var x, y int32
	var w, h int32

	x, y = I.block.GetPosition()
	w, h = I.block.GetSize()
	if xRef > x && xRef < x+w {
		if yRef > y && yRef < y+h {
			return true
		}
	}
	return false
}

// Click define a click on object
func (I *Input) Click() {
	if I.status != objects.SFix {
		I.status = objects.SClick
	}
}

// SetStatus change object's status
func (I *Input) SetStatus(s uint8) {
	switch s {
	case objects.SFix, objects.SBasic, objects.SOver, objects.SClick:
		I.status = s
	default:
		panic(errors.New(objects.ErrorStatus))
	}

	I.updateParamsByStatus()
}

// UpdatePosition object
func (I *Input) UpdatePosition(x, y int32) {
	I.fix.setPositionBlock(x, y)
	I.basic.setPositionBlock(x, y)
	I.over.setPositionBlock(x, y)
	I.click.setPositionBlock(x, y)

	I.updatePositionByStatus()
}

// GetPosition object (x, y)
func (I Input) GetPosition() (int32, int32) {
	return I.block.GetPosition()
}

// GetSize object (width, height)
func (I Input) GetSize() (int32, int32) {
	return I.block.GetSize()
}

// MoveTo by increment position with x and y parameters
func (I *Input) MoveTo(x, y int32) {
	I.fix.moveTo(x, y)
	I.basic.moveTo(x, y)
	I.over.moveTo(x, y)
	I.click.moveTo(x, y)

	I.updatePositionByStatus()
}

// Update object after done modification
func (I *Input) Update(r *sdl.Renderer) error {
	return nil
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

	wg.Add(1)
	I.block.Draw(wg, r)
	wg.Add(1)
	go I.txt.Draw(wg, r)
}
