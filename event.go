package gamebuilder

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/ymohl-cl/go-ui/objects"
)

// event : catch mouse and keyboard events
func event(e sdl.Event, s Scene) error {
	var err error

	switch e.(type) {
	case *sdl.MouseMotionEvent:
		err = mouseMotionEvent(e.(*sdl.MouseMotionEvent), s)
	case *sdl.MouseButtonEvent:
		err = mouseButtonEvent(e.(*sdl.MouseButtonEvent), s)
	case *sdl.KeyboardEvent:
		err = keyboardEvent(e.(*sdl.KeyboardEvent), s)
	}
	if err != nil {
		return err
	}
	return nil
}

// mouseMotionEvent define a new mouse's position
func mouseMotionEvent(mm *sdl.MouseMotionEvent, s Scene) error {
	layers, m := s.GetLayers()
	m.Lock()
	defer m.Unlock()

	size := len(layers)
	for i := size - 1; i >= 0; i-- {
		layer := layers[uint8(i)]
		for _, object := range layer {
			if object.IsOver(mm.X, mm.Y) {
				if object.GetStatus() != objects.SClick {
					go object.SetStatus(objects.SOver)
				}
			} else {
				go object.SetStatus(objects.SBasic)
			}
		}
	}
	return nil
}

// mouseButtonEvent define a new mouse's action (click)
func mouseButtonEvent(b *sdl.MouseButtonEvent, s Scene) error {
	if b.Button != sdl.BUTTON_LEFT {
		return nil
	}
	layers, m := s.GetLayers()
	m.Lock()
	defer m.Unlock()

	size := len(layers)
	for i := size - 1; i >= 0; i-- {
		layer := layers[uint8(i)]
		for _, object := range layer {
			if b.State == sdl.PRESSED {
				if object.GetStatus() == objects.SOver {
					go object.SetStatus(objects.SClick)
					break
				}
			} else if b.State == sdl.RELEASED {
				if object.GetStatus() == objects.SClick {
					go object.SetStatus(objects.SOver)
					go object.Click()
					break
				}
			}
		}
	}

	return nil
}

// keyboardEvent : _
func keyboardEvent(k *sdl.KeyboardEvent, s Scene) error {
	go s.KeyboardEvent(k)
	return nil
}
