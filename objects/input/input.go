package input

import (
	"errors"
	"sync"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/ymohl-cl/game-builder/objects"
	"github.com/ymohl-cl/game-builder/objects/block"
	"github.com/ymohl-cl/game-builder/objects/text"
)

type Input struct {
	// infos object
	status      uint8
	initialized bool

	// content object
	txt    *text.Text
	bFix   *block.Block
	bBasic *block.Block
	bOver  *block.Block
	bClick *block.Block
}

func (I *Input) Init(r *sdl.Renderer) error {
	var err error

	if r == nil {
		return errors.New(objects.ErrorRenderer)
	}

	if I.bFix == nil {
		return errors.New(objects.ErrorEmpty)
	}
	if err = I.bFix.Init(r); err != nil {
		return err
	}

	if I.bBasic == nil {
		return errors.New(objects.ErrorEmpty)
	}
	if err = I.bBasic.Init(r); err != nil {
		return err
	}

	if I.bOver == nil {
		return errors.New(objects.ErrorEmpty)
	}
	if err = I.bOver.Init(r); err != nil {
		return err
	}

	if I.bClick == nil {
		return errors.New(objects.ErrorEmpty)
	}
	if err = I.bClick.Init(r); err != nil {
		return err
	}

	var pos *objects.Position
	var size *objects.Size
	if pos, err = I.bBasic.GetPosition(); err != nil {
		return err
	}
	if size, err = I.bBasic.GetSize(); err != nil {
		return err
	}
	posTXT := new(objects.Position)
	posTXT.X = pos.X + (size.W / 2)
	posTXT.Y = pos.Y + (size.H / 2)
	if err = I.txt.SetPosition(posTXT); err != nil {
		return err
	}

	if err = I.txt.Init(r); err != nil {
		return err
	}

	I.initialized = true
	return nil
}

// IsInit return status initialize
func (I *Input) IsInit() bool {
	return I.initialized
}

func (I *Input) Close() error {
	I.initialized = false

	if err := I.bBasic.Close(); err != nil {
		return err
	}
	if err := I.bOver.Close(); err != nil {
		return err
	}
	if err := I.bClick.Close(); err != nil {
		return err
	}

	return nil
}

func (I *Input) GetStatus() uint8 {
	return I.status
}

func (I *Input) IsOver(x, y int32) bool {
	var pos *objects.Position
	var size *objects.Size
	var err error

	switch I.status {
	case objects.SFix:
		if pos, err = I.bFix.GetPosition(); err != nil {
			panic(err)
		}
		if size, err = I.bFix.GetSize(); err != nil {
			panic(err)
		}
	case objects.SBasic:
		if pos, err = I.bBasic.GetPosition(); err != nil {
			panic(err)
		}
		if size, err = I.bBasic.GetSize(); err != nil {
			panic(err)
		}
	case objects.SOver:
		if pos, err = I.bOver.GetPosition(); err != nil {
			panic(err)
		}
		if size, err = I.bOver.GetSize(); err != nil {
			panic(err)
		}
	case objects.SClick:
		if pos, err = I.bClick.GetPosition(); err != nil {
			panic(err)
		}
		if size, err = I.bClick.GetSize(); err != nil {
			panic(err)
		}
	}

	if x > pos.X && x < pos.X+size.W {
		if y > pos.Y && y < pos.Y+size.H {
			return true
		}
	}
	return false
}

func (I *Input) Click() {
	if I.status != objects.SFix {
		I.status = objects.SClick
	}
}

func (I *Input) SetStatus(s uint8) {
	if I.status != objects.SFix {
		I.status = s
	}
}

func (I *Input) Draw(wg *sync.WaitGroup, r *sdl.Renderer) {
	defer wg.Done()

	if I.initialized == false {
		return
	}
	if r == nil {
		panic(errors.New(objects.ErrorRenderer))
	}

	wg.Add(1)
	switch I.status {
	case objects.SFix:
		I.bFix.Draw(wg, r)
	case objects.SBasic:
		I.bBasic.Draw(wg, r)
	case objects.SOver:
		I.bOver.Draw(wg, r)
	case objects.SClick:
		I.bClick.Draw(wg, r)
	default:
		wg.Done()
	}
	wg.Add(1)
	go I.txt.Draw(wg, r)
}
