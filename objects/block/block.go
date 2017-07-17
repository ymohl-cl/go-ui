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
	rect     sdl.Rect
	renderer *sdl.Renderer
}

func (B *Block) Init(r *sdl.Renderer) error {
	if r == nil {
		return errors.New(objects.ErrorRenderer)
	}
	B.renderer = r

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

// Draw the object block.
func (B *Block) Draw(wg *sync.WaitGroup) {
	defer wg.Done()

	sdl.Do(func() {
		if B.initialized == false {
			panic(errors.New(objects.ErrorNotInit))
		}

		err := B.renderer.SetDrawColor(B.color.Red, B.color.Green, B.color.Blue, B.color.Opacity)
		if err != nil {
			panic(err)
		}

		switch B.style {
		case Filled:
			err = B.renderer.FillRect(&B.rect)
		case Empty:
			err = B.renderer.DrawRect(&B.rect)
		default:
			err = errors.New(objects.ErrorObjectStyle)
		}

		if err != nil {
			panic(err)
		}
	})
}
