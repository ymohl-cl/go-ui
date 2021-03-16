package loader

import (
	"errors"
	"sync"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/ymohl-cl/game-builder/objects"
)

/*
** interface functions
 */

// Build describe the scene with objects needest
func (l *DefaultLoader) Build() error {
	var err error

	l.layers = make(map[uint8][]objects.Object)

	if err = l.addTxt(); err != nil {
		return err
	}
	if err = l.addBlockLoading(); err != nil {
		return err
	}
	return nil
}

// Init the scene
func (l *DefaultLoader) Init() error {
	if l.renderer == nil {
		return errors.New(objects.ErrorRenderer)
	}

	if l.layers == nil {
		return errors.New(errorLayers)
	}
	if l.loadBlock == nil {
		return errors.New(errorLayers)
	}
	l.initialized = true
	return nil
}

// IsInit return status initialize
func (l *DefaultLoader) IsInit() bool {
	return l.initialized
}

// Run the scene
func (l *DefaultLoader) Run() error {
	l.addLoadingBar()
	return nil
}

// Stop scene
func (l *DefaultLoader) Stop() {
	l.resetLoadingBlock()
}

// Close the scene
func (l *DefaultLoader) Close() error {
	var err error

	l.initialized = false

	l.m.Lock()
	defer l.m.Unlock()
	if err = objects.Closer(l.layers); err != nil {
		return err
	}

	l.layers = nil
	return nil
}

// GetLayers provide all scene's objects to draw them
func (l *DefaultLoader) GetLayers() (map[uint8][]objects.Object, *sync.Mutex) {
	return l.layers, l.m
}

// KeyboardEvent provide key down to the scene
func (l *DefaultLoader) KeyboardEvent(keyboard *sdl.KeyboardEvent) {
	return
}

// SetSwitcher can be call to change scene with index scene and flag closer
func (l *DefaultLoader) SetSwitcher(f func(uint8, bool) error) {
	l.switcher = f
}

// Update : called on each frame
func (l *DefaultLoader) Update() {
	l.addLoadingBar()
	return
}

// SetCloser : allow quit the application
func (l *DefaultLoader) SetCloser(f func()) {
	// nothing to do
}
