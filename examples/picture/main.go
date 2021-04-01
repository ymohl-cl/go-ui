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
			Title:  "test widget picture",
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
	var p *widget.Picture

	if p, err = widget.NewPicture("../resource/puzzle.jpg"); err != nil {
		panic(err)
	}
	demo.AddWidget(p, goui.Layer(0))

	p.SetWidthRatioHeight(400)
	p.SetHoverColor(widget.Color{Green: 255, Alpha: 25})
	p.SetActionColor(widget.Color{Blue: 255, Alpha: 25})
	p.SetAction(widget.NewLinkAction("https://github.com/ymohl-cl/go-ui"))
	// End configuration scene

	// run application
	if err = ui.Run(demo); err != nil {
		panic(err)
	}
	return
}
