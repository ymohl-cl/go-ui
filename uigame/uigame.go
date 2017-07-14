package uigame

import (
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"github.com/ymohl-cl/game-builder/scenes/sinfos"
)

type Drivers struct {
	window   *sdl.Window
	surface  *sdl.Surface
	renderer *sdl.Renderer
}

// Destroy close sdl correctly
func (D *Drivers) Destroy() {
	if D.renderer != nil {
		D.renderer.Destroy()
	}

	if D.window != nil {
		D.window.Destroy()
	}
	ttf.Quit()
	mix.Quit()
	sdl.Quit()
}

// Init create sdl window and the renderer
func Init() (Drivers, error) {
	var D Drivers
	var err error

	sdl.Init(sdl.INIT_EVERYTHING)

	D.window, err = sdl.CreateWindow("Gomoku", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, sinfos.ScreenWidth, sinfos.ScreenHeight, sdl.WINDOW_OPENGL)
	if err != nil {
		return D, err
	}

	/*	D.surface, err = D.window.GetSurface()
		if err != nil {
			return D, err
		}

	D.renderer, err = sdl.CreateSoftwareRenderer(D.surface)*/
	D.renderer, err = sdl.CreateRenderer(D.window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return D, err
	}

	if err := mix.Init(sdl.INIT_AUDIO); err != nil {
		return D, err
	}

	if err := mix.OpenAudio(48000, mix.DEFAULT_FORMAT, mix.DEFAULT_CHANNELS, mix.DEFAULT_CHUNKSIZE); err != nil {
		return D, err
	}

	if err := ttf.Init(); err != nil {
		return D, err
	}

	D.window.SetTitle("Gomoku")
	D.renderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND)
	sdl.StopTextInput()
	return D, nil
}
