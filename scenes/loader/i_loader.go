package loader

import (
	"errors"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/ymohl-cl/game-builder/objects"
	"github.com/ymohl-cl/game-builder/scenes"
)

/*
** interface functions
 */

// Build describe the scene with objects needest
func (L *Load) Build() error {
	return nil
}

// Init the scene
func (L *Load) Init() error {
	var err error

	if L.renderer == nil {
		return errors.New(objects.ErrorRenderer)
	}

	if err = L.Build(); err != nil {
		return err
	}

	if L.layers == nil {
		return errors.New(scenes.ErrorLayers)
	}

	L.initialized = true
	return nil
}

// IsInit return status initialize
func (L Load) IsInit() bool {
	return L.initialized
}

// Run the scene
func (L Load) Run() error {
	/*	var wg sync.WaitGroup

		if ok := L.music.IsInit(); ok {
			wg.Add(1)
			go L.music.Draw(&wg, L.renderer)
			wg.Wait()
		}*/
	return nil
}

// Close the scene
func (L *Load) Close() error {
	L.initialized = false
	return nil
}

// GetLayers provide all scene's objects to draw them
func (L Load) GetLayers() map[uint8][]objects.Object {
	return L.layers
}

// KeyDownEvent provide key down to the scene
func (L *Load) KeyDownEvent(keyDown *sdl.KeyDownEvent) {
	return
}
