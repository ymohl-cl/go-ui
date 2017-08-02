package scenes

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/ymohl-cl/game-builder/objects"
)

const (
	// ErrorLayers is not allowed
	ErrorLayers = "no layers defined"
	// ErrorMissing define object missing
	ErrorMissing = "object missing on scene"
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
	// Close the scene
	Close() error
	// GetLayers get objects list by layers order
	GetLayers() map[uint8][]objects.Object
	// KeyDownEvent provide key down to the scene
	KeyDownEvent(*sdl.KeyDownEvent)
}
