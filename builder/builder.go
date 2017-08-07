package builder

import (
	"os"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/ymohl-cl/game-builder/drivers"
	"github.com/ymohl-cl/game-builder/loop"
)

// Run start the loop game
func Run(D drivers.VSDL) {
	// SDL.Main allow use sdl.Do() to queue sdl instructions.
	sdl.Main(func() {
		if err := loop.Start(D); err != nil {
			panic(err)
		}
		os.Exit(0)
	})
}
