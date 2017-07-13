package scripter

import (
	"sync"

	"github.com/ymohl-cl/game-builder/scene"
)

const (
	frame                      = 60
	layerLoading               = 0
	errorIndexLoader           = "index 0 is reserved to loader scene"
	errorIndexUsed             = "index to add new scene is already used"
	errorChangeScriptOnRunning = "script can't be modified when is running"
	errorScriptNoRunning       = "script don't be running"
	errorIndexUndefined        = "index scene is undefined"
)

// Script object : manage all scenes
type Script struct {
	list    map[uint8]scene.Scene
	running bool

	// lock state current
	m       *sync.Mutex
	current uint8
}

// New provide a script object
func New() *Script {
	s := Script{}

	s.list = make(map[uint8]scene.Scene)
	s.m = new(sync.Mutex)
	return &s
}
