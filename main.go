package main

import (
	"os"

	"github.com/42MrPiou42/game-builder/game"
	"github.com/42MrPiou42/game-builder/uigame"
	"github.com/veandco/go-sdl2/sdl"
)

// defer Destroy is call begin check err because Destroy() is safe
func main() {
	DriversUI, err := uigame.Init()
	defer DriversUI.Destroy()
	if err != nil {
		panic(err)
	}

	// SDL.Main allow use sdl.Do()
	sdl.Main(func() {
		if err := game.Start(DriversUI); err != nil {
			panic(err)
		}
		os.Exit(0)
	})
}
