package goui

import (
	"fmt"
	"sync"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/ymohl-cl/go-ui/widget"
)

// Scene is a interface and define the design model to your scenes
type Scene interface {
	AddWidget(w widget.Widget, l Layer)
	Render(r *sdl.Renderer)
	NewEvent(e sdl.Event) error
	Close()
}

type Layer uint8

type scene struct {
	layers  []Layer
	widgets map[Layer][]widget.Widget
}

// NewScene return the default scene to store your widget on 32 layers.
// On browse widget, if a layers represent an empty widget list, the browse action stop
func NewScene() (Scene, error) {
	var s scene

	for i := 0; i < 32; i++ {
		s.layers = append(s.layers, Layer(i))
	}
	s.widgets = make(map[Layer][]widget.Widget)
	return &s, nil
}

func (s *scene) AddWidget(w widget.Widget, layer Layer) {
	s.widgets[layer] = append(s.widgets[layer], w)
}

func (s *scene) Close() {
	for _, l := range s.layers {
		for _, w := range s.widgets[l] {
			w.Close()
		}
	}
}

func (s *scene) Render(r *sdl.Renderer) {
	var wg sync.WaitGroup

	for _, l := range s.layers {
		for _, w := range s.widgets[l] {
			wg.Add(1)
			go func(myWidget widget.Widget) {
				defer wg.Done()
				if err := myWidget.Render(r); err != nil {
					fmt.Printf("widget render failed: %s\n", err.Error())
				}
			}(w)
		}
		wg.Wait()
	}
}

func (s *scene) NewEvent(e sdl.Event) error {
	for _, l := range s.layers {
		for _, w := range s.widgets[l] {
			switch e.(type) {
			case *sdl.MouseMotionEvent:
				mouseEvent := e.(*sdl.MouseMotionEvent)
				if w.IsHover(mouseEvent.X, mouseEvent.Y) {
					w.SetState(widget.StateHover)
				} else {
					w.SetState(widget.StateBase)
				}
			case *sdl.MouseButtonEvent:
				mouseEvent := e.(*sdl.MouseButtonEvent)
				if w.IsHover(mouseEvent.X, mouseEvent.Y) {
					if mouseEvent.State == sdl.PRESSED {
						w.SetState(widget.StateAction)
					} else {
						w.SetState(widget.StateHover)
						w.Click()
					}
				} else {
					if mouseEvent.State == sdl.RELEASED {
						w.Unfocus()
					}
				}
			case *sdl.KeyboardEvent:
				w.KeyboardEvent(e.(*sdl.KeyboardEvent))
			}
		}
	}
	return nil
}
