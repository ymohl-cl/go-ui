package script

import (
	"errors"
	"sync"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/ymohl-cl/game-builder/conf"
	"github.com/ymohl-cl/game-builder/database"
	"github.com/ymohl-cl/game-builder/scenes"
	"github.com/ymohl-cl/game-builder/scenes/menu"
)

// Scenes manage the specific game.
type Script struct {
	Data *database.Data
	list map[uint8]scenes.Scene
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
	if s.list[conf.SMenu], err = menu.New(s.Data, r); err != nil {
		return nil, err
	}
	//s.list[sinfos.SceneStat] = new(sstat.SStat)
	//s.list[sinfos.SceneGame] = new(sgame.SGame)

	if err = s.list[conf.SMenu].Init(); err != nil {
		return nil, err
	}
	/*if err = s.list[sinfos.SceneStat].Init(s.Data); err != nil {
		return nil, err
	}
	if err = s.list[sinfos.SceneGame].Init(s.Data); err != nil {
		return nil, err
	}*/

	conf.Current = conf.SMenu
	//s.list[conf.Current].Update(s.Data)
	return s, nil
}

func (S Script) Draw(r *sdl.Renderer) {
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

func (S Script) Run(r *sdl.Renderer) error {
	if _, ok := S.list[conf.Current]; ok == false {
		return errors.New("Scene tried to execute don't exist")
	}
	S.list[conf.Current].Run()
	S.Draw(r)
	return nil
}

func (S Script) Close() error {
	var err error

	if _, ok := S.list[conf.Current]; ok == false {
		return errors.New("Scene tried to execute don't exist")
	}

	err = S.list[conf.Current].Close()
	return err
}
