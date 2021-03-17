package button

import (
	"errors"
	"sync"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/ymohl-cl/go-ui/objects"
)

// Init object to draw it. If error occurred, object can't be drawn
func (B *Button) Init(r *sdl.Renderer) error {
	var err error

	if r == nil {
		return errors.New(objects.ErrorRenderer)
	}

	if B.block != nil {
		if err = B.block.Init(r); err != nil {
			return err
		}
	}
	if B.img != nil {
		if err = B.img.Init(r); err != nil {
			return err
		}
	}

	if B.txt != nil {
		if err = B.txt.Init(r); err != nil {
			return err
		}
	}

	B.initialized = true
	return nil
}

// IsInit return status initialize
func (B *Button) IsInit() bool {
	return B.initialized
}

// Close sdl ressources needest to object
func (B *Button) Close() error {
	var err error

	B.initialized = false
	if B.block != nil {
		if err = B.block.Close(); err != nil {
			return err
		}
	}
	if B.img != nil {
		if err = B.img.Close(); err != nil {
			return err
		}
	}
	if B.txt != nil {
		if err = B.txt.Close(); err != nil {
			return err
		}
	}

	return nil
}

// SetAction to get it on click button
func (B *Button) SetAction(f func(...interface{}), d ...interface{}) {
	B.funcClick = f

	for _, v := range d {
		B.dataClick = append(B.dataClick, v)
	}
}

// SetStatus change object's status
func (B *Button) SetStatus(s uint8) {
	if B.block != nil {
		B.block.SetStatus(s)
	}
	if B.img != nil {
		B.img.SetStatus(s)
	}
	if B.txt != nil {
		B.txt.SetStatus(s)
	}
}

// MoveTo by increment position with x and y parameters
func (B *Button) MoveTo(x, y int32) {
	if B.block != nil {
		B.block.MoveTo(x, y)
	}
	if B.img != nil {
		B.img.MoveTo(x, y)
	}
	if B.txt != nil {
		B.txt.MoveTo(x, y)
	}
}

// UpdatePosition object
func (B *Button) UpdatePosition(x, y int32) {
	var diffX, diffY int32

	if B.block != nil {
		blX, blY := B.block.GetPosition()
		diffX = x - blX
		diffY = y - blY
		B.block.UpdatePosition(x, y)
	}
	if B.img != nil {
		if diffX == 0 && diffY == 0 {
			iX, iY := B.img.GetPosition()
			diffX = x - iX
			diffY = y - iY
			B.img.UpdatePosition(x, y)
		} else {
			B.img.MoveTo(diffX, diffY)
		}
	}
	if B.txt != nil {
		if diffX == 0 && diffY == 0 {
			B.txt.UpdatePosition(x, y)
		} else {
			B.txt.MoveTo(diffX, diffY)
		}
	}
}

// UpdateSize to change size of initialized object
func (B *Button) UpdateSize(w, h int32) {
	if B.block != nil {
		B.block.UpdateSize(w, h)
	}
	if B.img != nil {
		B.img.UpdateSize(w, h)
	}
	if B.txt != nil {
		B.txt.UpdateSize(w, h)
	}
	return
}

// UpdateColor to change color of initialized object
func (B *Button) UpdateColor(red, green, blue, opacity uint8, r *sdl.Renderer) {
	return
}

// IsOver define if object and position parameters matches
func (B Button) IsOver(xRef, yRef int32) bool {
	if B.block != nil {
		return B.block.IsOver(xRef, yRef)
	} else if B.img != nil {
		return B.img.IsOver(xRef, yRef)
	} else if B.txt != nil {
		return B.txt.IsOver(xRef, yRef)
	}

	return false
}

// GetStatus object
func (B Button) GetStatus() uint8 {
	if B.block != nil && B.block.GetStatus() != objects.SFix {
		return B.block.GetStatus()
	} else if B.img != nil && B.img.GetStatus() != objects.SFix {
		return B.img.GetStatus()
	} else if B.txt != nil && B.txt.GetStatus() != objects.SFix {
		return B.txt.GetStatus()
	}
	return objects.SFix
}

// GetPosition object (x, y)
func (B Button) GetPosition() (int32, int32) {
	if B.block != nil {
		return B.block.GetPosition()
	} else if B.img != nil {
		return B.img.GetPosition()
	} else if B.txt != nil {
		return B.txt.GetPosition()
	}
	return -1, -1
}

// GetColor object (current color by status)
func (B Button) GetColor() (r, g, b, a uint8) {
	return
}

// GetSize object (width, height)
func (B Button) GetSize() (int32, int32) {
	if B.block != nil {
		return B.block.GetSize()
	} else if B.img != nil {
		return B.img.GetSize()
	} else if B.txt != nil {
		return B.txt.GetSize()
	}
	return -1, -1
}

// Click define a click on object
func (B Button) Click() {
	if B.funcClick != nil {
		B.funcClick(B.dataClick...)
	}
}

// Draw object
func (B *Button) Draw(wg *sync.WaitGroup, r *sdl.Renderer) {
	defer wg.Done()

	if r == nil {
		panic(errors.New(objects.ErrorRenderer))
	}

	if B.block != nil {
		wg.Add(1)
		B.block.Draw(wg, r)
	}
	if B.img != nil {
		wg.Add(1)
		B.img.Draw(wg, r)
	}
	if B.txt != nil {
		wg.Add(1)
		B.txt.Draw(wg, r)
	}
}
