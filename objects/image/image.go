package image

import (
	"errors"
	"fmt"
	"sync"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/ymohl-cl/game-builder/objects"
)

const ()

// Image object with implementation of objet interface
type Image struct {
	// infos object
	status      uint8
	initialized bool

	// content object
	url      string
	size     *objects.Size
	position *objects.Position

	// sdl objects
	texture  *sdl.Texture
	rect     sdl.Rect
	renderer *sdl.Renderer
}

/*
** Functions image specifications
 */

// New create a new image object
func New(url string) (*Image, error) {
	i := new(Image)

	i.status = objects.SFix
	i.url = url
	return i, nil
}

// SetSize define the size
func (I *Image) SetSize(sz *objects.Size) error {
	if sz == nil {
		return errors.New("Can't add size because is nil")
	}

	I.size = sz
	return nil
}

// SetPosition define the position
func (I *Image) SetPosition(p *objects.Position) error {
	if p == nil {
		return errors.New("Can't add position because is nil")
	}

	I.position = p
	return nil
}

// GetSize provide size object
func (I Image) GetSize() *objects.Size {
	return I.size
}

// GetPosition provide position object
func (I Image) GetPosition() *objects.Position {
	return I.position
}

/*
** Interface objects functions
 */

// IsInit return status initialize
func (I Image) IsInit() bool {
	return I.initialized
}

// Init image object
func (I *Image) Init(r *sdl.Renderer) error {
	var surface *sdl.Surface
	var err error

	if r == nil {
		return errors.New("Can't init object because renderer is nil")
	}
	if I.size == nil {
		return errors.New("Size block not define")
	}
	if I.position == nil {
		return errors.New("Posisition block not define")
	}
	if I.url == "" {
		return errors.New("Image not specified")
	}

	surface, err = img.Load(I.url)
	if err != nil {
		return err
	}
	defer surface.Free()

	I.texture, err = r.CreateTextureFromSurface(surface)
	if err != nil {
		return err
	}

	I.rect.X = I.position.X
	I.rect.Y = I.position.Y
	I.rect.W = I.size.W
	I.rect.H = I.size.H

	I.renderer = r
	I.initialized = true
	return nil
}

// Close sdl objects
func (I *Image) Close() error {
	if I.texture != nil {
		I.texture.Destroy()
	}
	I.initialized = false
	return nil
}

// GetStatus provide the status
func (I *Image) GetStatus() uint8 {
	return I.status
}

// IsOver define if overred
func (I *Image) IsOver(x, y int32) bool {
	return false
}

// Click specify object is click
func (I *Image) Click() {
	return
}

// SetStatus change the object status if it is not Fix
func (I *Image) SetStatus(s uint8) {
	if I.status != objects.SFix {
		I.status = s
	}
}

// Draw the object Image.
func (I *Image) Draw(wg *sync.WaitGroup) error {
	wg.Add(1)
	defer wg.Done()

	if I.initialized == false {
		fmt.Println("...")
		return errors.New("Can't draw image object is not initialized")
	}

	sdl.Do(func() {
		I.renderer.Copy(I.texture, nil, &I.rect)
	})
	return nil
}
