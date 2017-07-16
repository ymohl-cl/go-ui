package menu

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/ymohl-cl/game-builder/conf"
	"github.com/ymohl-cl/game-builder/database"
	"github.com/ymohl-cl/game-builder/objects"
	"github.com/ymohl-cl/game-builder/objects/audio"
	"github.com/ymohl-cl/game-builder/objects/image"
)

func (M *Menu) build(d *database.Data, r *sdl.Renderer) error {
	var err error

	if err = M.addMusic(r); err != nil {
		return err
	}
	if err = M.addBackground(r); err != nil {
		return err
	}
	/*	if err = M.addStructurePage(); err != nil {
			return err
		}
		if err = M.addButtons(); err != nil {
			return err
		}
		if err = M.addPlayers(); err != nil {
			return err
		}*/

	return nil
}

func (M *Menu) addBackground(r *sdl.Renderer) error {
	i, err := image.New(conf.MenuBackground)
	if err != nil {
		return err
	}

	p := new(objects.Position)
	p.SetPosition(conf.OriginX, conf.OriginY)
	s := new(objects.Size)
	s.SetSize(conf.WindowWidth, conf.WindowHeight)

	if err = i.SetPosition(p); err != nil {
		return err
	}
	if err = i.SetSize(s); err != nil {
		return err
	}

	if err = i.Init(r); err != nil {
		return err
	}
	M.layers[layerBackground] = append(M.layers[layerBackground], i)
	return nil
}

func (M *Menu) addMusic(r *sdl.Renderer) error {
	var err error

	M.music, err = audio.New(conf.MenuMusic)
	if err != nil {
		return err
	}

	if err = M.music.Init(r); err != nil {
		return err
	}
	return nil
}
