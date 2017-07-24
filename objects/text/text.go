package text

import (
	"errors"
	"fmt"
	"sync"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"github.com/ymohl-cl/game-builder/objects"
)

// Text object with implementation of objet interface
type Text struct {
	// infos object
	status      uint8
	initialized bool
	underStyle  uint8

	// content object
	txt        string
	size       int
	color      *objects.Color
	underColor *objects.Color
	position   *objects.Position
	fontURL    string

	// sizeSDL is Width and height of txt on the screen.
	sizeSDL      *objects.Size
	idSDL        uint8
	rect         sdl.Rect
	underRect    sdl.Rect
	texture      *sdl.Texture
	underTexture *sdl.Texture
}

func (T *Text) Init(r *sdl.Renderer) error {
	var err error
	var surface *sdl.Surface
	var uSurface *sdl.Surface
	var font *ttf.Font

	if r == nil {
		return errors.New(objects.ErrorRenderer)
	}

	if T.position == nil {
		return errors.New(objects.ErrorPosition)
	}
	if T.color == nil {
		return errors.New(objects.ErrorColor)
	}
	if T.fontURL == "" {
		return errors.New(objects.ErrorTargetURL)
	}

	if font, err = ttf.OpenFont(T.fontURL, T.size); err != nil {
		return err
	}

	color := sdl.Color{
		R: T.color.Red,
		G: T.color.Green,
		B: T.color.Blue,
		A: T.color.Opacity,
	}
	sdl.Do(func() {
		surface, err = font.RenderUTF8_Solid(T.txt, color)
		if err != nil {
			panic(err)
		}
		defer surface.Free()

		T.sizeSDL = new(objects.Size)
		T.sizeSDL.SetSize(surface.W, surface.H)

		T.rect.X = T.position.X - (T.sizeSDL.W / 2)
		T.rect.Y = T.position.Y - (T.sizeSDL.H / 2)
		T.rect.W = T.sizeSDL.W
		T.rect.H = T.sizeSDL.H

		if T.texture, err = r.CreateTextureFromSurface(surface); err != nil {
			panic(err)
		}

		if T.underColor != nil {
			uColor := sdl.Color{
				R: T.underColor.Red,
				G: T.underColor.Green,
				B: T.underColor.Blue,
				A: T.underColor.Opacity,
			}
			uSurface, err = font.RenderUTF8_Solid(T.txt, uColor)
			if err != nil {
				panic(err)
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
			if T.underTexture, err = r.CreateTextureFromSurface(uSurface); err != nil {
				panic(err)
			}
		}
		T.initialized = true
	})
	return nil
}

// IsInit return status initialize
func (T Text) IsInit() bool {
	return T.initialized
}

func (T *Text) Close() error {
	T.initialized = false
	sdl.Do(func() {
		if T.texture != nil {
			T.texture.Destroy()
		}
		if T.underTexture != nil {
			T.underTexture.Destroy()
		}
	})
	return nil
}

func (T *Text) GetStatus() uint8 {
	return T.status
}

func (T *Text) IsOver(x, y int32) bool {
	if T.status == objects.SFix {
		return false
	}

	return false
}

func (T *Text) Click() {
	return
}

func (T *Text) SetStatus(s uint8) {
	T.status = s
}

func (T *Text) UpdatePosition(x, y int32) {
	if T.position == nil {
		return
	}

	fmt.Println("Update Position")
	fmt.Println("position x: ", x)
	fmt.Println("position y: ", y)
	T.position.X = x
	T.position.Y = y
	T.rect.X = T.position.X - (T.sizeSDL.W / 2)
	T.rect.Y = T.position.Y - (T.sizeSDL.H / 2)

	if T.underColor != nil {
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
	}
}

func (T Text) GetPosition() (int32, int32) {
	return T.position.X, T.position.Y
}

func (T *Text) MoveTo(x, y int32) {
	if T.position == nil {
		return
	}

	fmt.Println("Move TO")
	fmt.Println("position x: ", x)
	fmt.Println("position y: ", y)
	T.position.X += x
	T.position.Y += y
	T.rect.X += x
	T.rect.Y += y

	if T.underColor != nil {
		T.underRect.X += x
		T.underRect.Y += y
	}
}

// Draw the object block.
func (T *Text) Draw(wg *sync.WaitGroup, r *sdl.Renderer) {
	defer wg.Done()

	sdl.Do(func() {
		if T.initialized == false {
			return
		}
		if r == nil {
			panic(errors.New(objects.ErrorRenderer))
		}

		if T.underTexture != nil {
			err := r.SetDrawColor(T.underColor.Red, T.underColor.Green, T.underColor.Blue, T.underColor.Opacity)
			if err != nil {
				panic(err)
			}

			if err = r.Copy(T.underTexture, nil, &T.underRect); err != nil {
				panic(err)
			}
		}

		err := r.SetDrawColor(T.color.Red, T.color.Green, T.color.Blue, T.color.Opacity)
		if err != nil {
			panic(err)
		}
		if err = r.Copy(T.texture, nil, &T.rect); err != nil {
			panic(err)
		}
	})
}
