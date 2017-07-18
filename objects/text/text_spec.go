package text

import (
	"errors"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/ymohl-cl/game-builder/objects"
)

const (
	PositionTopLeft  = 0
	PositionTopRight = 1
	PositionBotLeft  = 2
	PositionBotRight = 3
)

// New create a new Text object
func New(txt string, size int, font string) (*Text, error) {
	t := new(Text)

	t.status = objects.SFix
	t.txt = txt
	t.size = size
	t.fontURL = font

	return t, nil
}

func (T *Text) SetParams(x, y int32, red, green, blue, opacity uint8) {
	T.position = new(objects.Position)
	T.position.SetPosition(x, y)

	T.color = new(objects.Color)
	T.color.SetColor(red, green, blue, opacity)
}

func (T *Text) SetUnderParams(red, green, blue, opacity uint8, underStyle uint8) {
	T.underColor = new(objects.Color)
	T.underColor.SetColor(red, green, blue, opacity)

	T.underStyle = underStyle
}

// SetSize
func (T *Text) SetSize(sz int) error {
	if sz == 0 {
		return errors.New(objects.ErrorSize)
	}

	T.size = sz
	return nil
}

// SetPosition
func (T *Text) SetPosition(p *objects.Position) error {
	if p == nil {
		return errors.New(objects.ErrorPosition)
	}

	T.position = p
	return nil
}

// SetColor
func (T *Text) SetColor(c *objects.Color) error {
	if c == nil {
		return errors.New(objects.ErrorColor)
	}

	T.color = c
	return nil
}

// SetUnderColor
func (T *Text) SetUnderColor(c *objects.Color) error {
	if c == nil {
		return errors.New(objects.ErrorColor)
	}

	T.underColor = c
	return nil
}

// SetFontTTF
func (T *Text) SetFontTTF(url string) error {
	if url == "" {
		return errors.New(objects.ErrorTargetURL)
	}

	T.fontURL = url
	return nil
}

func (T *Text) UpdateTxt(str string) error {
	if T.size <= 0 {
		return errors.New(objects.ErrorSize)
	}
	if T.position == nil {
		return errors.New(objects.ErrorPosition)
	}
	if T.color == nil {
		return errors.New(objects.ErrorColor)
	}
	if T.font == nil {
		return errors.New(objects.ErrorEmpty)
	}

	T.txt = str
	color := sdl.Color{
		R: T.color.Red,
		G: T.color.Green,
		B: T.color.Blue,
		A: T.color.Opacity,
	}
	surface, err := T.font.RenderUTF8_Solid(T.txt, color)
	if err != nil {
		return err
	}
	defer surface.Free()

	T.sizeSDL.SetSize(surface.W, surface.H)

	T.rect.X = T.position.X - (T.sizeSDL.W / 2)
	T.rect.Y = T.position.Y - (T.sizeSDL.H / 2)
	T.rect.W = T.sizeSDL.W
	T.rect.H = T.sizeSDL.H

	if T.texture != nil {
		T.texture.Destroy()
	}
	if T.texture, err = T.renderer.CreateTextureFromSurface(surface); err != nil {
		return err
	}

	if T.underColor != nil {
		uColor := sdl.Color{
			R: T.underColor.Red,
			G: T.underColor.Green,
			B: T.underColor.Blue,
			A: T.underColor.Opacity,
		}
		uSurface, err := T.font.RenderUTF8_Solid(T.txt, uColor)
		if err != nil {
			return err
		}
		defer uSurface.Free()
		if T.underStyle == PositionTopLeft || T.underStyle == PositionBotRight {
			T.underRect.Y = T.rect.Y - 1
		} else {
			T.underRect.Y = T.rect.Y + 1
		}
		if T.underStyle == PositionTopRight || T.underStyle == PositionBotRight {
			T.underRect.X = T.rect.X + 1
		} else {
			T.underRect.X = T.rect.X - 1
		}
		T.underRect.W = T.rect.W
		T.underRect.H = T.rect.H
		if T.underTexture != nil {
			T.underTexture.Destroy()
		}
		if T.underTexture, err = T.renderer.CreateTextureFromSurface(uSurface); err != nil {
			return err
		}
	}

	return nil
}

// GetSize provide size object (on the screen)
func (T Text) GetSize() (*objects.Size, error) {
	if T.sizeSDL == nil {
		return nil, errors.New(objects.ErrorSize)
	}
	return T.sizeSDL, nil
}

// GetPosition provide position object
func (T Text) GetPosition() (*objects.Position, error) {
	if T.position == nil {
		return nil, errors.New(objects.ErrorPosition)
	}
	return T.position, nil
}

// GetColor provide color object
func (T Text) GetColor() (*objects.Color, error) {
	if T.color == nil {
		return nil, errors.New(objects.ErrorColor)
	}
	return T.color, nil
}

// GetUnderColor provide color object
func (T Text) GetUnderColor() (*objects.Color, error) {
	if T.underColor == nil {
		return nil, errors.New(objects.ErrorColor)
	}
	return T.underColor, nil
}
