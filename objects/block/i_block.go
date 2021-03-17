package block

import (
	"errors"
	"sync"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/ymohl-cl/go-ui/objects"
)

// Init object to draw it. If error occurred, object can't be drawn
func (B *Block) Init(r *sdl.Renderer) error {
	B.initialized = true
	return nil
}

// IsInit return status initialize
func (B Block) IsInit() bool {
	return B.initialized
}

// Close sdl ressources needest to object
func (B *Block) Close() error {
	B.initialized = false
	return nil
}

// SetAction to get it on click button
func (B *Block) SetAction(f func(...interface{}), d ...interface{}) {
	B.funcClick = f

	for _, v := range d {
		B.dataClick = append(B.dataClick, v)
	}
}

// SetStatus change object's status
func (B *Block) SetStatus(s uint8) {
	if B.status == objects.SFix {
		return
	}

	switch s {
	case objects.SBasic, objects.SOver, objects.SClick:
		B.status = s
	default:
		panic(errors.New(objects.ErrorStatus))
	}
}

// MoveTo by increment position with x and y parameters
func (B *Block) MoveTo(x, y int32) {
	B.rect.X += x
	B.rect.Y += y
}

// UpdatePosition object
func (B *Block) UpdatePosition(x, y int32) {
	B.rect.X = x
	B.rect.Y = y
}

// UpdateSize object
func (B *Block) UpdateSize(w, h int32) {
	B.rect.W = w
	B.rect.H = h
}

// UpdateColor to change color of initialized object
func (B *Block) UpdateColor(red, green, blue, opacity uint8, r *sdl.Renderer) {
	B.colors[B.status] = sdl.Color{
		R: red,
		G: green,
		B: blue,
		A: opacity,
	}
}

// IsOver define if object and position parameters matches
func (B Block) IsOver(xRef, yRef int32) bool {
	if B.status == objects.SFix {
		return false
	}

	if xRef > B.rect.X && xRef < B.rect.X+B.rect.W {
		if yRef > B.rect.Y && yRef < B.rect.Y+B.rect.H {
			return true
		}
	}
	return false
}

// GetStatus object
func (B Block) GetStatus() uint8 {
	return B.status
}

// GetPosition object (x, y)
func (B Block) GetPosition() (int32, int32) {
	return B.rect.X, B.rect.Y
}

// GetColor object (current color by status)
func (B Block) GetColor() (r, g, b, a uint8) {
	return B.colors[B.status].R, B.colors[B.status].G, B.colors[B.status].B, B.colors[B.status].A
}

// GetSize object (width, height)
func (B Block) GetSize() (int32, int32) {
	return B.rect.W, B.rect.H
}

// Click define a click on object
func (B Block) Click() {
	if B.status == objects.SFix || B.funcClick == nil {
		return
	}
	B.funcClick(B.dataClick...)
	return
}

// Draw the object
func (B Block) Draw(wg *sync.WaitGroup, r *sdl.Renderer) {
	defer wg.Done()
	var err error

	sdl.Do(func() {
		if r == nil {
			panic(errors.New(objects.ErrorRenderer))
		}

		if c, ok := B.colors[B.status]; ok {
			if err = r.SetDrawColor(c.R, c.G, c.B, c.A); err != nil {
				panic(err)
			}
		}

		switch B.style {
		case Filled:
			err = r.FillRect(&B.rect)
		case Border:
			err = r.DrawRect(&B.rect)
		}

		if err != nil {
			panic(err)
		}
	})
}
