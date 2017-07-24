package button

import (
	"errors"
	"fmt"
	"sync"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/ymohl-cl/game-builder/objects"
	"github.com/ymohl-cl/game-builder/objects/block"
	"github.com/ymohl-cl/game-builder/objects/image"
	"github.com/ymohl-cl/game-builder/objects/text"
)

type Content struct {
	txt   *text.Text
	img   *image.Image
	block *block.Block
}

// New create a new Button object
func New(f func(...interface{}), d ...interface{}) *Button {
	b := new(Button)

	b.status = objects.SBasic
	b.funcClick = f
	for _, v := range d {
		b.dataClick = append(b.dataClick, v)
	}
	return b
}

func (B *Button) SetContentBasic(t *text.Text, i *image.Image, b *block.Block) {
	B.cBasic.txt = t
	B.cBasic.img = i
	B.cBasic.block = b
}

func (B *Button) SetContentOver(t *text.Text, i *image.Image, b *block.Block) {
	B.cOver.txt = t
	B.cOver.img = i
	B.cOver.block = b
}

func (B *Button) SetContentClick(t *text.Text, i *image.Image, b *block.Block) {
	B.cClick.txt = t
	B.cClick.img = i
	B.cClick.block = b
}

func (B *Button) SetContentFix(t *text.Text, i *image.Image, b *block.Block) {
	B.cFix.txt = t
	B.cFix.img = i
	B.cFix.block = b
}

func (B *Button) CopyStateToStates(stateSource uint8, stDests []uint8) error {
	var source Content

	switch stateSource {
	case objects.SFix:
		source = B.cFix
	case objects.SBasic:
		source = B.cBasic
	case objects.SOver:
		source = B.cOver
	case objects.SClick:
		source = B.cClick
	default:
		return errors.New(objects.ErrorStatus)
	}

	for _, v := range stDests {
		switch v {
		case objects.SFix:
			B.cFix.copyContent(source)
		case objects.SBasic:
			B.cBasic.copyContent(source)
		case objects.SOver:
			B.cOver.copyContent(source)
		case objects.SClick:
			B.cClick.copyContent(source)
		default:
			return errors.New(objects.ErrorStatus)
		}
	}
	return nil
}

/*
** Private function Text objects
 */
func (C Content) copyContent(s Content) {
	C.txt = s.txt
	C.img = s.img
	C.block = s.block
}

// checkContent and return err with the raison.
func (C Content) checkContent() error {
	var flag uint8

	if C.block != nil {
		flag++
	}
	if C.img != nil {
		flag++
	}
	if C.txt != nil {
		flag++
	}

	if flag == 0 {
		return errors.New(objects.ErrorEmpty)
	}
	return nil
}

func (C *Content) initContent(r *sdl.Renderer) error {
	var err error

	if C.block != nil {
		if C.block.IsInit() == false {
			if err = C.block.Init(r); err != nil {
				return err
			}
		}
	}
	if C.img != nil {
		if C.img.IsInit() == false {
			if err = C.img.Init(r); err != nil {
				return err
			}
		}
	}
	if C.txt != nil {
		if C.txt.IsInit() == false {
			if err = C.txt.Init(r); err != nil {
				return err
			}
		}
	}
	return nil
}

func (C *Content) closeContent() error {
	var err error

	if C.block != nil {
		if err = C.block.Close(); err != nil {
			return err
		}
	}
	if C.img != nil {
		if err = C.img.Close(); err != nil {
			return err
		}
	}
	if C.txt != nil {
		if err = C.txt.Close(); err != nil {
			return err
		}
	}
	return nil
}

func (C Content) drawContent(wg *sync.WaitGroup, r *sdl.Renderer) {

	if C.block != nil {
		wg.Add(1)
		C.block.Draw(wg, r)
	}
	if C.img != nil {
		wg.Add(1)
		C.img.Draw(wg, r)
	}
	if C.txt != nil {
		wg.Add(1)
		C.txt.Draw(wg, r)
	}
}

func (C Content) getPosition() (*objects.Position, error) {
	var x, y int32

	if C.block != nil {
		x, y = C.block.GetPosition()
	} else if C.img != nil {
		x, y = C.img.GetPosition()
	} else if C.txt != nil {
		x, y = C.txt.GetPosition()
	} else {
		return nil, errors.New(objects.ErrorEmpty)
	}
	return &objects.Position{X: x, Y: y}, nil
}

func (C Content) getSize() (*objects.Size, error) {
	if C.block != nil {
		return C.block.GetSize()
	}
	if C.img != nil {
		return C.img.GetSize()
	}
	if C.txt != nil {
		return C.txt.GetSize()
	}
	return nil, errors.New(objects.ErrorEmpty)
}

func (C *Content) UpdatePosition(x, y int32) {
	var diferX, diferY int32
	var blockX, blockY int32
	var imgX, imgY int32

	if C.block != nil {
		blockX, blockY = C.block.GetPosition()
		fmt.Println("Block position X: ", blockX, " | Y: ", blockY)
		C.block.UpdatePosition(x, y)
		diferX = x - blockX
		diferY = y - blockY
		fmt.Println("Block difer position X: ", diferX, " | Y: ", diferY)
	}
	if C.img != nil {
		if diferX == 0 && diferY == 0 {
			imgX, imgY = C.img.GetPosition()
			C.img.UpdatePosition(x, y)
			diferX = imgX - x
			diferY = imgY - y
		} else {
			C.img.MoveTo(diferX, diferY)
		}
	}
	if C.txt != nil {
		if diferX == 0 && diferY == 0 {
			fmt.Println("Call update from Content")
			fmt.Println("position x: ", x)
			fmt.Println("position y: ", y)
			fmt.Println("difer x: ", diferX)
			fmt.Println("difer y: ", diferY)
			fmt.Println(".............................")
			C.txt.UpdatePosition(x, y)
		} else {
			fmt.Println("Call move to from Content")

			fmt.Println("position x: ", x)
			fmt.Println("position y: ", y)
			fmt.Println("difer x: ", diferX)
			fmt.Println("difer y: ", diferY)
			fmt.Println(".............................")
			C.txt.MoveTo(diferX, diferY)
		}
	}
	return
}

func (C *Content) MoveTo(x, y int32) {
	if C.block != nil {
		C.block.MoveTo(x, y)
	}
	if C.img != nil {
		C.img.MoveTo(x, y)
	}
	if C.txt != nil {
		C.txt.MoveTo(x, y)
	}
}
