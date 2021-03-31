package widget

import (
	"fmt"

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
	t.block.width = int32(w)
	t.block.height = int32(h)

	t.state = StateBase
	return &t, nil
}

func (t *Text) Close() {
	t.font.Close()
}

func (t *Text) Update(txt string) {
	t.str = txt
}

func (t *Text) Render(r *sdl.Renderer) {
	var surface *sdl.Surface
	var texture *sdl.Texture
	var err error

	myColor := t.Color(t.state)
	color := sdl.Color{
		R: myColor.Red,
		G: myColor.Green,
		B: myColor.Blue,
		A: myColor.Alpha,
	}

	var str string
	if str, err = t.strRenderable(); err != nil {
		fmt.Printf("error occured to return the STR: %s", err.Error())
		return
	}
	if surface, err = t.font.driver.RenderUTF8Blended(str, color); err != nil {
		fmt.Printf("error occured while renderUTF8Blended: %s", err.Error())
		return
	}
	defer surface.Free()
	if texture, err = r.CreateTextureFromSurface(surface); err != nil {
		fmt.Printf("error occured while create texture from surface: %s", err.Error())
		return
	}
	defer texture.Destroy()

	var w, h int
	if w, h, err = t.font.driver.SizeUTF8(str); err != nil {
		fmt.Printf("error occured while get size utf8: %s", err.Error())
		return
	}
	if err = r.Copy(texture, nil, &sdl.Rect{
		X: t.position.X,
		Y: t.position.Y,
		W: int32(w),
		H: int32(h),
	}); err != nil {
		fmt.Printf("error occured while copy sdl texture: %s", err.Error())
		return
	}
	return
}

func (t *Text) strRenderable() (string, error) {
	var w int
	var err error
	str := t.str

	for w, _, err = t.font.driver.SizeUTF8(str); err == nil && w > int(t.block.width); w, _, err = t.font.driver.SizeUTF8(str) {
		fmt.Printf("w: %d and t.block.width: %d\n", w, t.block.width)
		str = str[:len(str)-1]
	}
	if err != nil {
		return "", err
	}
	return str, nil
}
