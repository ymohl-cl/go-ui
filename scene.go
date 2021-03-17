package gamebuilder

import (
	"sync"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/ymohl-cl/go-ui/objects"
)

// Scene is a interface and define the design model to your scenes.
type Scene interface {
	// Build the scene
	Build() error
	// Init the scene. Create static objects. Data is provide if you need.
	Init() error
	// IsInit return status initialize
	IsInit() bool
	// Run the scene
	Run() error
	// Stop the scene, it's possible to Run later
	Stop()
	// Close the scene at the end game
	Close() error
	// GetLayers get objects list by layers order
	GetLayers() (map[uint8][]objects.Object, *sync.Mutex)
	// KeyboardEvent provide key down to the scene
	KeyboardEvent(*sdl.KeyboardEvent)
	// SetSwitcher can be called to change scene with index scene on
	SetSwitcher(func(indexScene string) error)
	// SetCloser call the closer scene
	SetCloser(func(indexScene string) error)
	// Update : called on each frame
	Update()
}
