package scripter

import (
	"errors"
	"sync"

	"github.com/ymohl-cl/game-builder/data"
	"github.com/ymohl-cl/game-builder/scene"
)

const (
	loader                     = 0
	errorIndexLoader           = "index 0 is reserved to loader scene"
	errorIndexUsed             = "index to add new scene is already used"
	errorChangeScriptOnRunning = "script can't be modified when is running"
	errorScriptNoRunning       = "script don't be running"
	errorIndexUndefined        = "index scene is undefined"
)

// Script manage the specific game.
type Script struct {
	data    *data.Model
	list    map[uint8]scene.Scene
	running bool

	// lock state current
	m       *sync.Mutex
	current uint8
}

// New provide a script object
func New(model *data.Model) *Script {
	S := Script{data: model}

	S.list = make(map[uint8]scene.Scene)
	S.m = new(sync.Mutex)
	return &S
}

// AddLoader define the scene loader on the reserved index (0)
// if loader already exist, the old older is destroy and replace
// by the new
func (s *Script) AddLoader(l scene.Scene) error {
	if s.running == true {
		return errors.New(errorChangeScriptOnRunning)
	}
	s.list[loader] = l
	return nil
}

// AddScene define the new scene on the index specified
// if index is 0 (reserved by loader), or index already used,
// a error is occured
// if first is true, the scene will be first loaded
func (s *Script) AddScene(newScene scene.Scene, index uint8, first bool) error {
	if s.running == true {
		return errors.New(errorChangeScriptOnRunning)
	}

	if index == 0 {
		return errors.New(errorIndexLoader)
	}
	if _, ok := s.list[index]; ok {
		return errors.New(errorIndexUsed)
	}

	s.list[index] = newScene
	if first {
		s.current = index
	}
	return nil
}

// Switch change the scene by the loader step
func (s *Script) Switch(index uint8, closeOld bool) error {
	var err error

	// check conditions to do the switch scene
	if _, ok := s.list[index]; !ok {
		return errors.New(errorIndexUndefined)
	}
	if !s.running {
		return errors.New(errorScriptNoRunning)
	}

	// save previous scene
	previous := s.current

	// switch to loader
	s.m.Lock()
	s.current = loader
	if err = s.list[s.current].Run(); err != nil {
		return err
	}
	s.m.Unlock()

	// close or stop the previous scene
	if closeOld {
		if err = s.list[previous].Close(); err != nil {
			return err
		}
	} else {
		s.list[previous].Stop()
	}

	// build and init new scene
	if !s.list[index].IsInit() {
		s.buildNewScene(index)
	}

	// switch to the new scene
	s.m.Lock()
	s.current = index
	if err = s.list[index].Run(); err != nil {
		return err
	}
	s.m.Unlock()

	// stop loader scene
	s.list[loader].Stop()
	return nil
}
