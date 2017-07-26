package text

import (
	"errors"
	"sync"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/ymohl-cl/game-builder/objects"
)

// Init object to draw it. If error occurred, object can't be drawn
func (T *Text) Init(r *sdl.Renderer) error {
	var err error
	var surface *sdl.Surface
	var uSurface *sdl.Surface

	if r == nil {
		return errors.New(objects.ErrorRenderer)
	}

	sdl.Do(func() {
		if surface, err = T.font.RenderUTF8_Solid(T.txt, T.color); err != nil {
			panic(err)
		}
		defer surface.Free()

		if T.texture, err = r.CreateTextureFromSurface(surface); err != nil {
			panic(err)
		}

		if T.style.exist {
			if uSurface, err = T.font.RenderUTF8_Solid(T.txt, T.style.color); err != nil {
				panic(err)
			}
			defer uSurface.Free()

			if T.style.texture, err = r.CreateTextureFromSurface(uSurface); err != nil {
				panic(err)
			}
		}

		T.setSize(surface.W, surface.H)
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
		if T.font != nil {
			T.font.Close()
		}
		if T.texture != nil {
			T.texture.Destroy()
		}
		if T.style.texture != nil {
			T.style.texture.Destroy()
		}
	})
	return nil
}

// GetStatus object
func (T Text) GetStatus() uint8 {
	return T.status
}

// IsOver define if object and position parameters matches
func (T Text) IsOver(x, y int32) bool {
	return false
}

// Click define a click on object
func (T *Text) Click() {
	return
}

// SetStatus change object's status
func (T *Text) SetStatus(s uint8) {
	return
}

// UpdatePosition object
func (T *Text) UpdatePosition(x, y int32) {
	if T.style.exist {
		T.setUnderPosition(x, y)
	}
	T.rect.X = x
	T.rect.Y = y
}

// GetPosition object (x, y)
func (T Text) GetPosition() (int32, int32) {
	return T.rect.X, T.rect.Y
}

// GetSize object (width, height)
func (T Text) GetSize() (int32, int32) {
	return T.rect.W, T.rect.H
}

// MoveTo by increment position with x and y parameters
func (T *Text) MoveTo(x, y int32) {
	T.rect.X += x
	T.rect.Y += y

	if T.style.exist {
		T.style.rect.X += x
		T.style.rect.Y += y
	}
}

// Update object after done modification
func (T *Text) Update(r *sdl.Renderer) error {
	T.initialized = false

	sdl.Do(func() {
		if T.texture != nil {
			T.texture.Destroy()
		}
		if T.style.texture != nil {
			T.style.texture.Destroy()
		}
	})

	if err := T.Init(r); err != nil {
		return err
	}
	return nil
}

// Draw object
func (T Text) Draw(wg *sync.WaitGroup, r *sdl.Renderer) {
	defer wg.Done()

	sdl.Do(func() {
		if T.initialized == false {
			return
		}
		if r == nil {
			panic(errors.New(objects.ErrorRenderer))
		}

		if T.style.exist {
			color := T.style.color
			if err := r.SetDrawColor(color.R, color.G, color.B, color.A); err != nil {
				panic(err)
			}

			rect := T.style.rect
			rect.X -= (rect.W / 2)
			rect.Y -= (rect.H / 2)
			if err := r.Copy(T.style.texture, nil, &rect); err != nil {
				panic(err)
			}
		}

		if err := r.SetDrawColor(T.color.R, T.color.G, T.color.B, T.color.A); err != nil {
			panic(err)
		}

		rect := T.rect
		rect.X -= (rect.W / 2)
		rect.Y -= (rect.H / 2)
		if err := r.Copy(T.texture, nil, &rect); err != nil {
			panic(err)
		}
	})
}
