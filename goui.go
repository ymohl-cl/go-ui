package goui

import (
	"sync"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

const (
	frame = 60
)

// GoUI make the relation between sdl2 driver and the application service. Provide an easy way to setup application's ui or game'ui
type GoUI interface {
	Close() error
	Renderer() Renderer
	Run(scene Scene) error
	loop(scene Scene) error
}

type goUI struct {
	renderer Renderer
}

// New configure a new gamebuilder instance and setup a sdl driver ready to use
func New(c ConfigUI) (GoUI, error) {
	var err error
	var ui goUI

	sdl.Init(sdl.INIT_EVERYTHING)
	// renderer configuration
	if ui.renderer, err = NewRenderer(c); err != nil {
		return nil, err
	}
	// audio configuration
	if err = mix.Init(sdl.INIT_AUDIO); err != nil {
		return nil, err
	}
	mix.AllocateChannels(255)
	if err = mix.OpenAudio(48000,
		mix.DEFAULT_FORMAT,
		mix.DEFAULT_CHANNELS,
		mix.DEFAULT_CHUNKSIZE); err != nil {
		return nil, err
	}
	// font type configuration
	if err = ttf.Init(); err != nil {
		return nil, err
	}
	sdl.StopTextInput()
	return &ui, nil
}

// Close the sdl resources
func (ui *goUI) Close() error {
	ui.renderer.Close()
	ttf.Quit()
	mix.Quit()
	img.Quit()
	sdl.Quit()
	return nil
}

// Renderer getter
func (ui goUI) Renderer() Renderer {
	return ui.renderer
}

// Run the sdl loop and start with the scene targeted by sceneIndex
func (ui *goUI) Run(scene Scene) error {
	// SDL.Main allow use sdl.Do() to queue sdl instructions.
	c := make(chan error, 1)
	sdl.Main(func() {
		err := ui.loop(scene)
		c <- err
		return
	})
	return <-c
}

// loop app/game to cach event and draw scene / frame
func (ui *goUI) loop(s Scene) error {
	var err error
	var wg sync.WaitGroup

	for ok := true; ok; {
		sdl.Do(func() {
			for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
				switch e.(type) {
				case *sdl.QuitEvent:
					ok = false
					break
				default:
					if err = s.NewEvent(e); err != nil {
						ok = false
					}
					break
				}
			}
		})
		if err != nil || !ok {
			break
		}
		// clear image
		wg.Add(1)
		sdl.Do(func() {
			defer wg.Done()
			if err = ui.renderer.Clear(); err != nil {
				ok = false
			}
		})
		if err != nil {
			break
		}
		wg.Wait()
		// prepare new image
		s.Render(ui.renderer.Driver())
		// draw new image
		wg.Add(1)
		sdl.Do(func() {
			defer wg.Done()
			ui.renderer.Draw()
		})
		wg.Wait()
		// wait frame
		sdl.Do(func() {
			sdl.Delay(1000 / frame)
		})
	}
	if err != nil {
		return err
	}
	return nil
}
