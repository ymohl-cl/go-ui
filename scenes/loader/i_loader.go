package loader

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
func (L *Load) Build() error {
	var err error

	if err = L.addMusic(); err != nil {
		return err
	}
	if err = L.addBackground(); err != nil {
		return err
	}
	if err = L.addStructure(); err != nil {
		return err
	}
	if err = L.addTxt(); err != nil {
		return err
	}
	if err = L.addBlockLoading(); err != nil {
		return err
	}
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

	if L.music == nil {
		return errors.New(scenes.ErrorMissing)
	}

	L.initialized = true
	return nil
}

// IsInit return status initialize
func (L Load) IsInit() bool {
	return L.initialized
}

// Run the scene
func (L *Load) Run() error {
	//	var err error
	var wg sync.WaitGroup

	if ok := L.music.IsInit(); ok {
		wg.Add(1)
		go L.music.Play(&wg, L.renderer)
		wg.Wait()
	}
	go L.addLoadingBar()
	return nil
}

// Close the scene
func (L *Load) Close() error {
	var err error

	L.initialized = false
	L.closer <- true
	if ok := L.music.IsInit(); ok {
		if err = L.music.Close(); err != nil {
			return err
		}
	}
	L.resetLoadingBlock()
	if err = L.lastLoadBlock.Close(); err != nil {
		return err
	}

	return nil
}

// GetLayers provide all scene's objects to draw them
func (L *Load) GetLayers() map[uint8][]objects.Object {
	if L.refresh == true {
		L.resetLoadingBlock()
		L.refresh = false
	}
	return L.layers
}

// KeyDownEvent provide key down to the scene
func (L *Load) KeyDownEvent(keyDown *sdl.KeyDownEvent) {
	return
}
