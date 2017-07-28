package text

import (
	"errors"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"github.com/ymohl-cl/game-builder/objects"
)

// Text object implementation
type Text struct {
	// infos object
	status      uint8
	initialized bool
	idSDL       uint8

	// content object
	txt     string
	size    int
	fontURL string

	// parameters object
	rect        sdl.Rect
	underRect   sdl.Rect
	colors      map[uint8]sdl.Color
	underColors map[uint8]sdl.Color

	// action click
	funcClick func(...interface{})
	dataClick []interface{}

	// sdl ressources
	font          *ttf.Font
	textures      map[uint8]*sdl.Texture
	underTextures map[uint8]*sdl.Texture
}

/*
** Builder method
 */

// New create Text object, it's necessary to call SetParams before call Init
func New(txt string, size int, fontURL string, x, y int32) (*Text, error) {
	var err error
	t := Text{txt: txt, status: objects.SFix, size: size, fontURL: fontURL}

	sdl.Do(func() {
		t.font, err = ttf.OpenFont(fontURL, size)
		if err != nil {
			panic(err)
		}
	})

	t.rect = sdl.Rect{
		X: x,
		Y: y,
	}
	t.colors = make(map[uint8]sdl.Color)
	t.underColors = make(map[uint8]sdl.Color)
	t.textures = make(map[uint8]*sdl.Texture)
	t.underTextures = make(map[uint8]*sdl.Texture)
	return &t, nil
}

// Clone object and return a new
func (T Text) Clone(r *sdl.Renderer) (*Text, error) {
	var err error
	var prime *Text

	// Create object
	if prime, err = New(T.txt, T.size, T.fontURL, T.rect.X, T.rect.Y); err != nil {
		return prime, err
	}
	// Set color
	for id, c := range T.colors {
		if err = prime.SetVariantStyle(c.R, c.G, c.B, c.A, id); err != nil {
			return nil, err
		}
	}
	// Set underColors
	for id, c := range T.underColors {
		if err = prime.SetVariantUnderStyle(c.R, c.G, c.B, c.A, id); err != nil {
			return nil, err
		}
	}
	// Set position under
	prime.SetUnderPosition(T.underRect.X, T.underRect.Y)
	// Set functionClick
	prime.funcClick = T.funcClick
	prime.dataClick = T.dataClick

	if T.IsInit() {
		if err = prime.Init(r); err != nil {
			return nil, err
		}
	}
	return prime, nil
}

/*
** Setter method
 */

// SetVariantStyle define styles to interact with object.
func (T *Text) SetVariantStyle(r, g, b, a uint8, status ...uint8) error {
	for _, s := range status {
		switch s {
		case objects.SFix, objects.SBasic, objects.SOver, objects.SClick:
			T.colors[s] = sdl.Color{
				R: r,
				G: g,
				B: b,
				A: a,
			}
		default:
			return errors.New(objects.ErrorStatus)
		}
		if T.status == objects.SFix && s != objects.SFix {
			T.status = objects.SBasic
		}
	}
	return nil
}

// SetVariantUnderStyle define styles under to interact with object.
func (T *Text) SetVariantUnderStyle(r, g, b, a uint8, status ...uint8) error {
	for _, s := range status {
		switch s {
		case objects.SFix, objects.SBasic, objects.SOver, objects.SClick:
			T.underColors[s] = sdl.Color{
				R: r,
				G: g,
				B: b,
				A: a,
			}
		default:
			return errors.New(objects.ErrorStatus)
		}
	}
	return nil
}

// SetUnderPosition to underText
func (T *Text) SetUnderPosition(x, y int32) {
	T.underRect = sdl.Rect{
		X: x,
		Y: y,
	}
}

/*
** Getter method
 */

// GetTxt provide txt string
func (T Text) GetTxt() string {
	return T.txt
}

// GetIDSDL provide the idSDL session. Useful for timer display
func (T Text) GetIDSDL() uint8 {
	return T.idSDL
}

/*
** Updater method
 */

// NewIDSDL provide and set a new idSDL session.
func (T *Text) NewIDSDL() uint8 {
	var check uint8
	check--

	if T.idSDL == check {
		T.idSDL = 0
	} else {
		T.idSDL++
	}
	return T.idSDL
}

// UpdateText to change text of initialized object
func (T *Text) UpdateText(str string, r *sdl.Renderer) error {
	var err error

	T.txt = str
	if err = T.Close(); err != nil {
		return err
	}
	sdl.Do(func() {
		if err = T.updateTextures(r); err != nil {
			panic(err)
		}
		T.initialized = true
	})
	return nil
}
