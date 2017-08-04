package menu

import (
	"errors"
	"fmt"
	"time"

	"github.com/ymohl-cl/game-builder/conf"
	"github.com/ymohl-cl/game-builder/database"
	"github.com/ymohl-cl/game-builder/objects"
)

/*
** Endpoint action from objects click
 */

// LoadGame : start a new party
func (M *Menu) LoadGame(values ...interface{}) {
	fmt.Println("Load Game")
}

// DeletePlayer : _
func (M *Menu) DeletePlayer(values ...interface{}) {
	var p *database.Player
	var err error
	var ok bool
	var id int

	if len(values) == 1 {
		p, ok = values[0].(*database.Player)
		if !ok {
			panic(errorInterface)
		}
	} else {
		panic(errorValuesEmpty)
	}

	if id, err = M.data.DeletePlayer(p); err != nil {
		go M.setNotice(err.Error())
		return
	}
	if err = M.removeUIPlayer(id); err != nil {
		panic(err)
	}
	M.updateVS()
}

// DrawStat : on user selected
func (M *Menu) DrawStat(values ...interface{}) {
	fmt.Println("Draw stat")
}

// SelectPlayer : to the futur game
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

// NewPlayer : on the database
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

// Play start the game
func (M *Menu) Play(values ...interface{}) {
	var err error

	conf.Current = conf.SGame
	if err = M.Close(); err != nil {
		panic(err)
	}
}

// ResetName : reset the input value
func (M *Menu) ResetName(values ...interface{}) {
	M.input.Reset(M.renderer)
}

// DefaultPlayer : init the defaults player to the game
func (M *Menu) DefaultPlayer(values ...interface{}) {
	var err error

	if err = M.data.DefaultPlayers(); err != nil {
		panic(err.Error)
	}
	M.updateVS()
}

/*
** Change object from Endpoint management
 */
func (M *Menu) setNotice(str string) {
	idSDL := M.notice.NewIDSDL()
	if M.notice.IsInit() == true {
		M.notice.Close()
	}
	M.notice.UpdateText(str, M.renderer)
	if err := M.notice.Init(M.renderer); err != nil {
		panic(errors.New(objects.ErrorRenderer))
	}
	time.Sleep(3 * time.Second)
	if M.notice.GetIDSDL() == idSDL {
		M.notice.Close()
	}
}
