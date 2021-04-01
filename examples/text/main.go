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
	var t *widget.Text

	if f, err = widget.NewFont("../resource/NewYork.otf", 22); err != nil {
		panic(err)
	}
	if t, err = widget.NewText("Ceci est un message", f); err != nil {
		panic(err)
	}
	demo.AddWidget(t, goui.Layer(0))

	t.SetColor(widget.ColorRed)
	t.SetHoverColor(widget.ColorGreen)
	t.SetActionColor(widget.ColorBlue)
	t.SetAction(widget.NewLinkAction("https://github.com/ymohl-cl/go-ui"))
	// End configuration scene

	// run application
	if err = ui.Run(demo); err != nil {
		panic(err)
	}
	return
}
