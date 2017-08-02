package script

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/ymohl-cl/game-builder/conf"
	"github.com/ymohl-cl/game-builder/objects"
)

func (S *Script) Events(E sdl.Event) {
	var err error

	switch E.(type) {
	case *sdl.MouseMotionEvent:
		err = S.mouseMotionEvent(E.(*sdl.MouseMotionEvent))
	case *sdl.MouseButtonEvent:
		err = S.mouseButtonEvent(E.(*sdl.MouseButtonEvent))
	case *sdl.KeyDownEvent:
		err = S.keyDownEvent(E.(*sdl.KeyDownEvent))
	case *sdl.KeyUpEvent:
		err = S.keyUpEvent(E.(*sdl.KeyUpEvent))
	}
	if err != nil {
		panic(err)
	}
}

func (S *Script) mouseMotionEvent(mouse *sdl.MouseMotionEvent) error {
	layers := S.list[conf.Current].GetLayers()

	size := len(layers)
	for i := size - 1; layers[uint8(i)] != nil; i-- {
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
	return nil
}

func (S *Script) mouseButtonEvent(button *sdl.MouseButtonEvent) error {
	if button.Button == sdl.BUTTON_LEFT {

		layers := S.list[conf.Current].GetLayers()

		size := len(layers)
		for i := size - 1; layers[uint8(i)] != nil; i-- {
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
	}
	return nil
}

func (S *Script) keyDownEvent(keyDown *sdl.KeyDownEvent) error {
	go S.list[conf.Current].KeyDownEvent(keyDown)
	//	fmt.Println("Key down: ", keyDown)
	return nil
}

func (S *Script) keyUpEvent(keyUp *sdl.KeyUpEvent) error {
	//	fmt.Println("HELLO KEY UP: ", keyUp.Keysym.Scancode)
	/*	if keyUp.Keysym.Scancode == sdl.SDL_SCANCODE_RETURN {
		if sdl.IsTextInputActive() == true {
			sdl.StopTextInput()
		}
	}*/
	return nil
}
