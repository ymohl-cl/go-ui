package drivers

import (
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"github.com/ymohl-cl/game-builder/conf"
)

// VSdl veandco sdl drivers
type VSdl struct {
	window   *sdl.Window
	renderer *sdl.Renderer
}

// Destroy close sdl objects correctly
func (V *VSdl) Destroy() {
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
func Init() (VSdl, error) {
	var V VSdl
	var err error

	sdl.Init(sdl.INIT_EVERYTHING)

	V.window, err = sdl.CreateWindow(conf.Title, sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, conf.WindowWidth, conf.WindowHeight, sdl.WINDOW_OPENGL)
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

func (V *VSdl) GetRenderer() *sdl.Renderer {
	return V.renderer
}
