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
	var x, y int32

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

	var size *objects.Size

	x, y = I.bBasic.GetPosition()
	if size, err = I.bBasic.GetSize(); err != nil {
		return err
	}
	posTXT := new(objects.Position)
	posTXT.X = x + (size.W / 2)
	posTXT.Y = y + (size.H / 2)
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

func (I *Input) IsOver(xRef, yRef int32) bool {
	var x, y int32
	var size *objects.Size
	var err error

	switch I.status {
	case objects.SFix:
		x, y = I.bFix.GetPosition()
		if size, err = I.bFix.GetSize(); err != nil {
			panic(err)
		}
	case objects.SBasic:
		x, y = I.bBasic.GetPosition()
		if size, err = I.bBasic.GetSize(); err != nil {
			panic(err)
		}
	case objects.SOver:
		x, y = I.bOver.GetPosition()
		if size, err = I.bOver.GetSize(); err != nil {
			panic(err)
		}
	case objects.SClick:
		x, y = I.bClick.GetPosition()
		if size, err = I.bClick.GetSize(); err != nil {
			panic(err)
		}
	}

	if xRef > x && xRef < x+size.W {
		if yRef > y && yRef < y+size.H {
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

func (I *Input) UpdatePosition(x, y int32) {
	var diferX, diferY int32
	var blockX, blockY int32

	if I.bFix != nil {
		blockX, blockY = I.bFix.GetPosition()
		diferX = blockX - x
		diferY = blockY - y
		I.bFix.UpdatePosition(x, y)
	}
	if I.bBasic != nil {
		I.bBasic.UpdatePosition(x, y)
	}
	if I.bOver != nil {
		I.bOver.UpdatePosition(x, y)
	}
	if I.bClick != nil {
		I.bClick.UpdatePosition(x, y)
	}

	if I.txt != nil {
		if diferX == 0 && diferY == 0 {
			I.txt.UpdatePosition(x, y)
		} else {
			I.txt.MoveTo(diferX, diferY)
		}
	}
}

func (I *Input) MoveTo(x, y int32) {
	if I.txt != nil {
		I.txt.MoveTo(x, y)
	}
	if I.bFix != nil {
		I.bFix.MoveTo(x, y)
	}
	if I.bBasic != nil {
		I.bBasic.MoveTo(x, y)
	}
	if I.bOver != nil {
		I.bOver.MoveTo(x, y)
	}
	if I.bClick != nil {
		I.bClick.MoveTo(x, y)
	}
}

func (I *Input) GetPosition() (int32, int32) {
	if I.txt != nil {
		return I.txt.GetPosition()
	}
	return -1, -1
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
	I.txt.Draw(wg, r)
}
