package image

import (
	"errors"
	"sync"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/ymohl-cl/game-builder/objects"
)

const ()

// Image object with implementation of objet interface
type Image struct {
	// infos object
	status      uint8
	initialized bool

	// content object
	url      string
	size     *objects.Size
	position *objects.Position

	// sdl objects
	texture *sdl.Texture
	rect    sdl.Rect
}

// Init image object
func (I *Image) Init(r *sdl.Renderer) error {
	var surface *sdl.Surface
	var err error

	if r == nil {
		return errors.New(objects.ErrorRenderer)
	}

	if I.size == nil {
		return errors.New(objects.ErrorSize)
	}
	if I.position == nil {
		return errors.New(objects.ErrorPosition)
	}
	if I.url == "" {
		return errors.New(objects.ErrorTargetURL)
	}

	surface, err = img.Load(I.url)
	if err != nil {
		return err
	}
	defer surface.Free()

	I.texture, err = r.CreateTextureFromSurface(surface)
	if err != nil {
		return err
	}

	I.rect.X = I.position.X
	I.rect.Y = I.position.Y
	I.rect.W = I.size.W
	I.rect.H = I.size.H

	I.initialized = true
	return nil
}

// IsInit return status initialize
func (I Image) IsInit() bool {
	return I.initialized
}

// Close sdl objects
func (I *Image) Close() error {
	I.initialized = false
	if I.texture != nil {
		I.texture.Destroy()
	}
	return nil
}

// GetStatus provide the status
func (I *Image) GetStatus() uint8 {
	return I.status
}

// IsOver define if overred
func (I *Image) IsOver(x, y int32) bool {
	return false
}

// Click specify object is click
func (I *Image) Click() {
	return
}

// SetStatus change the object status if it is not Fix
func (I *Image) SetStatus(s uint8) {
	if I.status != objects.SFix {
		I.status = s
	}
}

// Draw the object Image.
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
