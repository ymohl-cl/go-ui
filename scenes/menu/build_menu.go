package menu

import (
	"errors"

	"github.com/ymohl-cl/game-builder/audio"
	"github.com/ymohl-cl/game-builder/conf"
	"github.com/ymohl-cl/game-builder/objects"
	"github.com/ymohl-cl/game-builder/objects/block"
	"github.com/ymohl-cl/game-builder/objects/button"
	"github.com/ymohl-cl/game-builder/objects/image"
	"github.com/ymohl-cl/game-builder/objects/input"
	"github.com/ymohl-cl/game-builder/objects/text"
)

func (M *Menu) addMusic() error {
	var err error

	M.music, err = audio.New(conf.MenuMusic, 1)
	if err != nil {
		return err
	}

	if err = M.music.Init(M.renderer); err != nil {
		return err
	}
	return nil
}

func (M *Menu) addBackground() error {
	var i *image.Image
	var err error

	i = image.New(conf.MenuBackground, conf.OriginX, conf.OriginY, conf.WindowWidth, conf.WindowHeight)
	if err = i.Init(M.renderer); err != nil {
		return err
	}
	M.layers[layerBackground] = append(M.layers[layerBackground], i)
	return nil
}

func (M *Menu) addStructuresPage() error {
	var b *block.Block
	var err error
	var x, y int32

	// Create blockheader
	if b, err = block.New(block.Filled); err != nil {
		return err
	}
	b.SetVariantStyle(conf.ColorBlockRed, conf.ColorBlockGreen, conf.ColorBlockBlue, conf.ColorBlockOpacity, objects.SFix)
	b.UpdatePosition(conf.OriginX, conf.MarginTop)
	b.UpdateSize(conf.WindowWidth, conf.MenuHeaderHeight)
	if err = b.Init(M.renderer); err != nil {
		return nil
	}
	M.layers[layerStructure] = append(M.layers[layerStructure], b)

	// Create blockLeft
	y = conf.MarginTop + conf.MenuHeaderHeight + conf.PaddingBlock
	if b, err = block.New(block.Filled); err != nil {
		return err
	}
	b.SetVariantStyle(conf.ColorBlockRed, conf.ColorBlockGreen, conf.ColorBlockBlue, conf.ColorBlockOpacity, objects.SFix)
	b.UpdatePosition(conf.MarginLeft, y)
	b.UpdateSize(conf.MenuContentBlockWidth, conf.MenuContentLargeBlockHeight)
	if err = b.Init(M.renderer); err != nil {
		return nil
	}
	M.layers[layerStructure] = append(M.layers[layerStructure], b)

	// Create blockTopRight
	x = conf.WindowWidth - conf.MarginRight - conf.MenuContentBlockWidth
	if b, err = block.New(block.Filled); err != nil {
		return err
	}
	b.SetVariantStyle(conf.ColorBlockRed, conf.ColorBlockGreen, conf.ColorBlockBlue, conf.ColorBlockOpacity, objects.SFix)
	b.UpdatePosition(x, y)
	b.UpdateSize(conf.MenuContentBlockWidth, conf.MenuContentMediumBlockHeight)
	if err = b.Init(M.renderer); err != nil {
		return nil
	}
	M.layers[layerStructure] = append(M.layers[layerStructure], b)

	// Create blockBottomRight
	y = conf.MarginTop + conf.MenuHeaderHeight + conf.PaddingBlock + conf.MenuContentMediumBlockHeight + conf.PaddingBlock
	if b, err = block.New(block.Filled); err != nil {
		return err
	}
	b.SetVariantStyle(conf.ColorBlockRed, conf.ColorBlockGreen, conf.ColorBlockBlue, conf.ColorBlockOpacity, objects.SFix)
	b.UpdatePosition(x, y)
	b.UpdateSize(conf.MenuContentBlockWidth, conf.MenuContentMediumBlockHeight)
	if err = b.Init(M.renderer); err != nil {
		return nil
	}
	M.layers[layerStructure] = append(M.layers[layerStructure], b)

	// Create blockFooter
	y = conf.WindowHeight - conf.MarginBot - conf.MenuFooterHeight
	if b, err = block.New(block.Filled); err != nil {
		return err
	}
	b.SetVariantStyle(conf.ColorBlockRed, conf.ColorBlockGreen, conf.ColorBlockBlue, conf.ColorBlockOpacity, objects.SFix)
	b.UpdatePosition(conf.OriginX, y)
	b.UpdateSize(conf.WindowWidth, conf.MenuHeaderHeight)
	if err = b.Init(M.renderer); err != nil {
		return nil
	}
	M.layers[layerStructure] = append(M.layers[layerStructure], b)

	return nil
}

