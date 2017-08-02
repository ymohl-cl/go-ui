package loader

import (
	"errors"
	"sync"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/ymohl-cl/game-builder/audio"
	"github.com/ymohl-cl/game-builder/database"
	"github.com/ymohl-cl/game-builder/objects"
)

type Load struct {
	layers      map[uint8][]objects.Object
	initialized bool

	// content Load scene
	music *audio.Audio

	/* sdl objects */
	renderer *sdl.Renderer
}

/* Functions */
func (L *Load) Init(d *database.Data, r *sdl.Renderer) error {
	var err error

	if r == nil {
		return errors.New(objects.ErrorRenderer)
	}
	L.renderer = r
	L.layers = make(map[uint8][]objects.Object)

	if err = L.Build(); err != nil {
		return err
	}

	if err = L.check(); err != nil {
		return err
	}
	L.initialized = true
	return nil
}

func (L Load) Run() error {
	var wg sync.WaitGroup

	if ok := L.music.IsInit(); ok {
		wg.Add(1)
		go L.music.Draw(&wg, L.renderer)
		wg.Wait()
	}
	return nil
}

func (L *Load) Close() error {
	var err error

	L.initialized = false
	return nil
}

func (L Load) GetLayers() map[uint8][]objects.Object {
	return L.layers
}

func (L Load) check() error {

}
