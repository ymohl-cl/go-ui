package image

import (
	"errors"

	"github.com/ymohl-cl/game-builder/objects"
)

// New create a new object
func New(url string) (*Image, error) {
	i := new(Image)

	i.status = objects.SFix
	i.url = url
	return i, nil
}

// SetParams define the object on one time
func (I *Image) SetParams(x, y, w, h int32) {
	I.position = new(objects.Position)
	I.position.SetPosition(x, y)

	I.size = new(objects.Size)
	I.size.SetSize(w, h)
}

// SetSize define the size
func (I *Image) SetSize(sz *objects.Size) error {
	if sz == nil {
		return errors.New(objects.ErrorSize)
	}

	I.size = sz
	return nil
}

// SetPosition define the position
func (I *Image) SetPosition(p *objects.Position) error {
	if p == nil {
		return errors.New(objects.ErrorPosition)
	}

	I.position = p
	return nil
}

// GetSize provide size object
func (I Image) GetSize() (*objects.Size, error) {
	if I.size == nil {
		return nil, errors.New(objects.ErrorSize)
	}
	return I.size, nil
}

// GetPosition provide position object
func (I Image) getPosition() (*objects.Position, error) {
	if I.position == nil {
		return nil, errors.New(objects.ErrorPosition)
	}
	return I.position, nil
}
