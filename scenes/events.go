package scenes

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/ymohl-cl/game-builder/objects"
	"github.com/ymohl-cl/game-builder/scenes/sinfos"
)

func (S *Scenes) Events(E sdl.Event) {
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
	case *sdl.TextInputEvent:
		err = S.textInputEvent(E.(*sdl.TextInputEvent))
	}
	if err != nil {
		panic(err)
	}
}

func (S *Scenes) textInputEvent(input *sdl.TextInputEvent) error {
	if sdl.IsTextInputActive() == true {
		S.list[sinfos.Current].AddLetterToInput(string(input.Text[0]))
		if err := S.list[sinfos.Current].Update(S.Data); err != nil {
			return err
		}
	}
	return nil
}

func (S *Scenes) mouseMotionEvent(mouse *sdl.MouseMotionEvent) error {
	var ret bool

	ret = false
	objs := S.list[sinfos.Current].GetDynamicObjs()
	_ = defineIsOver(objs, mouse.X, mouse.Y)
	if ret == false {
		objs = S.list[sinfos.Current].GetStaticObjs()
		_ = defineIsOver(objs, mouse.X, mouse.Y)
	}

	return nil
}

func defineIsOver(slc []*objects.ObjectType, mx int32, my int32) bool {
	for _, tobj := range slc {
		if len(tobj.Childs) > 0 {
			_ = defineIsOver(tobj.GetChildsWithType(), mx, my)
			/*			if ret == true {
						return ret
					}*/
		}
		if tobj.Status == objects.StatusFix {
			fmt.Println("Status Fix")
			continue
		}
		fmt.Println("Status traitment")
		x, y, checkPos := tobj.GetPos()
		w, h, checkSize := tobj.GetSize()
		if checkPos == true && checkSize == true {
			if mx >= x && mx <= x+w && my >= y && my <= y+h {
				tobj.SetStatus(objects.StatusOver)
			} else {
				tobj.SetStatus(objects.StatusNormal)
			}
		}
	}
	return false
}

func (S *Scenes) mouseButtonEvent(button *sdl.MouseButtonEvent) error {
	var ret bool

	ret = false
	if button.Button == sdl.BUTTON_LEFT {
		objs := S.list[sinfos.Current].GetDynamicObjs()
		_ = S.defineIsPressed(objs, button.State)
		if ret == false {
			objs = S.list[sinfos.Current].GetStaticObjs()
			ret = S.defineIsPressed(objs, button.State)
		}
	}

	//S.list[sinfos.Current].Update(S.data)
	/*	if ret == true {
		S.list[sinfos.Current].Update(S.data)
	}*/
	return nil
}

func (S *Scenes) defineIsPressed(slc []*objects.ObjectType, state uint8) bool {
	for _, tobj := range slc {
		if len(tobj.Childs) > 0 {
			_ = S.defineIsPressed(tobj.GetChildsWithType(), state)
			/*			if r == true {
						return true
					}*/
		}
		if tobj.Status == objects.StatusOver && state == sdl.PRESSED {
			tobj.SetStatus(objects.StatusClicDown)
			return true
		} else if tobj.Status == objects.StatusClicDown && state == sdl.RELEASED {
			tobj.SetStatus(objects.StatusOver)
			str := tobj.Action(tobj.ActionDatas...)
			S.list[sinfos.Current].SetNotice(str)
			S.list[sinfos.Current].Update(S.Data)
			return true
		}
	}
	return false
}

func (S *Scenes) keyDownEvent(keyDown *sdl.KeyDownEvent) error {
	//	fmt.Println("Key down: ", keyDown)
	return nil
}

func (S *Scenes) keyUpEvent(keyUp *sdl.KeyUpEvent) error {
	/*	if keyUp.Keysym.Scancode == sdl.SDL_SCANCODE_RETURN {
		if sdl.IsTextInputActive() == true {
			sdl.StopTextInput()
		}
	}*/
	return nil
}
