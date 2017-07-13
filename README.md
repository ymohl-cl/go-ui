# game-builder
Game-builder is a game builder lite using [SDL2 veandco](https://github.com/veandco/go-sdl2).
It's not convention to compose a game, just proof of concept today.

The wrapper is implemented with goroutines.

I try to define a good architecture and simple of use.
Inspired by Unity 3D (without interface home machine).
I try to provide a good tool to creates models scenes / layout and objects.

I would be pleased to talk about that with you on mail:
`mohl.clauzade@gmail.com`

# Requirements
See official github to SDL2 binding for Go by [veandco](https://github.com/veandco/go-sdl2).

# Installation
To install Golang see [getting started](https://golang.org/doc/install)

To get SDL2 wrapper see [veandco](https://github.com/veandco/go-sdl2)

`go get -v github.com/veandco/go-sdl2/{sdl,mix,img,ttf}`

To get Game-Builder:

`go get -v github.com/ymohl-cl/game-builder`

#### OSX
Install SDL2:

`brew install sdl2 sdl2_gfx sdl2_image sdl2_mixer sdl2_net sdl2_ttf`

# Example
You can view implementation example on this project [gomoku game](https://github.com/ymohl-cl/gomoku)

Please, read godoc to know the specifications

```
package main

import (
	"github.com/ymohl-cl/game-builder/drivers"
	"github.com/ymohl-cl/game-builder/scripter"
	"github.com/ymohl-cl/gomoku/database"
	"github.com/ymohl-cl/gomoku/scenes/gomoku"
	"github.com/ymohl-cl/gomoku/scenes/loader"
	"github.com/ymohl-cl/gomoku/scenes/menu"
)

const (
	windowWidth = 800
	windowHeight = 600

	indexMenu = 1 << iota
	indexGomoku
)

func main() {
	var err error
	var d drivers.VSDL
	var data *database.Data

	// init drivers sdl from game-builder
	if d, err = drivers.Init(windowWidth, windowHeight, "Title of my windows"); err != nil {
		panic(err)
	}
	defer d.Destroy()

	// get your data app from your package, for the example, we used database package
	if data, err = database.Get(); err != nil {
		panic(err)
	}

	// get new scripter application from game-builder
	s := scripter.New()

	// get loader scene from my app
	var loaderScene *loader.Load
	if loaderScene, err = loader.New(nil, d.GetRenderer()); err != nil {
		panic(err)
	}
	// add scene on the scripter (game-builder)
	if err = s.AddLoader(loaderScene); err != nil {
		panic(err)
	}

	// get menu scene from my app
	var menuScene *menu.Menu
	if menuScene, err = menu.New(data, d.GetRenderer()); err != nil {
		panic(err)
	}
	// add scene on the scripter (game-builder)
	if err = s.AddScene(menuScene, indexMenu, true); err != nil {
		panic(err)
	}

	// get gomoku scene from my app
	var gameScene *gomoku.Gomoku
	if gameScene, err = gomoku.New(data, d.GetRenderer()); err != nil {
		panic(err)
	}
	// add scene on the scripter (game-builder)
	if err = s.AddScene(gameScene, indexGomoku, false); err != nil {
		panic(err)
	}

	// run application from game-builder
	s.Run(d)
}
```

# FAQ
#### Why shaders aren't implemented ?
Game-builder is a proof of concept for the moment. This lib provide that which are needest to make a simple project.
If you need shaders, please contact us.

#### How do I contribute ?
Contact me by mail: `mohl.clauzade@gmail.com`

# Acknowledgment
Thanks at [veandco](https://github.com/veandco/go-sdl2) for their work.

# License
game-builder is BSD 3-clause licensed.

# Version
V-0.1.1: implement library game-builder
