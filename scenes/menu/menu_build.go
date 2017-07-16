package menu

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/ymohl-cl/game-builder/conf"
	"github.com/ymohl-cl/game-builder/database"
	"github.com/ymohl-cl/game-builder/objects"
	"github.com/ymohl-cl/game-builder/objects/audio"
	"github.com/ymohl-cl/game-builder/objects/block"
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
	if err = M.addStructuresPage(r); err != nil {
		return err
	}
	/*		if err = M.addButtons(); err != nil {
				return err
			}
			if err = M.addPlayers(); err != nil {
				return err
			}*/

	return nil
}

func (M *Menu) addStructuresPage(r *sdl.Renderer) error {
	var b *block.Block
	var err error
	var x, y int32

	// Create blockheader
	b, err = M.addStructurePage(conf.OriginX, conf.MarginTop, conf.WindowWidth, conf.MenuHeaderHeight, conf.ColorBlockRed, conf.ColorBlockGreen, conf.ColorBlockBlue, conf.ColorBlockOpacity)
	if err != nil {
		return err
	}
	if err = b.Init(r); err != nil {
		return nil
	}
	M.layers[layerStructure] = append(M.layers[layerStructure], b)

	// Create blockLeft
	y = conf.MarginTop + conf.MenuHeaderHeight + conf.PaddingBlock
	b, err = M.addStructurePage(conf.MarginLeft, y, conf.MenuContentBlockWidth, conf.MenuContentBlockHeight, conf.ColorBlockRed, conf.ColorBlockGreen, conf.ColorBlockBlue, conf.ColorBlockOpacity)
	if err != nil {
		return err
	}
	if err = b.Init(r); err != nil {
		return nil
	}
	M.layers[layerStructure] = append(M.layers[layerStructure], b)

	// Create blockRight
	x = conf.WindowWidth - conf.MarginRight - conf.MenuContentBlockWidth
	b, err = M.addStructurePage(x, y, conf.MenuContentBlockWidth, conf.MenuContentBlockHeight, conf.ColorBlockRed, conf.ColorBlockGreen, conf.ColorBlockBlue, conf.ColorBlockOpacity)
	if err != nil {
		return err
	}
	if err = b.Init(r); err != nil {
		return nil
	}
	M.layers[layerStructure] = append(M.layers[layerStructure], b)

	// Create blockFooter
	y = conf.WindowHeight - conf.MarginBot - conf.MenuFooterHeight
	b, err = M.addStructurePage(conf.OriginX, y, conf.WindowWidth, conf.MenuHeaderHeight, conf.ColorBlockRed, conf.ColorBlockGreen, conf.ColorBlockBlue, conf.ColorBlockOpacity)
	if err != nil {
		return err
	}
	if err = b.Init(r); err != nil {
		return nil
	}
	M.layers[layerStructure] = append(M.layers[layerStructure], b)
	return nil
}

func (M *Menu) addStructurePage(x, y, w, h int32, red, green, blue, opacity uint8) (*block.Block, error) {
	b, err := block.New(block.Filled)
	if err != nil {
		return nil, err
	}

	p := new(objects.Position)
	p.SetPosition(x, y)
	s := new(objects.Size)
	s.SetSize(w, h)
	c := new(objects.Color)
	c.SetColor(red, green, blue, opacity)

	if err = b.SetPosition(p); err != nil {
		return nil, err
	}
	if err = b.SetSize(s); err != nil {
		return nil, err
	}
	if err = b.SetColor(c); err != nil {
		return nil, err
	}

	return b, nil
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
