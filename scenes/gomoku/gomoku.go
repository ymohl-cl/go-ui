package gomoku

import (
	"errors"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/ymohl-cl/game-builder/audio"
	"github.com/ymohl-cl/game-builder/database"
	"github.com/ymohl-cl/game-builder/objects"
)

const (
	// order layers of scene
	layerBackground = 0
)

// Gomoku is a scene which used when build other scene
type Gomoku struct {
	/* infos scene */
	initialized bool

	/* objects by layers */
	layers map[uint8][]objects.Object
	music  *audio.Audio

	/* specific objects */

	/* sdl ressources */
	renderer *sdl.Renderer
}

/*
** constructor
 */

// New provide a new object
func New(d *database.Data, r *sdl.Renderer) (*Gomoku, error) {
	if r == nil {
		return nil, errors.New(objects.ErrorRenderer)
	}

	g := Gomoku{renderer: r}
	g.layers = make(map[uint8][]objects.Object)
	return &g, nil
}
