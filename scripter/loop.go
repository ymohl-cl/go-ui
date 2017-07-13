package scripter

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/ymohl-cl/game-builder/drivers"
)

// Run start the loop game on the sdl Main function
func (s *Script) Run(D drivers.VSDL) {
	// SDL.Main allow use sdl.Do() to queue sdl instructions.
	sdl.Main(func() {
		if err := s.start(D); err != nil {
			panic(err)
		}
	})
}

// Start is a loop game
func (s *Script) start(d drivers.VSDL) error {
	var err error

	w, h := d.GetScreenSize()
	if err = s.init(d.GetRenderer(), w, h); err != nil {
		return err
	}
	for s.isRunning() {
		sdl.Do(func() {
			for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
				switch event.(type) {
				case *sdl.QuitEvent:
					s.running = false
				default:
					s.events(event)
				}
			}
		})
		s.drawScene(d.GetRenderer())
		sdl.Do(func() {
			sdl.Delay(1000 / frame)
		})
	}
	s.close()
	return nil
}
