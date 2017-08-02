package menu

import (
	"errors"

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

// Menu is a scene
type Menu struct {
	/* infos scene */
	initialized bool
	closer      chan (uint8)

	/* objects by layers */
	layers map[uint8][]objects.Object

	/* specific objects */
	input  *input.Input
	notice *text.Text
	music  *audio.Audio
	vs     *text.Text

	/* sdl ressources */
	renderer *sdl.Renderer
	/* data game */
	data *database.Data
}

/*
** constructor
 */

// New provide a new object
func New(d *database.Data, r *sdl.Renderer) (*Menu, error) {
	if r == nil {
		return nil, errors.New(objects.ErrorRenderer)
	}
	if d == nil {
		return nil, errors.New(objects.ErrorData)
	}

	m := Menu{renderer: r, data: d}
	m.layers = make(map[uint8][]objects.Object)
	return &m, nil
}
