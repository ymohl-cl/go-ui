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
			Title:  "test widget ellipse",
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
	e := widget.NewEllipse(200, 400)

	demo.AddWidget(e, goui.Layer(0))
	e.SetColor(widget.ColorRed)
	e.SetPosition(400, 300)
	e.SetHoverColor(widget.ColorGreen)
	e.SetActionColor(widget.ColorBlue)
	e.SetAction(widget.NewLinkAction("https://github.com/ymohl-cl/go-ui"))
	// End configuration scene

	// run application
	if err = ui.Run(demo); err != nil {
		panic(err)
	}
	return
}
