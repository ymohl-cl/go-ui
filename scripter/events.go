package scripter

import (
	"errors"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/ymohl-cl/game-builder/objects"
)

// events : catch mouse and keyboard events
func (s *Script) events(E sdl.Event) {
	var err error

	switch E.(type) {
	case *sdl.MouseMotionEvent:
		err = s.mouseMotionEvent(E.(*sdl.MouseMotionEvent))
	case *sdl.MouseButtonEvent:
		err = s.mouseButtonEvent(E.(*sdl.MouseButtonEvent))
	case *sdl.KeyboardEvent:
		err = s.keyboardEvent(E.(*sdl.KeyboardEvent))
	}
	if err != nil {
		panic(err)
	}
}

// mouseMotionEvent define a new mouse's position
func (s *Script) mouseMotionEvent(mouse *sdl.MouseMotionEvent) error {
	if _, ok := s.list[s.current]; !ok {
		return errors.New(errorIndexUndefined)
	}
	layers, m := s.list[s.current].GetLayers()
	m.Lock()
	size := len(layers)
	for i := size - 1; i >= 0; i-- {
		layer := layers[uint8(i)]
		for _, object := range layer {
			if object.IsOver(mouse.X, mouse.Y) {
				if object.GetStatus() != objects.SClick {
					go object.SetStatus(objects.SOver)
				}
			} else {
				go object.SetStatus(objects.SBasic)
			}
		}
	}
	m.Unlock()
	return nil
}

// mouseButtonEvent define a new mouse's action (click)
func (s *Script) mouseButtonEvent(button *sdl.MouseButtonEvent) error {
	if button.Button != sdl.BUTTON_LEFT {
		return nil
	}
	if _, ok := s.list[s.current]; !ok {
		return errors.New(errorIndexUndefined)
	}
	layers, m := s.list[s.current].GetLayers()
	m.Lock()
	size := len(layers)
	for i := size - 1; i >= 0; i-- {
		layer := layers[uint8(i)]
		for _, object := range layer {
			if button.State == sdl.PRESSED {
				if object.GetStatus() == objects.SOver {
					go object.SetStatus(objects.SClick)
					break
				}
			} else if button.State == sdl.RELEASED {
				if object.GetStatus() == objects.SClick {
					go object.SetStatus(objects.SOver)
					go object.Click()
					break
				}
			}
		}
	}
	m.Unlock()
	return nil
}

// keyboardEvent : _
func (s *Script) keyboardEvent(keyDown *sdl.KeyboardEvent) error {
	if _, ok := s.list[s.current]; !ok {
		return errors.New(errorIndexUndefined)
	}
	go s.list[s.current].KeyboardEvent(keyDown)
	return nil
}
