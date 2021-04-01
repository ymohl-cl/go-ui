package main

import (
	goui "github.com/ymohl-cl/go-ui"
	"github.com/ymohl-cl/go-ui/widget"
)

func main() {
	var err error
	var ui goui.GoUI

	// init drivers sdl from go-ui
	if ui, err = goui.New(goui.ConfigUI{
		Window: goui.Window{
			Title:  "test widget text",
			Width:  int32(800),
			Height: int32(600),
		},
	}); err != nil {
		panic(err)
	}
	defer func() {
		if err := ui.Close(); err != nil {
			panic(err)
		}
	}()

	var demo goui.Scene
	if demo, err = goui.NewScene(); err != nil {
		panic(err)
	}
	defer demo.Close()

	// Configuration scene
	var f widget.Font
	var i *widget.Input
	var r *widget.Rect

	if f, err = widget.NewFont("../resource/NewYork.otf", 22); err != nil {
		panic(err)
	}
	if i, err = widget.NewInput(f); err != nil {
		panic(err)
	}
	demo.AddWidget(i, goui.Layer(1))

	r = widget.NewRect(750, 50)
	demo.AddWidget(r, goui.Layer(0))
	r.SetPosition(20, 20)
	r.SetColor(widget.ColorWhite)

	i.SetPosition(20, 20)
	i.SetSize(750, 50)
	i.SetColor(widget.ColorRed)
	i.SetHoverColor(widget.ColorGreen)
	i.SetActionColor(widget.ColorBlue)
	// End configuration scene

	// run application
	if err = ui.Run(demo); err != nil {
		panic(err)
	}
	return
}
