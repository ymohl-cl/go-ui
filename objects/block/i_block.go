package block

import (
	"errors"
	"sync"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/ymohl-cl/game-builder/objects"
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

// GetStatus object
func (B Block) GetStatus() uint8 {
	return B.status
}

// IsOver define if object and position parameters matches
func (B Block) IsOver(x, y int32) bool {
	return false
}

// Click define a click on object
func (B *Block) Click() {
	return
}

// SetStatus change object's status
func (B *Block) SetStatus(s uint8) {
	return
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

// GetPosition object (x, y)
func (B Block) GetPosition() (int32, int32) {
	return B.rect.X, B.rect.Y
}

// GetSize object (width, height)
func (B Block) GetSize() (int32, int32) {
	return B.rect.W, B.rect.H
}

// MoveTo by increment position with x and y parameters
func (B *Block) MoveTo(x, y int32) {
	B.rect.X += x
	B.rect.Y += y
}

// Update object after done modification
func (B *Block) Update(r *sdl.Renderer) error {
	return nil
}

// Draw the object
func (B *Block) Draw(wg *sync.WaitGroup, r *sdl.Renderer) {
	defer wg.Done()
	var err error

	sdl.Do(func() {
		if B.initialized == false {
			return
		}
		if r == nil {
			panic(errors.New(objects.ErrorRenderer))
		}

		if err = r.SetDrawColor(B.color.R, B.color.G, B.color.B, B.color.A); err != nil {
			panic(err)
		}

		switch B.style.block {
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
