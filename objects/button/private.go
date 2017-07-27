package button

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/ymohl-cl/game-builder/objects"
)

/*
** Private method
 */
func (B *Button) setParamsByStatus() {
	switch B.status {
	case objects.SFix:
		B.txt.SetParams(B.fix.txtRect.X, B.fix.txtRect.Y, B.fix.txtColor.R, B.fix.txtColor.G, B.fix.txtColor.B, B.fix.txtColor.A)
		B.block.SetParams(B.fix.blockRect.X, B.fix.blockRect.Y, B.fix.blockRect.W, B.fix.blockRect.H, B.fix.blockColor.R, B.fix.blockColor.G, B.fix.blockColor.B, B.fix.blockColor.A)
		if B.fix.img != nil {
			B.fix.img.SetParams(B.fix.blockRect.X, B.fix.blockRect.Y, B.fix.blockRect.W, B.fix.blockRect.H)
		}
	case objects.SBasic:
		B.txt.SetParams(B.basic.txtRect.X, B.basic.txtRect.Y, B.basic.txtColor.R, B.basic.txtColor.G, B.basic.txtColor.B, B.basic.txtColor.A)
		B.block.SetParams(B.basic.blockRect.X, B.basic.blockRect.Y, B.basic.blockRect.W, B.basic.blockRect.H, B.basic.blockColor.R, B.basic.blockColor.G, B.basic.blockColor.B, B.basic.blockColor.A)
		if B.basic.img != nil {
			B.basic.img.SetParams(B.basic.blockRect.X, B.basic.blockRect.Y, B.basic.blockRect.W, B.basic.blockRect.H)
		}
	case objects.SOver:
		B.txt.SetParams(B.over.txtRect.X, B.over.txtRect.Y, B.over.txtColor.R, B.over.txtColor.G, B.over.txtColor.B, B.over.txtColor.A)
		B.block.SetParams(B.over.blockRect.X, B.over.blockRect.Y, B.over.blockRect.W, B.over.blockRect.H, B.over.blockColor.R, B.over.blockColor.G, B.over.blockColor.B, B.over.blockColor.A)
		if B.over.img != nil {
			B.over.img.SetParams(B.over.blockRect.X, B.over.blockRect.Y, B.over.blockRect.W, B.over.blockRect.H)
		}
	case objects.SClick:
		B.txt.SetParams(B.click.txtRect.X, B.click.txtRect.Y, B.click.txtColor.R, B.click.txtColor.G, B.click.txtColor.B, B.click.txtColor.A)
		B.block.SetParams(B.click.blockRect.X, B.click.blockRect.Y, B.click.blockRect.W, B.click.blockRect.H, B.click.blockColor.R, B.click.blockColor.G, B.click.blockColor.B, B.click.blockColor.A)
		if B.click.img != nil {
			B.click.img.SetParams(B.click.blockRect.X, B.click.blockRect.Y, B.click.blockRect.W, B.click.blockRect.H)
		}
	}
}

