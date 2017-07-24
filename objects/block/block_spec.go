package block

import (
	"errors"

	"github.com/ymohl-cl/game-builder/objects"
)

const (
	Filled = 1
	Empty  = 2
)

// New create a new object
func New(bStyle uint8) (*Block, error) {
	b := new(Block)

	switch bStyle {
	case Filled:
		b.style = Filled
	case Empty:
		b.style = Empty
	default:
		return nil, errors.New(objects.ErrorObjectStyle)
	}

	b.status = objects.SFix
	return b, nil
}

// SetParams define the object on one time
func (B *Block) SetParams(x, y, w, h int32, red, green, blue, opacity uint8) {
	B.position = new(objects.Position)
	B.position.SetPosition(x, y)

	B.size = new(objects.Size)
	B.size.SetSize(w, h)

	B.color = new(objects.Color)
	B.color.SetColor(red, green, blue, opacity)
}

// SetSize define size
func (B *Block) SetSize(sz *objects.Size) error {
	if sz == nil {
		return errors.New(objects.ErrorSize)
	}

	B.size = sz
	return nil
}

// SetPosition define position
func (B *Block) SetPosition(p *objects.Position) error {
	if p == nil {
		return errors.New(objects.ErrorPosition)
	}

	B.position = p
	return nil
}

// SetColor define color
func (B *Block) SetColor(c *objects.Color) error {
	if c == nil {
		return errors.New(objects.ErrorColor)
	}

	B.color = c
	return nil
}

// GetSize provide size object
func (B Block) GetSize() (*objects.Size, error) {
	if B.size == nil {
		return nil, errors.New(objects.ErrorSize)
	}
	return B.size, nil
}

// GetPosition provide position object
func (B Block) getPosition() (*objects.Position, error) {
	if B.position == nil {
		return nil, errors.New(objects.ErrorPosition)
	}
	return B.position, nil
}

// GetColor provide color object
func (B Block) GetColor() *objects.Color {
	return B.color
}
