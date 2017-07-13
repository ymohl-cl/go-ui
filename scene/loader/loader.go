package loader

import (
	"errors"
	"sync"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/ymohl-cl/game-builder/objects"
	"github.com/ymohl-cl/game-builder/objects/block"
)

const (
	// order layers of scene
	//layerBackground = iota
	layerText = iota
	layerLoadingBar

	// text to the scene loader
	font                = "/github.com/ymohl-cl/game-builder/Ressources/stocky.ttf"
	mainTxt             = "LOADING ..."
	txtBottomToLoad     = "this app use gogame-builder github.com/ymohl-cl/game-builder"
	sizeMainTxt         = 78
	sizeTxtBottomToLoad = 20

	// color block to loading bar
	colorRed     = 52
	colorGreen   = 201
	colorBlue    = 36
	colorOpacity = 255

	// position
	originX = 0
	originY = 1

	// errors message
	errorLayers = "layers not define on the default loader"
)

// DefaultLoader is the default loader if it don't provided
type DefaultLoader struct {
	/* infos scene */
	initialized bool
	refresh     bool
	switcher    func(uint8, bool) error

	/* objects by layers */
	m         *sync.Mutex
	layers    map[uint8][]objects.Object
	loadBlock *block.Block

	/* size to load objects */
	widthScreen  int32
	heightScreen int32
	widthBlock   int32
	heightBlock  int32

	/* sdl ressources */
	renderer *sdl.Renderer
}

/*
** constructor
 */

// New provide a new object
func New(r *sdl.Renderer, width, height int32) (*DefaultLoader, error) {
	if r == nil {
		return nil, errors.New(objects.ErrorRenderer)
	}

	l := DefaultLoader{renderer: r}
	l.layers = make(map[uint8][]objects.Object)
	l.m = new(sync.Mutex)
	l.widthScreen = width
	l.heightScreen = height
	l.widthBlock = (width / 100) * 2
	l.heightBlock = (height / 100) * 3
	return &l, nil
}
