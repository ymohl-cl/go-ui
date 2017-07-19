package scenes

import (
	"errors"
	"sync"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/ymohl-cl/game-builder/conf"
	"github.com/ymohl-cl/game-builder/database"
	"github.com/ymohl-cl/game-builder/objects"
	"github.com/ymohl-cl/game-builder/scenes/menu"
)

// Scenes manage the specific game.
type Scenes struct {
	Data *database.Data
	list map[uint8]Scene
}

func (S Scenes) Draw(r *sdl.Renderer) {
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

func (S Scenes) activeDraw(r *sdl.Renderer) {
	sdl.Do(func() {
		r.Present()
	})
}

func (S Scenes) clearDraw(r *sdl.Renderer, wg *sync.WaitGroup) {
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

// Scene is a interface and define the design model to your scenes.
type Scene interface {
	// Init the scene. Create static objects. Data is provide if you need.
	Init(*database.Data, *sdl.Renderer) error

	// Run start the scene
	Run() error

	// Close the scene
	Close() error

	// GetLayers get objects list by layers order
	GetLayers() map[uint8][]objects.Object

	// Add txt string typed by player to the input field
	//	AddLetterToInput(string)
	// GetStaticObjs provide the static objects
	//	GetStaticObjs() []*objects.ObjectType
	// GetDynamicObjs provide the dynamic objects
	//	GetDynamicObjs() []*objects.ObjectType
	// SetNotice provide a information to the user when a bad usage is done.
	//SetNotice(string)
}

// Build create a new scene manager. Define here all scenes which you want use.
func Build(r *sdl.Renderer) (*Scenes, error) {
	var err error
	s := new(Scenes)

	s.Data, err = database.Get()
	if err != nil {
		return nil, err
	}

	s.list = make(map[uint8]Scene)
	s.list[conf.SMenu] = new(menu.Menu)
	//s.list[sinfos.SceneStat] = new(sstat.SStat)
	//s.list[sinfos.SceneGame] = new(sgame.SGame)

	if err = s.list[conf.SMenu].Init(s.Data, r); err != nil {
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

func (S Scenes) Run(r *sdl.Renderer) error {
	if _, ok := S.list[conf.Current]; ok == false {
		return errors.New("Scene tried to execute don't exist")
	}
	S.list[conf.Current].Run()
	S.Draw(r)
	return nil
}

func (S Scenes) Close() error {
	var err error

	if _, ok := S.list[conf.Current]; ok == false {
		return errors.New("Scene tried to execute don't exist")
	}

	err = S.list[conf.Current].Close()
	return err
}

/*func (S *Scenes) GetObjects() []*objects.ObjectType {
	return S.list[sinfos.Current].GetObjs()
}*/
