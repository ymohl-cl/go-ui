package goui

import (
	"github.com/veandco/go-sdl2/sdl"
)

// Renderer type to manage the drawing
type Renderer interface {
	Close()
	Clear() error
	Draw()
	Driver() *sdl.Renderer
}

type renderer struct {
	driver *sdl.Renderer
	window *sdl.Window
}

// NewRenderer sdl
func NewRenderer(c ConfigUI) (Renderer, error) {
	var r renderer
	var err error

	if r.window, err = sdl.CreateWindow(c.Window.Title,
		sdl.WINDOWPOS_CENTERED,
		sdl.WINDOWPOS_CENTERED,
		c.Window.Width,
		c.Window.Height,
		sdl.WINDOW_OPENGL); err != nil {
		return nil, err
	}
	if r.driver, err = sdl.CreateRenderer(r.window, -1, sdl.RENDERER_ACCELERATED); err != nil {
		return nil, err
	}
	r.driver.SetDrawBlendMode(sdl.BLENDMODE_BLEND)

	return &r, nil
}

// Close sdl renderer resources
func (r *renderer) Close() {
	if r.driver != nil {
		r.driver.Destroy()
	}
	if r.window != nil {
		r.window.Destroy()
	}
}

// Clear image before write a new
func (r *renderer) Clear() error {
	var err error

	if err = r.driver.SetDrawColor(0, 0, 0, 0); err != nil {
		return err
	}
	if err = r.driver.Clear(); err != nil {
		return err
	}
	return nil
}

// Draw the current renderer
func (r *renderer) Draw() {
	r.driver.Present()
}

// Driver getter (*sdl.Renderer)
func (r *renderer) Driver() *sdl.Renderer {
	return r.driver
}
