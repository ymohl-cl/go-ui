package block

import (
	"errors"
	"sync"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/ymohl-cl/game-builder/objects"
)

const (
	Filled = 1
	Empty  = 2
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

/*
** Functions block specifications
 */
// New create a new Block object
func New(bStyle uint8) (*Block, error) {
	b := new(Block)

	switch bStyle {
	case Filled:
		b.style = Filled
	case Empty:
		b.style = Empty
	default:
		return nil, errors.New("Type block not recognized")
	}

	b.status = objects.SFix
	return b, nil
}

// SetSize
func (B *Block) SetSize(sz *objects.Size) error {
	if sz == nil {
		return errors.New("Can't add size because is nil")
	}

	B.size = sz
	return nil
}

// SetPosition
func (B *Block) SetPosition(p *objects.Position) error {
	if p == nil {
		return errors.New("Can't add position because is nil")
	}

	B.position = p
	return nil
}

// SetColor
func (B *Block) SetColor(c *objects.Color) error {
	if c == nil {
		return errors.New("Can't add color because is nil")
	}

	B.color = c
	return nil
}

// GetSize provide size object
func (B Block) GetSize() *objects.Size {
	return B.size
}

// GetPosiion provide position object
func (B Block) GetPosition() *objects.Position {
	return B.position
}

// GetColor provide color object
func (B Block) GetColor() *objects.Color {
	return B.color
}

/*
** Interface objects functions
 */

// IsInit return status initialize
func (B Block) IsInit() bool {
	return B.initialized
}

func (B *Block) Init(r *sdl.Renderer) error {
	if r == nil {
		return errors.New("Can't init object because renderer is nil")
	}
	if B.size == nil {
		return errors.New("Size block not define")
	}
	if B.position == nil {
		return errors.New("Posisition block not define")
	}
	if B.color == nil {
		return errors.New("Color block not define")
	}

	B.rect.X = B.position.X
	B.rect.Y = B.position.Y
	B.rect.W = B.size.W
	B.rect.H = B.size.H

	B.renderer = r
	B.initialized = true
	return nil
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
			panic(errors.New("Can't draw block object is not initialized"))
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
			err = errors.New("Draw block type no recognized")
		}

		if err != nil {
			panic(err)
		}
	})
}
