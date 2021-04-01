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
			Title:  "test widget rect",
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
	r := widget.NewRect(100, 200)

	demo.AddWidget(r, goui.Layer(0))
	r.SetColor(widget.ColorRed)
	r.SetHoverColor(widget.ColorGreen)
	r.SetActionColor(widget.ColorBlue)
	r.SetAction(widget.NewLinkAction("https://github.com/ymohl-cl/go-ui"))
	// End configuration scene

	// run application
	if err = ui.Run(demo); err != nil {
		panic(err)
	}
	return
}