func (M *Menu) addButtons() error {
	var err error
	var b *button.Button

	if b, err = M.getButtonNewPlayer(); err != nil {
		return err
	}
	if err = b.Init(M.renderer); err != nil {
		return err
	}
	M.layers[layerButton] = append(M.layers[layerButton], b)

	if b, err = M.getButtonResetName(); err != nil {
		return err
	}
	if err = b.Init(M.renderer); err != nil {
		return err
	}
	M.layers[layerButton] = append(M.layers[layerButton], b)

	if b, err = M.getButtonPlay(); err != nil {
		return err
	}
	if err = b.Init(M.renderer); err != nil {
		return err
	}
	M.layers[layerButton] = append(M.layers[layerButton], b)

	if b, err = M.getButtonDefaultPlayers(); err != nil {
		return err
	}
	if err = b.Init(M.renderer); err != nil {
		return err
	}
	M.layers[layerButton] = append(M.layers[layerButton], b)

	return nil
}

func (M *Menu) addNotice() error {
	var t *text.Text
	var err error
	var x, y int32

	x = conf.WindowWidth / 2
	y = conf.WindowHeight - conf.MarginBot - (conf.MenuFooterHeight / 2)
	if t, err = text.New("", conf.TxtLittle, conf.Font, x, y); err != nil {
		return err
	}
	t.SetVariantStyle(conf.ColorTxtRed, conf.ColorTxtGreen, conf.ColorTxtBlue, conf.ColorTxtOpacity, objects.SFix)
	t.SetVariantUnderStyle(conf.ColorUnderTxtRed, conf.ColorUnderTxtGreen, conf.ColorUnderTxtBlue, conf.ColorUnderTxtOpacity, objects.SFix)
	t.SetUnderPosition(x-conf.TxtUnderPadding, y-conf.TxtUnderPadding)
	M.notice = t
	M.layers[layerNotice] = append(M.layers[layerNotice], M.notice)
	return nil
}

func (M *Menu) addText() error {
	var t *text.Text
	var err error
	var x, y int32

	x = conf.WindowWidth / 2
	y = conf.MarginTop + (conf.MenuHeaderHeight / 2)
	// add title
	if t, err = text.New("GOMOKU", conf.TxtLarge, conf.Font, x, y); err != nil {
		return err
	}
	t.SetVariantStyle(conf.ColorTxtRed, conf.ColorTxtGreen, conf.ColorTxtBlue, conf.ColorTxtOpacity, objects.SFix)
	t.SetVariantUnderStyle(conf.ColorUnderTxtRed, conf.ColorUnderTxtGreen, conf.ColorUnderTxtBlue, conf.ColorUnderTxtOpacity, objects.SFix)
	t.SetUnderPosition(x-conf.TxtUnderPadding, y-conf.TxtUnderPadding)
	if err = t.Init(M.renderer); err != nil {
		return err
	}
	M.layers[layerText] = append(M.layers[layerText], t)

	// add signature
	y = conf.WindowHeight - (conf.MarginBot / 2)
	signature := "Gomuku is present to you by Anis (agadhgad) and MrPiou (ymohl-cl), Enjoy !"
	if t, err = text.New(signature, conf.TxtLittle, conf.Font, x, y); err != nil {
		return err
	}
	t.SetVariantStyle(conf.ColorTxtRed, conf.ColorTxtGreen, conf.ColorTxtBlue, conf.ColorTxtOpacity, objects.SFix)
	t.SetVariantUnderStyle(conf.ColorUnderTxtRed, conf.ColorUnderTxtGreen, conf.ColorUnderTxtBlue, conf.ColorUnderTxtOpacity, objects.SFix)
	t.SetUnderPosition(x-conf.TxtUnderPadding, y-conf.TxtUnderPadding)
	if err = t.Init(M.renderer); err != nil {
		return err
	}
	M.layers[layerText] = append(M.layers[layerText], t)

	return nil
}

