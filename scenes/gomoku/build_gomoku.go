package gomoku

import (
	"github.com/ymohl-cl/game-builder/audio"
	"github.com/ymohl-cl/game-builder/conf"
	"github.com/ymohl-cl/game-builder/objects/image"
)

func (G *Gomoku) addMusic() error {
	var err error

	G.music, err = audio.New(conf.GameMusic, 3)
	if err != nil {
		return err
	}

	if err = G.music.Init(G.renderer); err != nil {
		return err
	}
	return nil
}

func (G *Gomoku) addBackground() error {
	var i *image.Image
	var err error

	i = image.New(conf.MenuBackground, conf.OriginX, conf.OriginY, conf.WindowWidth, conf.WindowHeight)
	if err = i.Init(G.renderer); err != nil {
		return err
	}
	G.layers[layerBackground] = append(G.layers[layerBackground], i)
	return nil
}
