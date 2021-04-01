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
			Title:  "test widget audio",
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
	var backgroundMusic *widget.Audio
	var clickAudio *widget.Audio
	var button *widget.Rect

	if backgroundMusic, err = widget.NewAudio("../resource/Muriel-Bobby-Richards.wav", true); err != nil {
		panic(err)
	}
	demo.AddWidget(backgroundMusic, goui.Layer(0))

	if clickAudio, err = widget.NewAudio("../resource/Slapping-Three-Faces.wav", false); err != nil {
		panic(err)
	}
	demo.AddWidget(clickAudio, goui.Layer(0))

	clickAudio.SetStateToPlay(widget.StateAction)
	clickAudio.SetPosition(10, 10)
	clickAudio.SetSize(200, 200)

	button = widget.NewRect(10, 10)
	demo.AddWidget(button, goui.Layer(1))

	button.SetColor(widget.ColorRed)
	button.SetPosition(10, 10)
	button.SetHoverColor(widget.ColorGreen)
	button.SetActionColor(widget.ColorBlue)
	button.SetSize(200, 200)

	// End configuration scene

	// run application
	if err = ui.Run(demo); err != nil {
		panic(err)
	}
	return
}
