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
		for id, v := range I.urls {
			if surface, err = img.Load(v); err != nil {
				panic(err)
			}

			if I.textures[id], err = r.CreateTextureFromSurface(surface); err != nil {
				panic(err)
			}
			surface.Free()
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
		for _, texture := range I.textures {
			texture.Destroy()
		}
	})
	return nil
}

// SetAction to get it on click button
func (I *Image) SetAction(f func(...interface{}), d ...interface{}) {
	I.funcClick = f

	for _, v := range d {
		I.dataClick = append(I.dataClick, v)
	}
}

// SetStatus change object's status
func (I *Image) SetStatus(s uint8) {
	if I.status == objects.SFix {
		return
	}

	switch s {
	case objects.SBasic, objects.SOver, objects.SClick:
		I.status = s
	default:
		panic(errors.New(objects.ErrorStatus))
	}
}

// MoveTo by increment position with x and y parameters
func (I *Image) MoveTo(x, y int32) {
	I.rect.X += x
	I.rect.Y += y
}

// UpdatePosition object
func (I *Image) UpdatePosition(x, y int32) {
	I.rect.X = x
	I.rect.Y = y
}

// UpdateSize to change size of initialized object
func (I *Image) UpdateSize(w, h int32) {
	I.rect.W = w
	I.rect.H = h
}

// UpdateColor to change color of initialized object
func (I *Image) UpdateColor(red, g, b, a uint8, r *sdl.Renderer) {
	return
}

// IsOver define if object and position parameters matches
func (I Image) IsOver(xRef, yRef int32) bool {
	if I.status == objects.SFix {
		return false
	}

	if xRef > I.rect.X && xRef < I.rect.X+I.rect.W {
		if yRef > I.rect.Y && yRef < I.rect.Y+I.rect.H {
			return true
		}
	}
	return false
}

// GetStatus to object
func (I Image) GetStatus() uint8 {
	return I.status
}

// GetPosition object (x, y)
func (I Image) GetPosition() (int32, int32) {
	return I.rect.X, I.rect.Y
}

// GetColor object (current color by status)
func (I Image) GetColor() (r, g, b, a uint8) {
	return
}

// GetSize object (width, height)
func (I Image) GetSize() (int32, int32) {
	return I.rect.W, I.rect.H
}

// Click define a click on object
func (I Image) Click() {
	if I.status == objects.SFix || I.funcClick == nil {
		return
	}
	I.funcClick(I.dataClick)
}

// Draw the object
func (I Image) Draw(wg *sync.WaitGroup, r *sdl.Renderer) {
	defer wg.Done()

	sdl.Do(func() {
		if I.initialized == false {
			return
		}
		if r == nil {
			panic(errors.New(objects.ErrorRenderer))
		}

		if texture, ok := I.textures[I.status]; ok {
			r.Copy(texture, nil, &I.rect)
		} else {
			panic(errors.New(objects.ErrorBuild))
		}
	})
}
