package widget

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	StateBase StateWidget = iota
	StateHover
	StateAction
	StateOff
)

type StateWidget uint8

type Widget interface {
	SetColor(c Color)
	SetHoverColor(c Color)
	SetActionColor(c Color)
	SetAction(Action)
	Color(StateWidget) Color
	Position() Position
	SetPosition(x, y int32)
	IsHover(x, y int32) bool
	SetState(s StateWidget)
	State() StateWidget
	SetSize(w, h int32)
	Click()
	Unfocus()
	KeyboardEvent(key *sdl.KeyboardEvent)
	Size() (int32, int32)

	Render(r *sdl.Renderer) error
	Close()
}

type widget struct {
	action   Action
	colors   Colors
	position Position
	block    Block
	state    StateWidget
	focus    bool
}

func (w *widget) SetColor(c Color) {
	w.colors.base = &Color{
		Red:   c.Red,
		Green: c.Green,
		Blue:  c.Blue,
		Alpha: c.Alpha,
	}
}

func (w *widget) SetHoverColor(c Color) {
	w.colors.hover = &Color{
		Red:   c.Red,
		Green: c.Green,
		Blue:  c.Blue,
		Alpha: c.Alpha,
	}
}

func (w *widget) SetActionColor(c Color) {
	w.colors.action = &Color{
		Red:   c.Red,
		Green: c.Green,
		Blue:  c.Blue,
		Alpha: c.Alpha,
	}
}

func (w *widget) SetPosition(x, y int32) {
	w.position.X = x
	w.position.Y = y
}

func (w widget) Color(state StateWidget) Color {
	switch state {
	case StateBase:
		if w.colors.base != nil {
			return *w.colors.base
		}
	case StateHover:
		if w.colors.hover != nil {
			return *w.colors.hover
		} else if w.colors.base != nil {
			return *w.colors.base
		}
	case StateAction:
		if w.colors.action != nil {
			return *w.colors.action
		} else if w.colors.base != nil {
			return *w.colors.base
		}
	}
	return Color{}
}

func (w widget) Position() Position {
	return w.position
}

func (w widget) IsHover(x, y int32) bool {
	if x >= w.position.X && x <= w.position.X+w.block.width {
		if y >= w.position.Y && y <= w.position.Y+w.block.height {
			return true
		}
	}
	return false
}

// Unfocus remove the focus on the widget
// focus define the last interaction with the widget (ex with input)
func (w *widget) Unfocus() {
	w.focus = false
}

func (w *widget) SetState(s StateWidget) {
	w.state = s
}

// State getter
func (w widget) State() StateWidget {
	return w.state
}

func (w *widget) SetAction(a Action) {
	w.action = a
}

func (w *widget) SetSize(width, height int32) {
	w.block.width = width
	w.block.height = height
}

// Size return the width and height block
func (w widget) Size() (int32, int32) {
	return w.block.width, w.block.height
}

func (w *widget) Click() {
	w.focus = true
	if w.action == nil {
		return
	}
	if err := w.action.Run(); err != nil {
		// log error
		fmt.Printf("error to click action")
	}
}

func (w *widget) KeyboardEvent(key *sdl.KeyboardEvent) {}
