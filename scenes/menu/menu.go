package menu

import (
	"errors"
	"sync"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/ymohl-cl/game-builder/audio"
	"github.com/ymohl-cl/game-builder/database"
	"github.com/ymohl-cl/game-builder/objects"
	"github.com/ymohl-cl/game-builder/objects/input"
	"github.com/ymohl-cl/game-builder/objects/text"
)

const (
	// order layers of scene
	layerBackground = 0
	layerStructure  = 1
	layerButton     = 2
	layerNotice     = 3
	layerText       = 4
	layerVS         = 5
	layerInput      = 6
	layerPlayers    = 7
)

type Menu struct {
	layers map[uint8][]objects.Object

	input  *input.Input
	notice *text.Text
	music  *audio.Audio
	vs     *text.Text
	data   *database.Data

	/* sdl objects */
	renderer *sdl.Renderer
}

/*
** Functions
 */

func (M *Menu) Init(d *database.Data, r *sdl.Renderer) error {
	var err error

	if r == nil {
		return errors.New(objects.ErrorRenderer)
	}
	M.renderer = r
	M.data = d

	M.layers = make(map[uint8][]objects.Object)

	if err = M.build(); err != nil {
		return err
	}

	if err = M.check(); err != nil {
		return err
	}
	return nil
}

func (M Menu) Run() error {
	var wg sync.WaitGroup

	if ok := M.music.IsInit(); ok {
		wg.Add(1)
		go M.music.Draw(&wg, M.renderer)
		wg.Wait()
	}
	return nil
}

func (M Menu) Close() error {
	var err error

	if ok := M.music.IsInit(); ok {
		if err = M.music.Close(); err != nil {
			return err
		}
	}
	return nil
}

func (M Menu) GetLayers() map[uint8][]objects.Object {
	return M.layers
}

func (M Menu) check() error {
	if M.layers == nil {
		return errors.New("Objects not define for menu scene")
	}
	if M.input == nil {
		return errors.New("Object to input not define")
	}
	if M.notice == nil {
		return errors.New("Object to notice not define")
	}

	return nil
}
