package text

import (
	"errors"

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

func (T *Text) SetText(str string) {
	T.txt = str
}

func (T *Text) GetIdSDL() uint8 {
	return T.idSDL
}

func (T *Text) NewIDSDL() uint8 {
	var check uint8
	check -= 1

	if T.idSDL == check {
		T.idSDL = 0
	} else {
		T.idSDL++
	}
	return T.idSDL
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

// GetTxt provide txt string
func (T Text) GetTxt() string {
	return T.txt
}
