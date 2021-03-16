package text

import "github.com/veandco/go-sdl2/sdl"

func (T *Text) updateTextures(r *sdl.Renderer) error {
	var err error
	var surface *sdl.Surface

	for id, color := range T.colors {
		if surface, err = T.font.RenderUTF8Solid(T.txt, color); err != nil {
			return err
		}
		defer surface.Free()

		if T.textures[id], err = r.CreateTextureFromSurface(surface); err != nil {
			return err
		}
	}
	if surface != nil {
		T.rect.W = surface.W
		T.rect.H = surface.H
		surface = nil
	}

	for id, color := range T.underColors {
		if surface, err = T.font.RenderUTF8Solid(T.txt, color); err != nil {
			return err
		}
		defer surface.Free()

		if T.underTextures[id], err = r.CreateTextureFromSurface(surface); err != nil {
			return err
		}
	}
	if surface != nil {
		T.underRect.W = surface.W
		T.underRect.H = surface.H
	}
	return err
}

func (T *Text) updateTextureByStatus(s uint8, r *sdl.Renderer) error {
	var err error
	var surface *sdl.Surface

	T.textures[s].Destroy()
	if surface, err = T.font.RenderUTF8Solid(T.txt, T.colors[s]); err != nil {
		return err
	}
	defer surface.Free()

	if T.textures[s], err = r.CreateTextureFromSurface(surface); err != nil {
		return err
	}
	return nil
}
