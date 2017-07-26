package image

import (
	"errors"
	"sync"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/ymohl-cl/game-builder/objects"
)

// Init object to draw it. If error occurred, object can't be drawn
func (I *Image) Init(r *sdl.Renderer) error {
	var surface *sdl.Surface
	var err error

	if r == nil {
		return errors.New(objects.ErrorRenderer)
	}

	sdl.Do(func() {
		if surface, err = img.Load(I.url); err != nil {
			panic(err)
		}
		defer surface.Free()

		if I.texture, err = r.CreateTextureFromSurface(surface); err != nil {
			panic(err)
		}
		I.initialized = true
	})
	return nil
}

// IsInit return status initialize
func (I Image) IsInit() bool {
	return I.initialized
}

// Close sdl ressources needest to object
func (I *Image) Close() error {
	I.initialized = false
	sdl.Do(func() {
		if I.texture != nil {
			I.texture.Destroy()
		}
	})
	return nil
}

// Close sdl objects
func (I Image) GetStatus() uint8 {
	return I.status
}

// IsOver define if object and position parameters matches
func (I Image) IsOver(x, y int32) bool {
	return false
}

// Click define a click on object
func (I *Image) Click() {
	return
}

// SetStatus change object's status
func (I *Image) SetStatus(s uint8) {
	return
}

// UpdatePosition object
func (I *Image) UpdatePosition(x, y int32) {
	I.rect.X = x
	I.rect.Y = y
}

// GetPosition object (x, y)
func (I Image) GetPosition() (int32, int32) {
	return I.rect.X, I.rect.Y
}

// GetSize object (width, height)
func (I Image) GetSize() (int32, int32) {
	return I.rect.W, I.rect.H
}

// MoveTo by increment position with x and y parameters
func (I *Image) MoveTo(x, y int32) {
	I.rect.X += x
	I.rect.Y += y
}

// Update object after done modification
func (I *Image) Update(r *sdl.Renderer) error {
	return nil
}

// Draw the object
func (I *Image) Draw(wg *sync.WaitGroup, r *sdl.Renderer) {
	defer wg.Done()

	sdl.Do(func() {
		if I.initialized == false {
			return
		}
		if r == nil {
			panic(errors.New(objects.ErrorRenderer))
		}

		r.Copy(I.texture, nil, &I.rect)
	})
}
