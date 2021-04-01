package widget

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	RectStyleFill RectStyle = iota
	RectStyleBorder
)

type RectStyle uint8

type Rect struct {
	widget
	style RectStyle
}

func NewRect(w, h int32) *Rect {
	r := Rect{}
	r.SetSize(w, h)

	return &r
}

func (r *Rect) Close() {}

func (r *Rect) SetStyle(s RectStyle) {
	r.style = s
}

func (r *Rect) Render(renderer *sdl.Renderer) error {
	var err error

	c := r.Color(r.state)
	if err = renderer.SetDrawColor(c.Red, c.Green, c.Blue, c.Alpha); err != nil {
		return err
	}

	sdlRect := sdl.Rect{
		X: r.position.X,
		Y: r.position.Y,
		W: r.block.width,
		H: r.block.height,
	}
	switch r.style {
	case RectStyleFill:
		err = renderer.FillRect(&sdlRect)
	case RectStyleBorder:
		err = renderer.DrawRect(&sdlRect)
	}
	if err != nil {
		return err
	}
	return nil
}
