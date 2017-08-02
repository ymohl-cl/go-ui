package loop

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/ymohl-cl/game-builder/drivers"
	"github.com/ymohl-cl/game-builder/script"
)

const (
	frame = 60
)

type LoopGame struct {
	close  bool
	script *script.Script
}

func newLoopGame(r *sdl.Renderer) (LoopGame, error) {
	var LG LoopGame
	var err error

	if LG.script, err = script.Build(r); err != nil {
		return LG, err
	}

	return LG, nil
}

func Start(v drivers.VSdl) error {
	LGame, err := newLoopGame(v.GetRenderer())
	if err != nil {
		return err
	}

	if err = LGame.script.Run(v.GetRenderer()); err != nil {
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
		LGame.script.Draw(v.GetRenderer())
		sdl.Do(func() {
			sdl.Delay(1000 / 60)
		})
	}
	return nil
}
