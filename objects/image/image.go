package image

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/ymohl-cl/game-builder/objects"
)

// Image object implementation
type Image struct {
	// infos object
	status      uint8
	initialized bool
	url         string

	// parameters object
	rect sdl.Rect

	// sdl ressources
	texture *sdl.Texture
}

/*
** Builder method
 */

// New create Image object, it's necessary to call SetParams before call Init
func New(url string) (*Image, error) {
	i := Image{status: objects.SFix, url: url}

	return &i, nil
}

// Clone object and return a new
func (I Image) Clone(r *sdl.Renderer) (*Image, error) {
	var err error
	var prime *Image

	if prime, err = New(I.url); err != nil {
		return prime, err
	}
	prime.SetParams(I.rect.X, I.rect.Y, I.color.R, I.color.G, I.color.B, I.color.A)

	if I.IsInit() {
		if err = prime.Init(r); err != nil {
			return nil, err
		}
	}
	return prime, nil
}

/*
** Setter method
 */

// SetParams define object's position and color
func (I *Image) SetParams(x, y, w, h int32) {
	I.rect = sdl.Rect{
		X: x,
		Y: y,
		W: w,
		H: h,
	}
}

/*
** Getter method
 */

/*
** Updater method
 */
// UpdateSize to change size of initialized object
func (I *Image) UpdateSize(w, h int32) {
	I.rect.W = w
	I.rect.H = h
}
