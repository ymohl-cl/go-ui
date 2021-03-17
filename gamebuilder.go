package gamebuilder

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

// GameBuilder make the relation between sdl2 driver and the application service. Provide an easy way to setup application's ui or game'ui
type GameBuilder interface {
	Close() error
	Script() Script
	Renderer() Renderer
	Run(scene string) error
	loop() error
}

type gameBuilder struct {
	renderer Renderer
	script   Script
	running  bool
}

// New configure a new gamebuilder instance and setup a sdl driver ready to use
func New(c ConfigUI) (GameBuilder, error) {
	var err error
	var g gameBuilder

	sdl.Init(sdl.INIT_EVERYTHING)
	// renderer configuration
	if g.renderer, err = NewRenderer(c); err != nil {
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
	g.script = NewScript()
	return &g, nil
}

// Close the sdl resources
func (g *gameBuilder) Close() error {
	var err error

	g.renderer.Close()
	ttf.Quit()
	mix.Quit()
	img.Quit()
	sdl.Quit()

	if err = g.script.Close(); err != nil {
		return err
	}
	return nil
}

// Script getter
func (g gameBuilder) Script() Script {
	return g.script
}

// Renderer getter
func (g gameBuilder) Renderer() Renderer {
	return g.renderer
}

// Run the sdl loop and start with the scene targeted by sceneIndex
func (g *gameBuilder) Run(sceneIndex string) error {
	// SDL.Main allow use sdl.Do() to queue sdl instructions.
	c := make(chan error, 1)
	sdl.Main(func() {
		if err := g.script.LoadScene(sceneIndex); err != nil {
			c <- err
			return
		}

		err := g.loop()
		c <- err
		return
	})
	return <-c
}

// loop app/game to cach event and draw scene / frame
func (g *gameBuilder) loop() error {
	var err error
	var s Scene
	var wg sync.WaitGroup

	for ok := true; ok; {
		index := g.script.SceneIndex()
		if s, err = g.script.Scene(index); err != nil {
			return err
		}

		sdl.Do(func() {
			for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
				switch e.(type) {
				case *sdl.QuitEvent:
					ok = false
					break
				default:
					if err = event(e, s); err != nil {
						ok = false
					}
					break
				}
			}
		})
		if err != nil || !ok {
			break
		}
		// update scene data
		go s.Update()
		// clear image
		wg.Add(1)
		sdl.Do(func() {
			defer wg.Done()
			if err = g.renderer.Clear(); err != nil {
				ok = false
			}
		})
		if err != nil {
			break
		}
		wg.Wait()
		// prepare new image
		g.renderer.Scene(s)
		// draw new image
		wg.Add(1)
		sdl.Do(func() {
			defer wg.Done()
			g.renderer.Draw()
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
