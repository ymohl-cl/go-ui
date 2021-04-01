package widget

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type Picture struct {
	widget
	surface *sdl.Surface
}

func NewPicture(path string) (*Picture, error) {
	var err error
	var p Picture

	if p.surface, err = img.Load(path); err != nil {
		return nil, err
	}
	p.SetSize(p.surface.W, p.surface.H)
	return &p, nil
}

func (p *Picture) Close() {
	p.surface.Free()
}

func (p *Picture) Render(r *sdl.Renderer) error {
	var err error
	var texture *sdl.Texture

	c := p.Color(p.state)
	if err = r.SetDrawColor(c.Red, c.Green, c.Blue, c.Alpha); err != nil {
		return err
	}
	rect := sdl.Rect{
		X: p.position.X,
		Y: p.position.Y,
		W: p.block.width,
		H: p.block.height,
	}

	if texture, err = r.CreateTextureFromSurface(p.surface); err != nil {
		return err
	}
	defer texture.Destroy()
	if err = r.Copy(texture, nil, &rect); err != nil {
		return err
	}
	if err = r.FillRect(&rect); err != nil {
		return err
	}

	return nil
}

func (p *Picture) SetWidthRatioHeight(width int32) {
	pourcent := float32(width) / float32(p.surface.W) * 100.0
	height := float32(p.surface.H) * (pourcent / 100.0)
	p.SetSize(width, int32(height))
}

func (p *Picture) SetHeightRatioWidth(height int32) {
	pourcent := float32(height) / float32(p.surface.H) * 100.0
	width := float32(p.surface.W) * (pourcent / 100.0)
	p.SetSize(int32(width), height)
}
