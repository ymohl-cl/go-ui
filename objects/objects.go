package objects

import (
	"sync"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	// list of status
	SFix   = 0
	Sbasic = 1
	SOver  = 2
	SClick = 3

	path = "Ressources/"
)

type Object interface {
	// IsInit provide a bool value to inisialization value
	IsInit() bool
	// Init initialize object with SDL
	Init(r *sdl.Renderer) error
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
	// Draw object with sdl
	Draw(*sync.WaitGroup) error
}
