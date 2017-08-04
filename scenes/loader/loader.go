package loader

import (
	"errors"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/ymohl-cl/game-builder/audio"
	"github.com/ymohl-cl/game-builder/database"
	"github.com/ymohl-cl/game-builder/objects"
	"github.com/ymohl-cl/game-builder/objects/block"
)

const (
	// order layers of scene
	layerBackground = 0
	layerStructure  = 1
	layerText       = 2
	layerLoadingBar = 3
)

// Load is a scene which used when build other scene
type Load struct {
	/* infos scene */
	initialized bool
	closer      chan (bool)
	refresh     bool

	/* objects by layers */
	layers        map[uint8][]objects.Object
	music         *audio.Audio
	lastLoadBlock *block.Block

	/* specific objects */

	/* sdl ressources */
	renderer *sdl.Renderer
}

/*
** constructor
 */

// New provide a new object
func New(d *database.Data, r *sdl.Renderer) (*Load, error) {
	if r == nil {
		return nil, errors.New(objects.ErrorRenderer)
	}

	l := Load{renderer: r}
	l.layers = make(map[uint8][]objects.Object)
	l.closer = make(chan (bool))
	return &l, nil
}
