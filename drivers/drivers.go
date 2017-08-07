package drivers

import (
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

// VSDL veandco sdl drivers
type VSDL struct {
	window   *sdl.Window
	renderer *sdl.Renderer
}

// Destroy close sdl objects correctly
func (V *VSDL) Destroy() {
	if V.renderer != nil {
		V.renderer.Destroy()
	}

	if V.window != nil {
		V.window.Destroy()
	}
	ttf.Quit()
	mix.Quit()
	sdl.Quit()
}

// Init create sdl window and the renderer objects
func Init(width, height int, title string) (VSDL, error) {
	var V VSDL
	var err error

	sdl.Init(sdl.INIT_EVERYTHING)

	V.window, err = sdl.CreateWindow(title, sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, width, height, sdl.WINDOW_OPENGL)
	if err != nil {
		return V, err
	}

	V.renderer, err = sdl.CreateRenderer(V.window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return V, err
	}

	if err := mix.Init(sdl.INIT_AUDIO); err != nil {
		return V, err
	}

	mix.AllocateChannels(255)
	if err := mix.OpenAudio(48000, mix.DEFAULT_FORMAT, mix.DEFAULT_CHANNELS, mix.DEFAULT_CHUNKSIZE); err != nil {
		return V, err
	}

	if err := ttf.Init(); err != nil {
		return V, err
	}

	V.renderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND)
	sdl.StopTextInput()
	return V, nil
}

// GetRenderer : return renderer SDL2
func (V *VSDL) GetRenderer() *sdl.Renderer {
	return V.renderer
}
