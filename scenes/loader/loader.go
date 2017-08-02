package loader

import (
	"errors"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/ymohl-cl/game-builder/database"
	"github.com/ymohl-cl/game-builder/objects"
)

// Load is a scene which used when build other scene
type Load struct {
	/* infos scene */
	initialized bool
	closer      chan (uint8)

	/* objects by layers */
	layers map[uint8][]objects.Object

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
	return &l, nil
}
