package button

type Content struct {
	contentTxt *text.Text
	contentImg *images.Img
	block      *block.Block
}

// New create a new Button object
func New(r *sdl.Renderer) *Button {
	b := new(Button)

	b.renderer = r
	b.status = objects.SBasic
	return b
}

func (B *Button) SetContentBasic(t *text.Text, i *images.Img, b *block.Block) {
	B.cBasic.contentTxt = t
	B.cBasic.contentImg = i
	B.cBasic.block = b
}

func (B *Button) SetContentOver(t *text.Text, i *images.Img, b *block.Block) {
	B.cOver.contentTxt = t
	B.cOver.contentImg = i
	B.cOver.block = b
}

func (B *Button) SetContentClick(t *text.Text, i *images.Img, b *block.Block) {
	B.cClick.contentTxt = t
	B.cClick.contentImg = i
	B.cClick.block = b
}

func (B *Button) SetContentFix(t *text.Text, i *images.Img, b *block.Block) {
	B.cFix.contentTxt = t
	B.cFix.contentImg = i
	B.cFix.block = b
}

// SetTxt to the status specified
func (B *Button) SetTxt(t *text.Text, s uint8) error {
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
func (B *Button) SetImg(i *images.Img, s uint8) error {
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
func (B *Button) SetSize(sz *objects.Size, st uint8) error {
	var err error

	if !s {
		return errors.New("Can't add size because is nil")
	}

	switch st {
	case objects.SFix:
		err = B.cFix.block.SetSize(sz)
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
func (B *Button) SetPosition(p *objects.Position, s uint8) error {
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
func (B *Button) SetAction(f func(...interface{})string, d []interface{}) {
	B.funcClick = f
	B.dataClick = d
}

// SetColor to the status specified
func (B *Button) SetColor(c *objects.Color, s uint8) error {
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

func (B *Button) CopyStateToStates(stateSource uint8, stDests []uint8) error {
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
		return errors.New("Button isn't define by one Content")
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
