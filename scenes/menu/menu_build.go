package menu

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/ymohl-cl/game-builder/conf"
	"github.com/ymohl-cl/game-builder/database"
	"github.com/ymohl-cl/game-builder/objects"
	"github.com/ymohl-cl/game-builder/objects/audio"
	"github.com/ymohl-cl/game-builder/objects/block"
	"github.com/ymohl-cl/game-builder/objects/button"
	"github.com/ymohl-cl/game-builder/objects/image"
	"github.com/ymohl-cl/game-builder/objects/text"
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
	if err = M.addButtons(r, d); err != nil {
		return err
	}
	/*			if err = M.addPlayers(); err != nil {
				return err
			}*/

	return nil
}

func (M *Menu) addButtons(r *sdl.Renderer, d *database.Data) error {
	var x, y int32

	// Create button create new Player
	y = conf.MarginTop + conf.MenuHeaderHeight + conf.PaddingBlock + conf.MenuContentMediumBlockHeight/2 + conf.PaddingBlock/2
	x = conf.WindowWidth - conf.MarginRight - (conf.MenuContentBlockWidth - (conf.ButtonWidth * 2 + conf.PaddingBlock)
	t, err := text.New("NEW PLAYER", conf.TxtMedium, conf.Font)
	c := new(objects.Color)
	c.SetColor(conf.ColorTxtRed, conf.ColorTxtGreen, conf.ColorTxtBlue, conf.ColorTxtOpacity)
	underC := new(objects.Color)
	underC.SetColor(conf.ColorUnderTxtRed, conf.ColorUnderTxtGreen, conf.ColorUnderTxtBlue, conf.ColorUnderTxtOpacity)
	p := new(objects.Position)
	p.SetPosition(x, y)
	b := button.New(r, M.AddUser, d)
	// Create button reset new Player
	// Create button Play
	// Create button default player
}

func (M *Menu) addButton() (*button, error) {
	var b *button.Button
	var err error
	var x, y int32

	return b, err
}

func (M *Menu) addStructuresPage(r *sdl.Renderer) error {
	var b *block.Block
	var err error
	var x, y int32

	// Create blockheader
	if b, err = block.New(block.Filled); err != nil {
		return err
	}
	b.SetParams(conf.OriginX, conf.MarginTop, conf.WindowWidth, conf.MenuHeaderHeight, conf.ColorBlockRed, conf.ColorBlockGreen, conf.ColorBlockBlue, conf.ColorBlockOpacity)
	if err = b.Init(r); err != nil {
		return nil
	}
	M.layers[layerStructure] = append(M.layers[layerStructure], b)

	// Create blockLeft
	y = conf.MarginTop + conf.MenuHeaderHeight + conf.PaddingBlock
	if b, err = block.New(block.Filled); err != nil {
		return err
	}
	b.SetParams(conf.MarginLeft, y, conf.MenuContentBlockWidth, conf.MenuContentLargeBlockHeight, conf.ColorBlockRed, conf.ColorBlockGreen, conf.ColorBlockBlue, conf.ColorBlockOpacity)
	if err = b.Init(r); err != nil {
		return nil
	}
	M.layers[layerStructure] = append(M.layers[layerStructure], b)

	// Create blockTopRight
	x = conf.WindowWidth - conf.MarginRight - conf.MenuContentBlockWidth
	if b, err = block.New(block.Filled); err != nil {
		return err
	}
	b.SetParams(x, y, conf.MenuContentBlockWidth, conf.MenuContentMediumBlockHeight, conf.ColorBlockRed, conf.ColorBlockGreen, conf.ColorBlockBlue, conf.ColorBlockOpacity)
	if err = b.Init(r); err != nil {
		return nil
	}
	M.layers[layerStructure] = append(M.layers[layerStructure], b)

	// Create blockBottomRight
	y = conf.MarginTop + conf.MenuHeaderHeight + conf.PaddingBlock + conf.MenuContentMediumBlockHeight + conf.PaddingBlock
	if b, err = block.New(block.Filled); err != nil {
		return err
	}
	b.SetParams(x, y, conf.MenuContentBlockWidth, conf.MenuContentMediumBlockHeight, conf.ColorBlockRed, conf.ColorBlockGreen, conf.ColorBlockBlue, conf.ColorBlockOpacity)
	if err = b.Init(r); err != nil {
		return nil
	}
	M.layers[layerStructure] = append(M.layers[layerStructure], b)

	// Create blockFooter
	y = conf.WindowHeight - conf.MarginBot - conf.MenuFooterHeight
	if b, err = block.New(block.Filled); err != nil {
		return err
	}
	b.SetParams(conf.OriginX, y, conf.WindowWidth, conf.MenuHeaderHeight, conf.ColorBlockRed, conf.ColorBlockGreen, conf.ColorBlockBlue, conf.ColorBlockOpacity)
	if err = b.Init(r); err != nil {
		return nil
	}
	M.layers[layerStructure] = append(M.layers[layerStructure], b)

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
