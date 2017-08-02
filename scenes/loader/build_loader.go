package loader

import (
	"github.com/ymohl-cl/game-builder/audio"
	"github.com/ymohl-cl/game-builder/conf"
	"github.com/ymohl-cl/game-builder/objects"
	"github.com/ymohl-cl/game-builder/objects/block"
	"github.com/ymohl-cl/game-builder/objects/image"
)

func (L *Load) addMusic() error {
	var err error

	L.music, err = audio.New(conf.LoadMusic)
	if err != nil {
		return err
	}

	if err = L.music.Init(L.renderer); err != nil {
		return err
	}
	return nil
}

func (L *Load) addBackground() error {
	var i *image.Image
	var err error

	i = image.New(conf.MenuBackground, conf.OriginX, conf.OriginY, conf.WindowWidth, conf.WindowHeight)
	if err = i.Init(L.renderer); err != nil {
		return err
	}
	L.layers[layerBackground] = append(L.layers[layerBackground], i)
	return nil
}

func (L *Load) addBlockLoading() error {
	var b *block.Block
	var err error

	if b, err = block.New(block.Filled); err != nil {
		return err
	}
	b.SetVariantStyle(conf.ColorBlockRed, conf.ColorBlockGreen, conf.ColorBlockBlue, conf.ColorBlockOpacity, objects.SFix)
	b.UpdateSize(conf.LoadBlockWidth, conf.LoadBlockHeight)
	b.UpdatePosition(conf.OriginX, conf.OriginY+(conf.WindowHeight/2))

	if err = b.Init(L.renderer); err != nil {
		return err
	}
	L.layers[layerLoadingBar] = append(L.layers[layerLoadingBar], b)
	L.lastLoadBlock = b
	return nil
}
