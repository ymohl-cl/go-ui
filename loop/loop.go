package loop

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/ymohl-cl/game-builder/drivers"
	"github.com/ymohl-cl/game-builder/scripter"
)

const (
	frame = 60
)

// Start : _
func Start(d drivers.VSDL, s *scripter.Script) error {
	var err error

	if err = s.Init(d.GetRenderer()); err != nil {
		return err
	}
	for s.IsRunning() {
		sdl.Do(func() {
			for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
				switch event.(type) {
				case *sdl.QuitEvent:
					s.Close()
				default:
					s.Events(event)
				}
			}
		})
		s.DrawScene(d.GetRenderer())
		sdl.Do(func() {
			sdl.Delay(1000 / frame)
		})
	}
	return nil
}
