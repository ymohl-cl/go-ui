package menu

import (
	"errors"
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/ymohl-cl/game-builder/database"
	"github.com/ymohl-cl/game-builder/objects"
)

/*
** Endpoint Event from scenes
 */

func (M *Menu) KeyDownEvent(keyDown *sdl.KeyDownEvent) {
	var err error

	if err = M.input.SetNewRune(keyDown.Keysym, M.renderer); err != nil {
		go M.setNotice(err.Error())
	}
}

/*
** Endpoint action from objects click
 */

func (M *Menu) LoadGame(values ...interface{}) {
	fmt.Println("Load Game")
}

func (M *Menu) DeletePlayer(values ...interface{}) {
	fmt.Println("Delete Player")
}

func (M *Menu) DrawStat(values ...interface{}) {
	fmt.Println("Draw stat")
}

func (M *Menu) SelectPlayer(values ...interface{}) {
	var p *database.Player
	var err error
	var ok bool

	if len(values) == 1 {
		p, ok = values[0].(*database.Player)
		if !ok {
			panic(errorInterface)
		}
	} else {
		panic(errorValuesEmpty)
	}
	if err = M.data.UpdateCurrent(p); err != nil {
		go M.setNotice(err.Error())
		return
	}
	M.updateVS()
}

func (M *Menu) InputNewPlayer(values ...interface{}) {
	fmt.Println("Input New player")
}

func (M *Menu) NewPlayer(values ...interface{}) {
	var name string
	var nbPlayer int
	var err error

	nbPlayer = len(M.data.Players)
	if nbPlayer >= playerMax {
		go M.setNotice(noticeMaxPlayer)
		return
	}
	name = M.input.GetTxt()
	if len(name) == 0 {
		go M.setNotice(noticeNameEmpty)
		return
	}

	for _, p := range M.data.Players {
		if p.Name == name {
			go M.setNotice(noticeNameExist)
			return
		}
	}
	p := database.CreatePlayer(name)
	M.data.AddPlayer(p)
	M.input.Reset(M.renderer)

	if err = M.addUIPlayer(nbPlayer, p); err != nil {
		panic(err)
	}
}

func (M *Menu) Play(values ...interface{}) {
	fmt.Println("Play")
	go M.setNotice("Play")
}

func (M *Menu) ResetName(values ...interface{}) {
	fmt.Println("ResetName")
	go M.setNotice("ResetName")
}

func (M *Menu) DefaultPlayer(values ...interface{}) {
	fmt.Println("Default Player")
	go M.setNotice("Default Player")
}

/*
** Change object from Endpoint management
 */
func (M *Menu) setNotice(str string) {
	idSDL := M.notice.NewIDSDL()
	if M.notice.IsInit() == true {
		M.notice.Close()
	}
	M.notice.SetText(str)
	if err := M.notice.Init(M.renderer); err != nil {
		panic(errors.New(objects.ErrorRenderer))
	}
	time.Sleep(3 * time.Second)
	if M.notice.GetIdSDL() == idSDL {
		M.notice.Close()
	}
}
