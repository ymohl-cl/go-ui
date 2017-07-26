package text

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"github.com/ymohl-cl/game-builder/objects"
)

// Text object implementation
type Text struct {
	// infos object
	status      uint8 // get && set
	initialized bool  // get
	size        int
	fontURL     string

	// content object (getter and setter available)
	txt  string
	font *ttf.Font

	// parameters objects (getter and setter available)
	rect  sdl.Rect
	color sdl.Color

	// style objects (getter and setter available)
	style Styler

	// sizeSDL is Width and height of txt on the screen.
	idSDL   uint8
	texture *sdl.Texture
}

/*
** Builder method
 */

// New create Text object, it's necessary to call SetParams before call Init
func New(txt string, size int, fontURL string) (*Text, error) {
	var err error
	t := Text{txt: txt, status: objects.SFix, size: size, fontURL: fontURL}

	sdl.Do(func() {
		t.font, err = ttf.OpenFont(fontURL, size)
		if err != nil {
			panic(err)
		}
	})

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

// SetParams define object's position and color
func (T *Text) SetParams(x, y int32, red, green, blue, opacity uint8) {
	T.rect = sdl.Rect{
		X: x,
		Y: y,
	}
	T.color = sdl.Color{
		R: red,
		G: green,
		B: blue,
		A: opacity,
	}
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
