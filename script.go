package gamebuilder

import (
	"errors"
	"sync"
)

// Script describe the application/game scenes and the current use
type Script interface {
	Close() error
	SetLoader(s Scene) error
	AddScene(index string, s Scene) error
	LoadScene(index string) error
	CloseScene(index string) error
	Scene(index string) (Scene, error)
	SceneIndex() string
}

type script struct {
	sync.Mutex
	loader  Scene
	scenes  map[string]Scene
	scene   string
	running bool
}

// NewScript instance
func NewScript() Script {
	return &script{scenes: make(map[string]Scene)}
}

// Close free all scenes resources
func (s *script) Close() error {
	var err error

	s.scene = ""
	if s.loader != nil && s.loader.IsInit() {
		if err = s.loader.Close(); err != nil {
			return err
		}
	}
	for _, sn := range s.scenes {
		if err = sn.Close(); err != nil {
			return err
		}
	}
	return nil
}

// SetLoader build a loader scene to use it in transition
func (s *script) SetLoader(sn Scene) error {
	var err error

	if s.running {
		return errors.New(errorScriptModifier)
	}
	if s.loader != nil {
		if err = s.loader.Close(); err != nil {
			return err
		}
	}

	s.loader = sn
	return nil
}

// AddScene ajoute the scene in the gamebuilder script. The scene should be called by the index
func (s *script) AddScene(index string, sn Scene) error {
	var err error

	if s.running {
		return errors.New(errorScriptModifier)
	}
	if old, ok := s.scenes[index]; ok {
		if err = old.Close(); err != nil {
			return err
		}
	}
	s.scenes[index] = sn
	sn.SetSwitcher(s.LoadScene)
	sn.SetCloser(s.CloseScene)
	return nil
}

// LoadScene by index reference
func (s *script) LoadScene(index string) error {
	var err error
	var sn Scene

	if !s.running {
		s.running = true
	}
	ok := false
	if sn, ok = s.scenes[index]; !ok {
		return errors.New(errorSceneNotFound)
	}

	s.Lock()
	defer s.Unlock()
	// start loader if configured
	if s.loader != nil {
		if !s.loader.IsInit() {
			if err = s.loader.Build(); err != nil {
				return err
			}
			if err = s.loader.Init(); err != nil {
				return err
			}
		}
		if err = s.loader.Run(); err != nil {
			return err
		}
	}

	// stop the previous scene
	if s.scene != "" {
		s.scenes[s.scene].Stop()
	}

	// setup the new scene
	if !sn.IsInit() {
		if err = sn.Build(); err != nil {
			return err
		}
		if err = sn.Init(); err != nil {
			return err
		}
	}

	// start the new scene
	s.scene = index
	if err = sn.Run(); err != nil {
		return err
	}

	// stop the loader if configured
	if s.loader != nil {
		s.loader.Stop()
	}

	return nil
}

// CloseScene by index reference
func (s *script) CloseScene(index string) error {
	var err error
	var sn Scene

	ok := false
	if sn, ok = s.scenes[index]; !ok {
		return errors.New(errorSceneNotFound)
	}

	s.Lock()
	defer s.Unlock()
	if err = sn.Close(); err != nil {
		return err
	}

	return nil
}

// SceneIndex getter return the current scene loaded
func (s *script) SceneIndex() string {
	return s.scene
}

// Scene getter return current scene by index reference
func (s *script) Scene(index string) (Scene, error) {
	var sn Scene

	ok := false
	if sn, ok = s.scenes[index]; !ok {
		return nil, errors.New(errorSceneNotFound)
	}

	return sn, nil
}
