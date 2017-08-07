package main

import (
	"os"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/ymohl-cl/game-builder/drivers"
	"github.com/ymohl-cl/game-builder/loop"
)

// Drivers.Destroy is safe.
func main() {
	Drivers, err := drivers.Init()
	defer Drivers.Destroy()
	if err != nil {
		panic(err)
	}

	// SDL.Main allow use sdl.Do() to queue sdl instructions.
	sdl.Main(func() {
		if err := loop.Start(Drivers); err != nil {
			panic(err)
		}
		os.Exit(0)
	})
}
