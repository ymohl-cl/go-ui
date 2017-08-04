package script

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/ymohl-cl/game-builder/conf"
	"github.com/ymohl-cl/game-builder/database"
	"github.com/ymohl-cl/game-builder/scenes"
	"github.com/ymohl-cl/game-builder/scenes/gomoku"
	"github.com/ymohl-cl/game-builder/scenes/loader"
	"github.com/ymohl-cl/game-builder/scenes/menu"
)

// Script manage the specific game.
type Script struct {
	Data   *database.Data
	list   map[uint8]scenes.Scene
	loader scenes.Scene
}

// Build create a new scene manager. Define here all scenes which you want use.
func Build(r *sdl.Renderer) (*Script, error) {
	var err error
	s := new(Script)

	s.Data, err = database.Get()
	if err != nil {
		return nil, err
	}

	s.list = make(map[uint8]scenes.Scene)
	if s.list[conf.Sload], err = loader.New(s.Data, r); err != nil {
		return nil, err
	}
	if err = s.list[conf.Sload].Init(); err != nil {
		return nil, err
	}
	conf.Current = conf.Sload

	go func() {
		if s.list[conf.SMenu], err = menu.New(s.Data, r); err != nil {
			panic(err)
		}
		if err = s.list[conf.SMenu].Init(); err != nil {
			panic(err)
		}
		if s.list[conf.SGame], err = gomoku.New(s.Data, r); err != nil {
			panic(err)
		}
		time.Sleep(5 * time.Second)
		conf.Current = conf.SMenu
		if err = s.list[conf.SMenu].Run(); err != nil {
			panic(err)
		}
		if err = s.list[conf.Current].Close(); err != nil {
			panic(err)
		}
	}()

	return s, nil
}

// Draw : current scene loaded
func (S Script) Draw(r *sdl.Renderer) {
	var err error

	if S.list[conf.Current].IsInit() == false {
		fmt.Println("IS FALSE")
		saveScene := conf.Current
		conf.Current = conf.Sload
		if err = S.list[conf.Current].Init(); err != nil {
			panic(err)
		}
		if err = S.list[conf.Current].Run(); err != nil {
			panic(err)
		}

		go func() {
			if err = S.list[saveScene].Init(); err != nil {
				panic(err)
			}
			time.Sleep(5 * time.Second)
			if err = S.list[conf.Current].Close(); err != nil {
				panic(err)
			}
			if err = S.list[saveScene].Run(); err != nil {
				panic(err)
			}
			conf.Current = saveScene
		}()
	}
	layers := S.list[conf.Current].GetLayers()

	var wg sync.WaitGroup

	wg.Add(1)
	go S.clearDraw(r, &wg)
	wg.Wait()
	for i := 0; layers[uint8(i)] != nil; i++ {
		layer := layers[uint8(i)]
		for _, object := range layer {
			wg.Add(1)
			go object.Draw(&wg, r)
		}
		wg.Wait()
	}
	S.activeDraw(r)
}

// activeDraw : call Present() to apply the new screen
func (S Script) activeDraw(r *sdl.Renderer) {
	sdl.Do(func() {
		r.Present()
	})
}

func (S Script) clearDraw(r *sdl.Renderer, wg *sync.WaitGroup) {
	defer wg.Done()
	var err error

	sdl.Do(func() {
		if err = r.SetDrawColor(conf.ClearRGBO, conf.ClearRGBO, conf.ClearRGBO, conf.ClearRGBO); err != nil {
			panic(err)
		}
		if err = r.Clear(); err != nil {
			panic(err)
		}
	})
}

// Run : current scene loaded
func (S Script) Run(r *sdl.Renderer) error {
	if _, ok := S.list[conf.Current]; ok == false {
		return errors.New("Scene tried to execute don't exist")
	}
	S.list[conf.Current].Run()
	S.Draw(r)
	return nil
}

// Close : current scene loaded
func (S Script) Close() error {
	var err error

	if _, ok := S.list[conf.Current]; ok == false {
		return errors.New("Scene tried to execute don't exist")
	}

	err = S.list[conf.Current].Close()
	return err
}
