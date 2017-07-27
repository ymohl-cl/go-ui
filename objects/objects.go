package objects

import (
	"sync"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	// SFix to StatusFix (can't be overable or clickable)
	SFix = 0
	// SBasic to StatusBasic (default status)
	SBasic = 1
	// SOver to statusOver
	SOver = 2
	// SClick to statusClick
	SClick = 3

	path = "Ressources/"
	// ErrorRenderer (nil pointer)
	ErrorRenderer = "sdl.Renderer is nil"
	// ErrorNotInit (object not initialized)
	ErrorNotInit = "object not initialized"
	// ErrorTargetURL (url parameters not known)
	ErrorTargetURL = "object does not know url target parameters"
	// ErrorStatus (status not known)
	ErrorStatus = "object does not know status"
	// ErrorStyle (style not known)
	ErrorStyle = "object does not know style"
	// ErrorBuild (object empty)
	ErrorBuild = "object don't build with new constructor"

/*	ErrorColor       = "object has no defined color"
	ErrorTxt         = "object has no defined text"
	ErrorPosition    = "object has no defined position"
	ErrorSize        = "object has no defined size"
	ErrorObjectStyle = "object does not know this style"
	ErrorData        = "data is empty"*/
)

// Object interface it's a model to the game builder.
type Object interface {
	// Init object to draw it. If error occurred, object can't be drawn
	Init(r *sdl.Renderer) error
	// IsInit return status initialize
	IsInit() bool
	// Close sdl ressources needest to object
	Close() error
	// SetAction to get it on click button
	SetAction(f func(...interface{}), d ...interface{})
	// SetStatus change object's status
	SetStatus(uint8)
	// MoveTo by increment position with x and y parameters
	MoveTo(x, y int32)
	// UpdatePosition object
	UpdatePosition(x, y int32)
	// UpdateSize object
	UpdateSize(w, h int32)
	// UpdateColor object
	UpdateColor(r, g, b, a uint8)
	// IsOver define if object and position parameters matches
	IsOver(int32, int32) bool
	// GetStatus object
	GetStatus() uint8
	// GetPosition object (x, y)
	GetPosition() (int32, int32)
	// GetSize object (width, height)
	GetSize() (int32, int32)
	// Click define a click on object
	Click()
	// Draw object
	Draw(*sync.WaitGroup, *sdl.Renderer)
}
