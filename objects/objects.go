package objects

import (
	"sync"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	// list of status
	SFix   = 0
	SBasic = 1
	SOver  = 2
	SClick = 3

	path             = "Ressources/"
	ErrorRenderer    = "sdl.Renderer is nil"
	ErrorNotInit     = "object not initialized"
	ErrorColor       = "object has no defined color"
	ErrorTxt         = "object has no defined text"
	ErrorPosition    = "object has no defined position"
	ErrorSize        = "object has no defined size"
	ErrorObjectStyle = "object does not know this style"
	ErrorTargetURL   = "object does not know url target"
	ErrorStatus      = "object does not know status"
	ErrorEmpty       = "object is empty"
	ErrorData        = "data is empty"
)

type Object interface {
	// Init initialize object with SDL
	Init(r *sdl.Renderer) error
	// IsInit provide a bool value to inisialization value
	IsInit() bool
	// Close close Sdl elements
	Close() error
	// GetStatus object
	GetStatus() uint8
	// Define if object is over by mouse x-y
	IsOver(int32, int32) bool
	// Click on object when is it over
	Click()
	// Define new status
	SetStatus(uint8)
	// UpdatePosition change position to define a movement
	UpdatePosition(x, y int32)

	// MoveTo(x, y) add x and y to the position
	MoveTo(x, y int32)

	// GetPosition to get x and y coordinates
	GetPosition() (int32, int32)

	// GetSize return width and height of object
	//	GetSize() (int32, int32)

	// Draw object with sdl
	Draw(*sync.WaitGroup, *sdl.Renderer)
}
