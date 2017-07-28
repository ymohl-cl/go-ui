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
	urls        map[uint8]string

	// parameters object
	rect sdl.Rect

	// action click
	funcClick func(...interface{})
	dataClick []interface{}

	// sdl ressources
	textures map[uint8]*sdl.Texture
}

/*
** Builder method
 */

// New create Image object
func New(url string, x, y, w, h int32) *Image {
	i := Image{status: objects.SFix}

	i.rect = sdl.Rect{
		X: x,
		Y: y,
		W: w,
		H: h,
	}
	i.urls = make(map[uint8]string)
	i.urls[objects.SFix] = url

	i.textures = make(map[uint8]*sdl.Texture)

	return &i
}

// Clone object and return a new
func (I Image) Clone(r *sdl.Renderer) (*Image, error) {
	var err error
	var prime *Image

	prime = New(I.urls[objects.SFix], I.rect.X, I.rect.Y, I.rect.W, I.rect.H)
	prime.cloneStatus(&I)

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

// SetVariantStyle define styles to  interact with object.
func (I *Image) SetVariantStyle(basicURL, overURL, clickURL string) {
	I.urls[objects.SBasic] = basicURL
	I.urls[objects.SOver] = overURL
	I.urls[objects.SClick] = clickURL

	I.status = objects.SBasic
}

/*
** Getter method
 */

/*
** Updater method
 */
