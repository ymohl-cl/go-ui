package scripter

import (
	"errors"
	"sync"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/ymohl-cl/game-builder/objects"
	"github.com/ymohl-cl/gomoku/conf"
	"github.com/ymohl-cl/gomoku/scenes/loader"
)

// Init : initialize the script begin and run it
func (s *Script) Init(r *sdl.Renderer) error {
	var err error

	// init loader
	if _, ok := s.list[loader]; !ok {
		s.list[loader] = loader.New(r)
	}
	if err = s.list[loader].Build(); err != nil {
		return err
	}
	if err = s.list[loader].Init(); err != nil {
		return err
	}
	if err = s.list[loader].Run(); err != nil {
		return err
	}

	// save first scene index
	first := s.current

	// start on the loader scene
	s.running = true
	s.current = loader

	// init first scene
	go func() {
		if err = s.buildNewScene(first); err != nil {
			panic(err)
		}
		if err = s.list[first].Run(); err != nil {
			panic(err)
		}
		// continue with the first scene
		s.m.Lock()
		s.current = first
		s.m.Unlock()

		// stop the loader
		s.list[loader].Stop()
	}()

	return nil
}

// IsRunning return status script
func (s Script) IsRunning() bool {
	return s.running
}

// Draw current scene loaded
func (s Script) DrawScene(r *sdl.Renderer) {
	var err error
	var wg sync.WaitGroup
	var layers map[uint8][]objects.Object

	// get objects to draw them
	s.m.Lock()
	if !s.list[s.current].IsInit() {
		s.m.Unlock()
		return
	}
	layers := S.list[s.current].GetLayers()
	s.m.Unlock()

	// clear the screen
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

func (S Script) clearScreen(r *sdl.Renderer, wg *sync.WaitGroup) {
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

// activeDraw : call Present() to apply the new screen
func (S Script) activeDraw(r *sdl.Renderer) {
	sdl.Do(func() {
		r.Present()
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
func (S *Script) Close() error {
	var err error

	S.running = false
	if _, ok := S.list[conf.Current]; ok == false {
		return errors.New("Scene tried to execute don't exist")
	}

	err = S.list[conf.Current].Close()
	return err
}
