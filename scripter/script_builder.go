package scripter

import (
	"errors"

	"github.com/ymohl-cl/game-builder/scene"
)

// AddLoader define the scene loader on the reserved index (0)
// if loader already exist, the old older is destroy and replace
// by the new
func (s *Script) AddLoader(l scene.Scene) error {
	if s.running == true {
		return errors.New(errorChangeScriptOnRunning)
	}
	s.list[layerLoading] = l
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
	s.current = layerLoading
	if err = s.list[s.current].Run(); err != nil {
		return err
	}
	s.m.Unlock()

	// stop and close if closeOld is true
	s.list[previous].Stop()
	if closeOld {
		if err = s.list[previous].Close(); err != nil {
			return err
		}
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
	s.list[layerLoading].Stop()
	return nil
}

// StopRun close the script
func (s *Script) StopRun() {
	s.running = false
}
