package buttun

import (
	"errors"
	"sync"

	"github.com/42MrPiou42/game-builder/objects"
	"github.com/42MrPiou42/game-builder/objects/block"
	"github.com/veandco/go-sdl2/sdl"
)

type Buttun struct {
	// infos object
	status      uint8
	initialized bool

	// content object
	cBasic    Content
	cOver     Content
	cClick    Content
	cFix      Content
	funcClick func(...interface{}) string
	dataClick []interface

	// object of sdl
}

type Content struct {
	contentTxt *text.Text
	contentImg *images.Img
	block      *block.Block
}

/*
** Functions buttun specifications
 */
 // New create a new Buttun object
func New() *Buttun {
	b := new(Buttun)

	b.status = objects.SBasic
	return b
}

// SetTxt to the status specified
func (B *Buttun) SetTxt(t *text.Text, s uint8) error {
	if !t {
		return errors.New("Can't add txt because is nil")
	}

	switch s {
	case objects.SFix:
		B.cFix.contentTxt = t
	case objects.SBasic:
		B.cBasic.contentTxt = t
	case objects.SOver:
		B.cOver.contentTxt = t
	case objects.SClick:
		B.cClick.contentTxt = t
	default:
		return errors.New("Status not available")
	}
	return nil
}

// SetSize to the status specified
func (B *Buttun) SetImg(i *images.Img, s uint8) error {
	if !i {
		return errors.New("Can't add img because is nil")
	}

	switch s {
	case objects.SFix:
		B.cFix.contentImg = i
	case objects.SBasic:
		B.cBasic.contentImg = i
	case objects.SOver:
		B.cOver.contentImg = i
	case objects.SClick:
		B.cClick.contentImg = i
	default:
		return errors.New("Status not available")
	}
	return nil
}

// SetSize to the status specified
func (B *Buttun) SetSize(sz *objects.Size, st uint8) error {
	var err error

	if !s {
		return errors.New("Can't add size because is nil")
	}

	switch st {
	case objects.SFix:
		err = B.cFix.block.SetSize(sz) = sz
	case objects.SBasic:
		err = B.cBasic.block.SetSize(sz)
	case objects.SOver:
		err = B.cOver.block.SetSize(sz)
	case objects.SClick:
		err = B.cClick.block.SetSize(sz)
	default:
		return errors.New("Status not available")
	}
	return err
}

// SetPosition to the status specified
func (B *Buttun) SetPosition(p *objects.Position, s uint8) error {
	var err error

	if !p {
		return errors.New("Can't add position because is nil")
	}

	switch s {
	case objects.SFix:
		err = B.cFix.block.SetPosition(p)
	case objects.SBasic:
		err = B.cBasic.block.SetPostion(p)
	case objects.SOver:
		err = B.cOver.block.SetPosition(p)
	case objects.SClick:
		err = B.cClick.block.Setposition(p)
	default:
		return errors.New("Status not available")
	}
	return err
}


// SetAction define action when the element is click
func (B *Buttun) SetAction(f func(...interface{})string, d []interface{}) {
	B.funcClick = f
	B.dataClick = d
}

// SetColor to the status specified
func (B *Buttun) SetColor(c *objects.Color, s uint8) error {
	var err error

	if !c {
		return errors.New("Can't add color because is nil")
	}

	switch s {
	case objects.SFix:
		err = B.cFix.SetColor(c)
	case objects.SBasic:
		err = B.cBasic.SetColor(c)
	case objects.SOver:
		err = B.cOver.SetColor(c)
	case objects.SClick:
		err = B.cClick.SetColor(c)
	default:
		return errors.New("Status not available")
	}
	return err
}

func (B *Buttun) CopyStateToStates(stateSource uint8, stDests []uint8) error {
	var source Content{}

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
		return errors.New("Status not available")
	}

	for _, v := range stDests {
		switch v {
		case objects.SFix:
			copy(B.cFix, source)
		case objects.SBasic:
			copy(B.cBasic, source)
		case objects.SOver:
			copy(B.cOver, source)
		case objects.SClick:
			copy(B.cClick, source)
		default:
			return errors.New("Status to dest copy not available")
		}
	}
	return nil
}


/*
** Interface objects functions
 */

// IsInit return status initialize
func (B *Buttun) IsInit() bool {
	return B.initialized
}