func (M *Menu) addVS() error {
	var t *text.Text
	var err error
	var y, x int32

	// add title
	p1 := M.data.Current.P1
	p2 := M.data.Current.P2
	y = conf.MarginTop + conf.MenuHeaderHeight + conf.PaddingBlock + conf.MenuContentMediumBlockHeight + conf.MenuContentMediumBlockHeight/2 - (conf.PaddingBlock / 2)
	x = conf.WindowWidth - conf.MarginRight - (conf.MenuContentBlockWidth / 2)
	if t, err = text.New(p1.Name+" VS "+p2.Name, conf.TxtMedium, conf.Font, x, y); err != nil {
		return err
	}
	t.SetVariantStyle(conf.ColorTxtRed, conf.ColorTxtGreen, conf.ColorTxtBlue, conf.ColorTxtOpacity, objects.SFix)
	t.SetVariantUnderStyle(conf.ColorUnderTxtRed, conf.ColorUnderTxtGreen, conf.ColorUnderTxtBlue, conf.ColorUnderTxtOpacity, objects.SFix)
	t.SetUnderPosition(x-conf.TxtUnderPadding, y-conf.TxtUnderPadding)
	if err = t.Init(M.renderer); err != nil {
		return err
	}

	M.layers[layerVS] = append(M.layers[layerVS], t)
	M.vs = t
	return nil
}

func (M *Menu) addInput() error {
	var err error
	var i *input.Input

	if i, err = M.createInput(); err != nil {
		return err
	}
	if err = i.Init(M.renderer); err != nil {
		return err
	}
	M.layers[layerInput] = append(M.layers[layerInput], i)

	M.input = i
	return nil
}

func (M *Menu) addPlayers() error {
	var err error
	var x, y int32
	var b1, b2, b3, b4 *button.Button

	if M.data == nil {
		return errors.New(errorData)
	}

	x = conf.MarginLeft
	y = conf.MarginTop + conf.MenuHeaderHeight + conf.PaddingBlock
	for _, p := range M.data.Players {
		if b1, err = M.addButtonPlayer(x, y, p); err != nil {
			return err
		}
		if err = b1.Init(M.renderer); err != nil {
			return err
		}
		M.layers[layerPlayers] = append(M.layers[layerPlayers], b1)

		if b2, err = M.addButtonDeletePlayer(x, y, p); err != nil {
			return err
		}
		if err = b2.Init(M.renderer); err != nil {
			return err
		}
		M.layers[layerPlayers] = append(M.layers[layerPlayers], b2)

		if b3, err = M.addButtonStat(x, y, p); err != nil {
			return err
		}
		if err = b3.Init(M.renderer); err != nil {
			return err
		}
		M.layers[layerPlayers] = append(M.layers[layerPlayers], b3)

		if b4, err = M.addLoadGame(x, y, p); err != nil {
			return err
		}
		if err = b4.Init(M.renderer); err != nil {
			return err
		}
		M.layers[layerPlayers] = append(M.layers[layerPlayers], b4)

		y += conf.MenuElementPlayerHeight + conf.MenuElementPadding
	}

	return nil
}
