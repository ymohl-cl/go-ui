package menu

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
func (M *Menu) Build() error {
	var err error

	if err = M.addMusic(); err != nil {
		return err
	}
	if err = M.addBackground(); err != nil {
		return err
	}
	if err = M.addStructuresPage(); err != nil {
		return err
	}
	if err = M.addButtons(); err != nil {
		return err
	}
	if err = M.addNotice(); err != nil {
		return err
	}
	if err = M.addText(); err != nil {
		return err
	}
	if err = M.addVS(); err != nil {
		return err
	}
	if err = M.addInput(); err != nil {
		return err
	}
	if err = M.addPlayers(); err != nil {
		return err
	}
	return nil
}

// Init the scene
func (M *Menu) Init() error {
	var err error

	if M.renderer == nil {
		return errors.New(objects.ErrorRenderer)
	}
	if M.data == nil {
		return errors.New(objects.ErrorData)
	}

	if err = M.Build(); err != nil {
		return err
	}

	if M.layers == nil {
		return errors.New(scenes.ErrorLayers)
	}
	if M.input == nil {
		return errors.New(scenes.ErrorMissing)
	}
	if M.notice == nil {
		return errors.New(scenes.ErrorMissing)
	}
	if M.music == nil {
		return errors.New(scenes.ErrorMissing)
	}
	if M.vs == nil {
		return errors.New(scenes.ErrorMissing)
	}

	M.initialized = true
	return nil
}

// IsInit return status initialize
func (M Menu) IsInit() bool {
	return M.initialized
}

// Run the scene
func (M Menu) Run() error {
	var wg sync.WaitGroup

	if ok := M.music.IsInit(); ok {
		wg.Add(1)
		go M.music.Play(&wg, M.renderer)
		wg.Wait()
	}
	return nil
}

// Close the scene
func (M *Menu) Close() error {
	var err error

	M.initialized = false
	if ok := M.music.IsInit(); ok {
		if err = M.music.Close(); err != nil {
			return err
		}
	}

	return nil
}

// GetLayers provide all scene's objects to draw them
func (M Menu) GetLayers() map[uint8][]objects.Object {
	return M.layers
}

// KeyDownEvent provide key down to the scene
func (M *Menu) KeyDownEvent(keyDown *sdl.KeyDownEvent) {
	var err error

	if err = M.input.SetNewRune(keyDown.Keysym, M.renderer); err != nil {
		go M.setNotice(err.Error())
	}
}