func (B *Buttun) Init(r *sdl.Renderer) error {
	if B.status == objects.SFix {
		if err := B.cFix.checkContent(); err != nil {
			return err
		}
		if err := B.cFix.initContent(r); err != nil {
			return err
		}
	} else {
		if B.funcClick == nil {
			return errors.New("Function not define to the buttun")
		}
		if err := B.cBasic.checkContent(); err != nil {
			return err
		}
		if err := B.cBasic.initContent(r); err != nil {
			return err
		}

		if err := B.cOver.checkContent(); err != nil {
			return err
		}
		if err := B.cOver.initContent(r); err != nil {
			return err
		}

		if err := B.cClick.checkContent(); err != nil {
			return err
		}
		if err := B.cClick.initContent(r); err != nil {
			return err
		}
	}

	B.initialized = true
}

func (B *Buttun) Close() error {
	if B.status == objects.SFix {
		if err := B.cFix.closeContent(); err != nil {
			return err
		}
	} else {
		if err := B.cBasic.closeContent(); err != nil {
			return err
		}
		if err := B.cBasic.closeContent(r); err != nil {
			return err
		}

		if err := B.cOver.closeContent(); err != nil {
			return err
		}
		if err := B.cOver.closeContent(r); err != nil {
			return err
		}

		if err := B.cClick.closeContent(); err != nil {
			return err
		}
		if err := B.cClick.closeContent(r); err != nil {
			return err
		}
	}

	B.initialized = false
	return nil
}

func (B *Buttun) GetStatus() uint8 {
	return B.status
}

func (B *Buttun) IsOver(x, y int32) bool {
	var pos *objects.Position
	var size *objects.Size

	switch B.status {
	case objects.SFix:
		pos = B.cFix.block.GetPosition()
		size = B.cFix.block.GetSize()
	case objects.SBasic:
		pos = B.cBasic.block.GetPosition()
		size = B.cBasic.block.GetSize()
	case objects.SOver:
		pos = B.cOver.block.GetPosition()
		size = B.cOver.block.GetSize()
	case objects.Click:
		pos = B.cClick.block.GetPosition()
		size = B.cClick.block.GetSize()
	}

	if !pos || size {
		return false
	}
	if x > pos.X && x < pos.X+size.W {
		if y > pos.Y && y < pos.Y+size.H {
			return true
		}
	}
	return false
}

func (B *Buttun) Click() {
	B.SetStatus(objects.SClick)
	B.funcClick(B.dataClick)
}

func (B *Buttun) SetStatus(s uint8) {
	if B.status != objects.SFix {
		B.status = s
	}
}

func (B *Buttun) Draw(r *sdl.Renderer, wg *sync.WaitGroup) error {
	if r == nil {
		return errors.New("Can't draw buttun because renderer is nil")
	}
	if wg == nil {
		return errors.New("Can't draw buttun because sync WaitGroup not define")
	}
	if B.initialized == false {
		return errors.New("Can't draw because object is not initialized")
	}

	wg.Add(1)
	defer wg.Done()

	switch B.status {
	case objects.SFix:
		if err := B.cFix.drawContent(); err != nil {
			return err
		}
	case objects.Sbasic:
		if err := B.cBasic.drawContent(); err != nil {
			return err
		}
	case objects.SOver:
		if err := B.cOver.drawContent(); err != nil {
			return err
		}
	case objects.SClick:
		if err := B.cClick.drawContent(); err != nil {
			return err
		}
	}

	return nil
}


/*
** Private function Text objects
 */
// checkContent and return err with the raison.
func (C Content) checkContent() error {
	var flag uint8
	var err error

	if C.block {
		flag++
	}
	if C.contentImg {
		flag++
	}
	if C.contentTxt {
		flag++
	}

	if flag == 0 {
		return errors.New("Buttun isn't define by one Content")
	}
	return nil
}

func (C *Content) initContent() error {
	var err error

	if C.block {
		if C.block.IsInit() == false  {
			err = C.block.Init()
			if err != nil {
				return err
			}
		}
	}
	if C.contentImg {
		if C.contentImg.IsInit() == false  {
			err = C.contentImg.Init()
			if err != nil {
				return err
			}
		}
	}
	if C.contentTxt {
		if C.contentTxt.IsInit() == false  {
			err = C.contentTxt.Init()
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (C *Content) closeContent() error {
	if C.block {
		if err := C.block.Close(); err != nil {
			return err
		}
	}
	if C.contentImg {
		if err := C.contentImg.Close(); err != nil {
			return err
		}
	}
	if C.contentTxt {
		if err := C.contentTxt.Close(); err != nil {
			return err
		}
	}
	return nil
}

func (C Content) drawContent() error {
	if C.block {
		if err := C.block.Draw(); err != nil {
			return err
		}
	}
	if C.contentImg {
		if err := C.contentImg.Draw(); err != nil {
			return err
		}
	}
	if C.contentTxt {
		if err := C.contentTxt.Draw(); err != nil {
			return err
		}
	}
	return nil
}
