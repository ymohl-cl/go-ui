package scripter

import (
	"sync"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/ymohl-cl/game-builder/objects"
	"github.com/ymohl-cl/game-builder/scene/loader"
)

// buildNewScene : build the scene specified by the index parameter
func (s *Script) buildNewScene(index uint8) error {
	var err error

	// set switcher function
	s.list[index].SetSwitcher(s.Switch)
	// build scene
	if err = s.list[index].Build(); err != nil {
		return err
	}
	// init scene
	if err = s.list[index].Init(); err != nil {
		return err
	}

	return nil
}

// Init : initialize the script begin and run it
func (s *Script) init(r *sdl.Renderer, widthScreen, heightScreen int32) error {
	var err error

	// init loader
	if _, ok := s.list[layerLoading]; !ok {
		if s.list[layerLoading], err = loader.New(r, widthScreen, heightScreen); err != nil {
			return err
		}
	}
	if err = s.buildNewScene(layerLoading); err != nil {
		return err
	}
	if err = s.list[layerLoading].Run(); err != nil {
		return err
	}

	// save first scene index
	first := s.current

	// start on the loader scene
	s.running = true
	s.current = layerLoading

	// if no scene loaded, leave function
	if first == layerLoading {
		return nil
	}
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
		s.list[layerLoading].Stop()
	}()

	return nil
}

// IsRunning return status script
func (s Script) isRunning() bool {
	return s.running
}

// DrawScene : current scene loaded
func (s Script) drawScene(r *sdl.Renderer) {
	var wg sync.WaitGroup
	var layers map[uint8][]objects.Object
	var m *sync.Mutex

	// Lock current acces list
	s.m.Lock()
	if !s.list[s.current].IsInit() {
		s.m.Unlock()
		return
	}
	// call update
	go s.list[s.current].Update()
	// get objects to draw them
	layers, m = s.list[s.current].GetLayers()
	// Lock to protect layers data
	m.Lock()
	// Unlock current access list
	s.m.Unlock()

	// clear the screen
	wg.Add(1)
	go s.clearScreen(r, &wg)
	wg.Wait()

	s.drawLayers(layers, r)
	// Unlock layers data
	m.Unlock()

	s.appliesScreen(r)
}

// clearScreen reset the screen
func (s Script) clearScreen(r *sdl.Renderer, wg *sync.WaitGroup) {
	defer wg.Done()
	var err error

	sdl.Do(func() {
		if err = r.SetDrawColor(0, 0, 0, 0); err != nil {
			panic(err)
		}
		if err = r.Clear(); err != nil {
			panic(err)
		}
	})
}

// drawLayers, draw objects by order layer
func (s Script) drawLayers(l map[uint8][]objects.Object, r *sdl.Renderer) {
	var wg sync.WaitGroup

	//	fmt.Println("DrawLayer")
	for i := 0; l[uint8(i)] != nil; i++ {
		//		fmt.Println("index: ", i)
		layer := l[uint8(i)]
		for _, object := range layer {
			//			fmt.Println("object number: ", it)
			if object.IsInit() {
				wg.Add(1)
				go object.Draw(&wg, r)
			}
		}
		wg.Wait()
	}
}

// appliesScreen : call Present() to applies the new screen
func (s Script) appliesScreen(r *sdl.Renderer) {
	sdl.Do(func() {
		r.Present()
	})
}

// close all ressources
func (s *Script) close() {
	var err error

	if scene, ok := s.list[s.current]; ok {
		scene.Stop()
	}

	for _, scene := range s.list {
		if scene.IsInit() {
			if err = scene.Close(); err != nil {
				panic(err)
			}
		}
	}
}
