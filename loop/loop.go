package loop

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/ymohl-cl/game-builder/drivers"
	"github.com/ymohl-cl/game-builder/scenes"
)

const (
	frame = 60
)

type LoopGame struct {
	close  bool
	script *scenes.Scenes
}

func newLoopGame(r *sdl.Renderer) (LoopGame, error) {
	var LG LoopGame
	var err error

	if LG.script, err = scenes.Build(r); err != nil {
		return LG, err
	}

	return LG, nil
}

func Start(v drivers.VSdl) error {
	LGame, err := newLoopGame(v.GetRenderer())
	if err != nil {
		return err
	}

	if err = LGame.script.Run(); err != nil {
		return err
	}
	for LGame.close == false {
		sdl.Do(func() {
			for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
				switch event.(type) {
				case *sdl.QuitEvent:
					LGame.script.Close()
					LGame.close = true
				default:
					LGame.script.Events(event)
				}
			}
		})
		LGame.script.Draw()
		sdl.Delay(1000 / 60)
	}
	return nil
}
