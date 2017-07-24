package block

import (
	"errors"
	"sync"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/ymohl-cl/game-builder/objects"
)

type Block struct {
	// infos object
	status      uint8
	initialized bool
	style       uint8

	size     *objects.Size
	position *objects.Position
	color    *objects.Color

	// objects of sdl
	rect sdl.Rect
}

func (B *Block) Init(r *sdl.Renderer) error {
	if r == nil {
		return errors.New(objects.ErrorRenderer)
	}

	if B.size == nil {
		return errors.New(objects.ErrorSize)
	}
	if B.position == nil {
		return errors.New(objects.ErrorPosition)
	}
	if B.color == nil {
		return errors.New(objects.ErrorColor)
	}

	B.rect.X = B.position.X
	B.rect.Y = B.position.Y
	B.rect.W = B.size.W
	B.rect.H = B.size.H

	B.initialized = true
	return nil
}

// IsInit return status initialize
func (B Block) IsInit() bool {
	return B.initialized
}

func (B *Block) Close() error {
	B.initialized = false
	return nil
}

func (B *Block) GetStatus() uint8 {
	return B.status
}

func (B *Block) IsOver(x, y int32) bool {
	return false
}

func (B *Block) Click() {
	return
}

func (B *Block) SetStatus(s uint8) {
	if B.status != objects.SFix {
		B.status = s
	}
}

func (B *Block) UpdatePosition(x, y int32) {
	if B.position == nil {
		return
	}
	B.position.X = x
	B.position.Y = y
	B.rect.X = x
	B.rect.Y = y
}

func (B *Block) MoveTo(x, y int32) {
	if B.position == nil {
		return
	}
	B.position.X += x
	B.position.Y += y
	B.rect.X += x
	B.rect.Y += y
}

func (B *Block) GetPosition() (int32, int32) {
	if B.position == nil {
		return -1, -1
	}
	return B.position.X, B.position.Y
}

// Draw the object block.
func (B *Block) Draw(wg *sync.WaitGroup, r *sdl.Renderer) {
	defer wg.Done()

	sdl.Do(func() {
		if B.initialized == false {
			return
		}
		if r == nil {
			panic(errors.New(objects.ErrorRenderer))
		}

		err := r.SetDrawColor(B.color.Red, B.color.Green, B.color.Blue, B.color.Opacity)
		if err != nil {
			panic(err)
		}

		switch B.style {
		case Filled:
			err = r.FillRect(&B.rect)
		case Empty:
			err = r.DrawRect(&B.rect)
		default:
			err = errors.New(objects.ErrorObjectStyle)
		}

		if err != nil {
			panic(err)
		}
	})
}
