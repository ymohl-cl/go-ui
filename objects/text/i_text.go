package text

import (
	"errors"
	"sync"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/ymohl-cl/go-ui/objects"
)

// Init object to draw it. If error occurred, object can't be drawn
func (T *Text) Init(r *sdl.Renderer) error {
	var err error

	if r == nil {
		return errors.New(objects.ErrorRenderer)
	}

	sdl.Do(func() {
		if err = T.updateTextures(r); err != nil {
			panic(err)
		}

		T.initialized = true
	})
	return nil
}

// IsInit return status initialize
func (T Text) IsInit() bool {
	return T.initialized
}

// Close sdl ressources needest to object
func (T *Text) Close() error {
	T.initialized = false

	sdl.Do(func() {
		for _, texture := range T.textures {
			texture.Destroy()
		}
		for _, texture := range T.underTextures {
			texture.Destroy()
		}
	})
	return nil
}

// SetAction to get it on click button
func (T *Text) SetAction(f func(...interface{}), d ...interface{}) {
	T.funcClick = f

	for _, v := range d {
		T.dataClick = append(T.dataClick, v)
	}
}

// SetStatus change object's status
func (T *Text) SetStatus(s uint8) {
	if T.status == objects.SFix {
		return
	}

	switch s {
	case objects.SBasic, objects.SOver, objects.SClick:
		T.status = s
	default:
		panic(errors.New(objects.ErrorStatus))
	}
}

// MoveTo by increment position with x and y parameters
func (T *Text) MoveTo(x, y int32) {
	T.rect.X += x
	T.rect.Y += y

	T.underRect.X += x
	T.underRect.Y += y
}

// UpdatePosition object
func (T *Text) UpdatePosition(x, y int32) {
	T.underRect.X = x - (T.underRect.X - T.rect.X)
	T.underRect.Y = y - (T.underRect.Y - T.rect.Y)
	T.rect.X = x
	T.rect.Y = y
}

// UpdateSize to change size of initialized object
func (T *Text) UpdateSize(w, h int32) {
	return
}

// UpdateColor to change color of initialized object
func (T *Text) UpdateColor(red, green, blue, opacity uint8, r *sdl.Renderer) {
	var err error

	T.colors[T.status] = sdl.Color{
		R: red,
		G: green,
		B: blue,
		A: opacity,
	}
	sdl.Do(func() {
		T.initialized = false
		if err = T.updateTextureByStatus(T.status, r); err != nil {
			panic(err)
		}
		T.initialized = true
	})
}

// IsOver define if object and position parameters matches
func (T Text) IsOver(xRef, yRef int32) bool {
	if T.status == objects.SFix {
		return false
	}

	if xRef > T.rect.X && xRef < T.rect.X+T.rect.W {
		if yRef > T.rect.Y && yRef < T.rect.Y+T.rect.H {
			return true
		}
	}
	return false
}

// GetStatus object
func (T Text) GetStatus() uint8 {
	return T.status
}

// GetPosition object (x, y)
func (T Text) GetPosition() (int32, int32) {
	return T.rect.X, T.rect.Y
}

// GetColor object (current color by status)
func (T Text) GetColor() (r, g, b, a uint8) {
	return T.colors[T.status].R, T.colors[T.status].G, T.colors[T.status].B, T.colors[T.status].A
}

// GetSize object (width, height)
func (T Text) GetSize() (int32, int32) {
	return T.rect.W, T.rect.H
}

// Click define a click on object
func (T Text) Click() {
	if T.status == objects.SFix || T.funcClick == nil {
		return
	}
	T.funcClick(T.dataClick...)
}

// Draw object
func (T Text) Draw(wg *sync.WaitGroup, r *sdl.Renderer) {
	defer wg.Done()

	sdl.Do(func() {
		if r == nil {
			panic(errors.New(objects.ErrorRenderer))
		}

		if texture, ok := T.underTextures[T.status]; ok {
			rect := sdl.Rect{
				X: T.rect.X - (T.rect.W / 2),
				Y: T.rect.Y - (T.rect.H / 2),
				W: T.rect.W,
				H: T.rect.H,
			}
			if err := r.Copy(texture, nil, &rect); err != nil {
				panic(err)
			}
		}
		if texture, ok := T.textures[T.status]; ok {
			rect := sdl.Rect{
				X: T.rect.X - (T.rect.W / 2),
				Y: T.rect.Y - (T.rect.H / 2),
				W: T.rect.W,
				H: T.rect.H,
			}
			if err := r.Copy(texture, nil, &rect); err != nil {
				panic(err)
			}
		}
	})
}
