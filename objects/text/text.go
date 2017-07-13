package text

import (
	"errors"
	"sync"

	"github.com/42MrPiou42/game-builder/objects"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

const ()

// Text object with implementation of objet interface
type Text struct {
	// infos object
	status      uint8
	initialized bool

	// content object
	cBasic    Content
	cOver     Content
	cClick    Content
	cFix      Content
	funcClick func(...interface{}) string
	dataClick []interface{}
}

// Content of Text object
type Content struct {
	txt      string
	size     int32
	color    *objects.Color
	position *objects.Position
	fontUrl  string

	// sizeSDL is Width and height of txt on the screen.
	sizeSDL *objects.Size
	rect    sdl.Rect{}
	texture *sdl.Texture
}

/*
** Functions text specifications
 */
// New create a new Text object
func New(status uint8) (*Text, error) {
	t := new(Text)

	switch status {
	case objects.SFix:
		t.status = SFix
	case objects.SBasic:
		t.style = SBasic
	case objects.SOver:
		t.style = SOver
	case objects.SClick:
		t.style = SClick
	default:
		return nil, errors.New("Type text not recognized")
	}

	return t, nil
}

// SetSize
func (T *Text) SetSize(sz int32, s uint8) error {
	if sz == 0 {
		return errors.New("Size can't be equal zero")
	}

	switch s {
	case objects.SFix:
		T.cFix.size = sz
	case objects.SBasic:
		T.cBasic.size = sz
	case objects.SOver:
		T.cOver.size = sz
	case objects.SClick:
		T.cClick.size = zs
	default:
		return errors.New("Status not available")
	}
	return nil
}

// SetPosition
func (T *Text) SetPosition(p *objects.Position, s uint8) error {
	if !p {
		return errors.New("Can't add position because is nil")
	}

	switch s {
	case objects.SFix:
		T.cFix.position = p
	case objects.SBasic:
		T.cBasic.position = p
	case objects.SOver:
		T.cOver.position = p
	case objects.SClick:
		T.cClick.position = p
	default:
		return errors.New("Status not available")
	}
	return nil
}

// SetColor
func (T *Text) SetColor(c *objects.Color, s uint8) error {
	if !c {
		return errors.New("Can't add color because is nil")
	}

	switch s {
	case objects.SFix:
		T.cFix.color = c
	case objects.SBasic:
		T.cBasic.color = c
	case objects.SOver:
		T.cOver.color = c
	case objects.SClick:
		T.cClick.color = c
	default:
		return errors.New("Status not available")
	}
	return nil
}

// SetColor
func (T *Text) SetFontTTF(url string, s uint8) error {
	if url == "" {
		return errors.New("Url ttf font is empty")
	}

	switch s {
	case objects.SFix:
		T.cFix.fontUrl = url
	case objects.SBasic:
		T.cBasic.fontUrl = url
	case objects.SOver:
		T.cOver.fontUrl = url
	case objects.SClick:
		T.cClick.fontUrl = url
	default:
		return errors.New("Status not available")
	}
	return nil
}

// SetAction define action when the element is click
func (T *Text) SetAction(f func(...interface{})string, d []interface{}) {
	T.funcClick = f
	T.dataClick = d
}

func (T *Text) CopyStateToStates(stateSource uint8, stDests []uint8) error {
	var source Content{}

	switch stateSource {
	case objects.SFix:
		source = T.cFix
	case objects.SBasic:
		source = T.cBasic
	case objects.SOver:
		source = T.cOver
	case objects.SClick:
		source = T.cClick
	default:
		return errors.New("Status not available")
	}

	for _, v := range stDests {
		switch v {
		case objects.SFix:
			copy(T.cFix, source)
		case objects.SBasic:
			copy(T.cBasic, source)
		case objects.SOver:
			copy(T.cOver, source)
		case objects.SClick:
			copy(T.cClick, source)
		default:
			return errors.New("Status to dest copy not available")
		}
	}
	return nil
}

// GetSize provide size object (on the screen)
func (T Text) GetSize() (*objects.Size, error) {
	var s *objets.Size

	switch T.status {
	case objects.SFix:
		s = T.cFix.sizeSDL
	case objects.SBasic:
		s = T.cBasic.sizeSDL
	case objects.SOver:
		s = T.cOver.sizeSDL
	case objects.SClick:
		s = T.cClick.sizeSDL
	}

	if s == nil {
		return nil, errors.New("Not size define on the status text")
	}
	return s, nil
}

// GetPosiion provide position object
func (T Text) GetPosition() (*objects.Position, error) {
	var p *objets.Position

	switch T.status {
	case objects.SFix:
		p = T.cFix.position
	case objects.SBasic:
		p = T.cBasic.position
	case objects.SOver:
		p = T.cOver.position
	case objects.SClick:
		p = T.cClick.position
	}

	if p == nil {
		return nil, errors.New("Not position define on the status text")
	}
	return p, nil
}

// GetColor provide color object
func (T Text) GetColor() (*objects.Color, error) {
	var c *objets.Position

	switch T.status {
	case objects.SFix:
		c = T.cFix.color
	case objects.SBasic:
		c = T.cBasic.color
	case objects.SOver:
		c = T.cOver.color
	case objects.SClick:
		c = T.cClick.color
	}

	if c == nil {
		return nil, errors.New("Not color define on the status text")
	}
	return c, nil
}

/*
** Interface objects functions
 */

// IsInit return status initialize
func (T Text) IsInit() bool {
	return T.initialized
}


func (T *Text) Init(r *sdl.Renderer) error {
	if T.status == objects.SFix {
		if err := T.cFix.checkContent(); err != nil {
			return err
		}
		if err := T.cFix.initContent(r); err != nil {
			return err
		}
	} else {
		if err := T.cBasic.checkContent(); err != nil {
			return err
		}
		if err := T.cBasic.initContent(r); err != nil {
			return err
		}

		if err := T.cOver.checkContent(); err != nil {
			return err
		}
		if err := T.cOver.initContent(r); err != nil {
			return err
		}

		if err := T.cClick.checkContent(); err != nil {
			return err
		}
		if err := T.cClick.initContent(r); err != nil {
			return err
		}
	}

	T.initialized = true
	return nil
}

func (T *Text) Close() error {
	if err := T.cFix.closeContent(); err != nil {
		return err
	}
	if err := T.cBasic.closeContent(); err != nil {
		return err
	}
	if err := T.cOver.closeContent(); err != nil {
		return err
	}
	if err := T.cClick.closeContent(); err != nil {
		return err
	}

	T.initialized = false
	return nil
}

func (T *Text) GetStatus() uint8 {
	return T.status
}

func (T *Text) IsOver(x, y int32) bool {
	if T.status == objects.SFix {
		return false
	}

	switch T.status {
	case objects.SBasic:
		return T.cBasic.isOverContent(x, y)
	case objects.SOver:
		return T.cOver.isOverContent(x, y)
	case objects.SClicl:
		return T.cClick.isOverContent(x, y)
	}
	return false
}

func (T *Text) Click() {
	T.SetStatus(objects.SClick)
	T.funcClick(T.dataClick)
}

func (T *Text) SetStatus(s uint8) {
	if T.status != objects.SFix {
		T.status = s
	}
}

// Draw the object block.
func (T *Text) Draw(r *sdl.Renderer, wg *sync.WaitGroup) error {
	if r == nil {
		return errors.New("Can't draw text because renderer is nil")
	}
	if wg == nil {
		return errors.New("Can't draw text because sync WaitGroup not define")
	}
	if T.initialized == false {
		return errors.New("Can't draw text object is not initialized")
	}

	wg.Add(1)
	defer wg.Done()

	sdl.Do(func() {
		err := D.renderer.SetDrawColor(B.color.Red, B.color.Gree, B.color.Blue, B.color.Opacity)
		if err != nil {
			panic(err)
		}

		switch T.status {
		case objects.SFix:
			err = r.Copy(T.cFix.texture, nil, &T.cFix.rect)
		case objects.SBasic:
			err = r.Copy(T.cBasic.texture, nil, &T.cBasic.rect)
		case objects.SOver:
			err = r.Copy(T.cOver.texture, nil, &T.cOver.rect)
		case objects.SClick:
			err = r.Copy(T.cClick.texture, nil, &T.cClick.rect)
		default:
			err = errors.New("text type no recognized")
		}

		if err != nil {
			panic(err)
		}
	})
	return nil
}

/*
** Private function Text objects
 */
// checkContent and return err with the raison.
func (C Content) checkContent() error {
	if C.txt == "" {
		return false, errors.New("Txt not define")
	}
	if C.size <= 0 {
		return false, errors.New("Size txt can't be equal zero")
	}
	if C.position == nil {
		return false, errors.New("Position not define")
	}
	if C.color == nil {
		return false, errors.New("Color not define")
	}
	if C.fontUrl == nil {
		return false, errors.New("Not font ttf provide")
	}

	return nil
}

// initContent initialized sdl object
func (C *Content) initContent(r *sdl.Renderer) error {
	font, err := ttf.OpenFont(C.fontUrl, C.size)
	if err != nil {
		panic(err)
	}
	defer font.Close()

	color := sdl.Color{
		R: C.color.Red,
		G: C.color.Green,
		B: C.color.Blue,
		A: C.color.Opacity,
	}
	surface, err := font.RenderUTF8_Solid(C.txt, color)
	if err != nil {
		panic(err)
	}
	defer surface.Free()

	C.sizeSDL = new(objects.Size)
	C.sizeSDL.SetSize(surface.W, surface.H)

	C.rect.X = C.position.X - (C.sizeSDL.W / 2)
	C.rect.Y = C.position.Y - (C.sizeSDL.H / 2)
	C.rect.W = C.sizeSDL.W
	C.rect.H = C.sizeSDL.H

	C.texture, err := D.renderer.CreateTextureFromSurface(surface)
	if err != nil {
		panic(err)
	}
	return nil
}

// closeContent close sdl objects
func (C *Content) closeContent() error {
	if C.texture == nil {
		return errors.New("Can't destroy texture because it nul")
	}
	if C.texture != nil {
		C.texture.Destroy()
	}
	return nil
}

// isOverContent
func (C Content) isOverContent(x, y int32) bool {
	if C.position || C.sizeSDL {
		return false
	}
	if x > C.position.X && x < C.position.X+sizeSDL.W {
		if y > C.position.Y && y < C.position.Y+sizeSDL.H {
			return true
		}
	}
	return false
}
