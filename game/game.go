package game

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/ymohl-cl/game-builder/scenes"
	"github.com/ymohl-cl/game-builder/uigame"
)

const (
	frame = 60
)

type Game struct {
	close  bool
	script *scenes.Scenes
}

func newGame() (Game, error) {
	var G Game
	var err error

	if G.script, err = scenes.Build(); err != nil {
		return G, err
	}

	return G, nil
}

func Start(UI uigame.Drivers) error {
	Game, err := newGame()
	if err != nil {
		return err
	}

	for Game.close == false {
		sdl.Do(func() {
			for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
				switch event.(type) {
				case *sdl.QuitEvent:
					Game.close = true
				default:
					Game.script.Events(event)
				}
			}
		})
		UI.Draw(Game.script.GetObjects())
		sdl.Delay(1000 / 60)
	}
	// save data
	return nil
}
