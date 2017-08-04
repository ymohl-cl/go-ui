package gomoku

import (
	"errors"
	"sync"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/ymohl-cl/game-builder/objects"
	"github.com/ymohl-cl/game-builder/scenes"
)

/*
** interface functions
 */

// Build describe the scene with objects needest
func (G *Gomoku) Build() error {
	var err error

	if err = G.addMusic(); err != nil {
		return err
	}
	if err = G.addBackground(); err != nil {
		return err
	}

	return nil
}

// Init the scene
func (G *Gomoku) Init() error {
	var err error

	if G.renderer == nil {
		return errors.New(objects.ErrorRenderer)
	}

	if err = G.Build(); err != nil {
		return err
	}

	if G.layers == nil {
		return errors.New(scenes.ErrorLayers)
	}

	if G.music == nil {
		return errors.New(scenes.ErrorMissing)
	}

	G.initialized = true
	return nil
}

// IsInit return status initialize
func (G *Gomoku) IsInit() bool {
	return G.initialized
}

// Run the scene
func (G *Gomoku) Run() error {
	//	var err error
	var wg sync.WaitGroup

	if ok := G.music.IsInit(); ok {
		wg.Add(1)
		go G.music.Play(&wg, G.renderer)
		wg.Wait()
	}
	return nil
}

// Close the scene
func (G *Gomoku) Close() error {
	var err error

	G.initialized = false
	if ok := G.music.IsInit(); ok {
		if err = G.music.Close(); err != nil {
			return err
		}
	}

	return nil
}

// GetLayers provide all scene's objects to draw them
func (G *Gomoku) GetLayers() map[uint8][]objects.Object {
	return G.layers
}

// KeyDownEvent provide key down to the scene
func (G *Gomoku) KeyDownEvent(keyDown *sdl.KeyDownEvent) {
	return
}
