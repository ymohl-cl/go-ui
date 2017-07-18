package button

import (
	"errors"
	"image"
	"sync"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/ymohl-cl/game-builder/objects"
	"github.com/ymohl-cl/game-builder/objects/block"
	"github.com/ymohl-cl/game-builder/objects/text"
)

type Content struct {
	txt   *text.Text
	img   *image.Image
	block *block.Block
}

// New create a new Button object
func New(f func(...interface{}) string, p []interface{}) *Button {
	b := new(Button)

	b.status = objects.SBasic
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
	var err error

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

func (C Content) drawContent(wg *sync.WaitGroup) {
	defer wg.Done()

	if C.block != nil {
		C.block.Draw(wg)
	}
	if C.img != nil {
		C.img.Draw(wg)
	}
	if C.txt != nil {
		C.txt.Draw(wg)
	}
}

func (C Content) getPosition() (*objects.Position, error) {
	if C.block != nil {
		return C.block.GetPosition()
	}
	if C.img != nil {
		return C.img.GetPosition()
	}
	if C.txt != nil {
		return C.txt.GetPosition()
	}
	return nil, errors.New(objects.ErrorEmpty)
}

func (C Content) getSize() (*objects.Size, error) {
	if C.block {
		return C.block.GetSize(), nil
	}
	if C.img {
		return C.img.GetSize(), nil
	}
	if C.txt {
		return C.txt.GetSize(), nil
	}
	return nil, errors.New(objects.ErrorEmpty)
}