func (B *Button) updateParamsByStatus() {
	switch B.status {
	case objects.SFix:
		B.txt.UpdateColor(B.fix.txtColor.R, B.fix.txtColor.G, B.fix.txtColor.B, B.fix.txtColor.A)
		B.block.UpdateColor(B.fix.blockColor.R, B.fix.blockColor.G, B.fix.blockColor.B, B.fix.blockColor.A)
		B.txt.UpdatePosition(B.fix.txtRect.X, B.fix.txtRect.W)
		B.block.UpdatePosition(B.fix.blockRect.X, B.fix.blockRect.W)
		if B.fix.img != nil {
			B.fix.img.UpdatePosition(B.fix.blockRect.X, B.fix.blockRect.Y)
		}
	case objects.SBasic:
		B.txt.UpdateColor(B.basic.txtColor.R, B.basic.txtColor.G, B.basic.txtColor.B, B.basic.txtColor.A)
		B.block.UpdateColor(B.basic.blockColor.R, B.basic.blockColor.G, B.basic.blockColor.B, B.basic.blockColor.A)
		B.txt.UpdatePosition(B.basic.txtRect.X, B.basic.txtRect.W)
		B.block.UpdatePosition(B.basic.blockRect.X, B.basic.blockRect.W)
		if B.basic.img != nil {
			B.basic.img.UpdatePosition(B.basic.blockRect.X, B.basic.blockRect.Y)
		}
	case objects.SOver:
		B.txt.UpdateColor(B.over.txtColor.R, B.over.txtColor.G, B.over.txtColor.B, B.over.txtColor.A)
		B.block.UpdateColor(B.over.blockColor.R, B.over.blockColor.G, B.over.blockColor.B, B.over.blockColor.A)
		B.txt.UpdatePosition(B.over.txtRect.X, B.over.txtRect.W)
		B.block.UpdatePosition(B.over.blockRect.X, B.over.blockRect.W)
		if B.over.img != nil {
			B.over.img.UpdatePosition(B.over.blockRect.X, B.over.blockRect.Y)
		}
	case objects.SClick:
		B.txt.UpdateColor(B.click.txtColor.R, B.click.txtColor.G, B.click.txtColor.B, B.click.txtColor.A)
		B.block.UpdateColor(B.click.blockColor.R, B.click.blockColor.G, B.click.blockColor.B, B.click.blockColor.A)
		B.txt.UpdatePosition(B.click.txtRect.X, B.click.txtRect.W)
		B.block.UpdatePosition(B.click.blockRect.X, B.click.blockRect.W)
		if B.click.img != nil {
			B.click.img.UpdatePosition(B.click.blockRect.X, B.click.blockRect.Y)
		}
	}
}

func (B *Button) updatePositionByStatus() {
	switch B.status {
	case objects.SFix:
		B.txt.UpdatePosition(B.fix.txtRect.X, B.fix.txtRect.W)
		B.block.UpdatePosition(B.fix.blockRect.X, B.fix.blockRect.W)
		if B.fix.img != nil {
			B.fix.img.UpdatePosition(B.fix.blockRect.X, B.fix.blockRect.Y)
		}
	case objects.SBasic:
		B.txt.UpdatePosition(B.basic.txtRect.X, B.basic.txtRect.W)
		B.block.UpdatePosition(B.basic.blockRect.X, B.basic.blockRect.W)
		if B.basic.img != nil {
			B.basic.img.UpdatePosition(B.basic.blockRect.X, B.basic.blockRect.Y)
		}
	case objects.SOver:
		B.txt.UpdatePosition(B.over.txtRect.X, B.over.txtRect.W)
		B.block.UpdatePosition(B.over.blockRect.X, B.over.blockRect.W)
		if B.over.img != nil {
			B.over.img.UpdatePosition(B.over.blockRect.X, B.over.blockRect.Y)
		}
	case objects.SClick:
		B.txt.UpdatePosition(B.click.txtRect.X, B.click.txtRect.W)
		B.block.UpdatePosition(B.click.blockRect.X, B.click.blockRect.W)
		if B.click.img != nil {
			B.click.img.UpdatePosition(B.click.blockRect.X, B.click.blockRect.Y)
		}
	}
}

func (B *Button) initImages(r *sdl.Renderer) error {
	var err error

	if B.fix.img != nil {
		if err = B.fix.img.Init(r); err != nil {
			return err
		}
	}
	if B.basic.img != nil {
		if err = B.basic.img.Init(r); err != nil {
			return err
		}
	}
	if B.over.img != nil {
		if err = B.over.img.Init(r); err != nil {
			return err
		}
	}
	if B.click.img != nil {
		if err = B.click.img.Init(r); err != nil {
			return err
		}
	}

	return nil
}

func (B *Button) closeImages() error {
	var err error

	if B.fix.img != nil {
		if err = B.fix.img.Close(); err != nil {
			return err
		}
	}
	if B.basic.img != nil {
		if err = B.basic.img.Close(); err != nil {
			return err
		}
	}
	if B.over.img != nil {
		if err = B.over.img.Close(); err != nil {
			return err
		}
	}
	if B.click.img != nil {
		if err = B.click.img.Close(); err != nil {
			return err
		}
	}

	return nil
}
