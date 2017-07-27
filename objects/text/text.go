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
	font         *ttf.Font
	texture      *sdl.Texture
	underTexture *sdl.Texture
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

	T.rect = sdl.Rect{
		X: x,
		Y: y,
	}
	t.colors = make(map[uint8]sdl.Color)
	t.underColors = make(map[uint8]sdl.Color)
	return &t, nil
}

// Clone object and return a new
func (T Text) Clone(r *sdl.Renderer) (*Text, error) {
	var err error
	var prime *Text

	if prime, err = New(T.txt, T.size, T.fontURL); err != nil {
		return prime, err
	}
	prime.SetParams(T.rect.X, T.rect.Y, T.color.R, T.color.G, T.color.B, T.color.A)
	if T.style.exist {
		prime.cloneUnderStyle(T.style)
	}

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

// SetFixStyle
func (T *Text) SetFixStyle(r, g, b, a uint8) error {

}

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
			T.status = objects.SBasic
		default:
			return errors.New(objects.ErrorStatus)
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

/*
** Getter method
 */

// GetColor provide color object
func (T Text) GetColor() (r, g, b, a uint8) {
	return T.color.R, T.color.G, T.color.B, T.color.A
}

// GetTxt provide txt string
func (T Text) GetTxt() string {
	return T.txt
}

// GetIDSDL provide the idSDL session. Useful for timer display
func (T Text) GetIDSDL() uint8 {
	return T.idSDL
}

// NewIDSDL provide a new idSDL session.
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

/*
** Updater method
 */

// UpdateColor to change color of initialized object
func (T *Text) UpdateColor(red, green, blue, opacity uint8) {
	T.color = sdl.Color{
		R: red,
		G: green,
		B: blue,
		A: opacity,
	}
}

// UpdateText to change text of initialized object
func (T *Text) UpdateText(str string, r *sdl.Renderer) error {
	T.txt = str
	if err := T.Update(r); err != nil {
		return err
	}
	return nil
}
