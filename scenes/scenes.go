package scenes

import (
	"github.com/42MrPiou42/game-builder/database"
	"github.com/42MrPiou42/game-builder/objects"
	"github.com/42MrPiou42/game-builder/scenes/sgame"
	"github.com/42MrPiou42/game-builder/scenes/sinfos"
	"github.com/42MrPiou42/game-builder/scenes/sstat"
	"github.com/42MrPiou42/game-builder/scenes/suser"
)

// Scenes manage the specific game.
type Scenes struct {
	Data *database.Data
	list map[uint8]Scene
}

// Scene is a interface and define the design model to your scenes.
type Scene interface {
	// Init the scene. Create static objects. Data is provide if you need.
	Init(*database.Data) error
	// Update dynamical objects
	Update(*database.Data) error
	// GetObjects return concatenation of static and dynamical objects.
	GetObjs() []*objects.ObjectType

	/*
	** follow functions are needest to interact by Events
	 */
	// Add txt string typed by player to the input field
	AddLetterToInput(string)
	// GetStaticObjs provide the static objects
	GetStaticObjs() []*objects.ObjectType
	// GetDynamicObjs provide the dynamic objects
	GetDynamicObjs() []*objects.ObjectType
	// SetNotice provide a information to the user when a bad usage is done.
	SetNotice(string)
}

// Build create a new scene manager. Define here all scenes which you want use.
func Build() (*Scenes, error) {
	var err error
	s := new(Scenes)

	s.Data, err = database.Get()
	if err != nil {
		return nil, err
	}

	s.list = make(map[uint8]Scene)
	s.list[sinfos.SceneUser] = new(suser.SUser)
	s.list[sinfos.SceneStat] = new(sstat.SStat)
	s.list[sinfos.SceneGame] = new(sgame.SGame)

	if err = s.list[sinfos.SceneUser].Init(s.Data); err != nil {
		return nil, err
	}
	if err = s.list[sinfos.SceneStat].Init(s.Data); err != nil {
		return nil, err
	}
	if err = s.list[sinfos.SceneGame].Init(s.Data); err != nil {
		return nil, err
	}

	sinfos.Current = sinfos.SceneUser
	s.list[sinfos.Current].Update(s.Data)
	return s, nil
}

func (S *Scenes) GetObjects() []*objects.ObjectType {
	return S.list[sinfos.Current].GetObjs()
}
