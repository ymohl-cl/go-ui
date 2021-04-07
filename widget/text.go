package widget

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type Font struct {
	driver *ttf.Font
}

func NewFont(file string, size int) (Font, error) {
	var err error
	var f Font

	if f.driver, err = ttf.OpenFont(file, size); err != nil {
		return Font{}, err
	}
	return f, nil
}

func (f *Font) Close() {
	if f.driver != nil {
		f.driver.Close()
	}
}

type Text struct {
	widget
	str  string
	font Font
}

func NewText(txt string, font Font) (*Text, error) {
	var err error

	t := Text{
		str:  txt,
		font: font,
	}
	var w, h int
	if w, h, err = t.font.driver.SizeUTF8(t.str); err != nil {
		return nil, err
	}

	t.SetSize(int32(w), int32(h))
	return &t, nil
}

func (t *Text) Close() {
	t.font.Close()
}

func (t *Text) Set(txt string) {
	t.str = txt
}

func (t *Text) Render(r *sdl.Renderer) error {
	var surface *sdl.Surface
	var texture *sdl.Texture
	var err error

	// skip empty text
	if len(t.str) == 0 {
		return nil
	}

	myColor := t.Color(t.state)
	color := sdl.Color{
		R: myColor.Red,
		G: myColor.Green,
		B: myColor.Blue,
		A: myColor.Alpha,
	}

	var str string
	if str, err = t.strRenderable(); err != nil {
		return err
	}
	if surface, err = t.font.driver.RenderUTF8Blended(str, color); err != nil {
		return err
	}
	defer surface.Free()
	if texture, err = r.CreateTextureFromSurface(surface); err != nil {
		return err
	}
	defer texture.Destroy()

	var w, h int
	if w, h, err = t.font.driver.SizeUTF8(str); err != nil {
		return err
	}
	if err = r.Copy(texture, nil, &sdl.Rect{
		X: t.position.X,
		Y: t.position.Y,
		W: int32(w),
		H: int32(h),
	}); err != nil {
		return err
	}
	return nil
}

// SizeSTR getter
func (t *Text) SizeSTR() (int32, int32, error) {
	w, h, err := t.font.driver.SizeUTF8(t.str)
	if err != nil {
		return 0, 0, err
	}
	return int32(w), int32(h), nil
}

func (t *Text) strRenderable() (string, error) {
	var w int
	var err error
	str := t.str

	for w, _, err = t.font.driver.SizeUTF8(str); err == nil && w > int(t.block.width); w, _, err = t.font.driver.SizeUTF8(str) {
		str = str[:len(str)-1]
	}
	if err != nil {
		return "", err
	}
	return str, nil
}
