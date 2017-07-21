package menu

import (
	"errors"

	"github.com/ymohl-cl/game-builder/conf"
	"github.com/ymohl-cl/game-builder/database"
	"github.com/ymohl-cl/game-builder/objects"
	"github.com/ymohl-cl/game-builder/objects/audio"
	"github.com/ymohl-cl/game-builder/objects/block"
	"github.com/ymohl-cl/game-builder/objects/button"
	"github.com/ymohl-cl/game-builder/objects/image"
	"github.com/ymohl-cl/game-builder/objects/input"
	"github.com/ymohl-cl/game-builder/objects/text"
)

func (M *Menu) build() error {
	var err error

	if err = M.addMusic(); err != nil {
		return err
	}
	if err = M.addBackground(); err != nil {
		return err
	}
	if err = M.addStructuresPage(); err != nil {
		return err
	}
	if err = M.addButtons(); err != nil {
		return err
	}
	if err = M.addNotice(); err != nil {
		return err
	}
	if err = M.addText(); err != nil {
		return err
	}
	if err = M.addVS(); err != nil {
		return err
	}
	if err = M.addInput(); err != nil {
		return err
	}
	if err = M.addPlayers(); err != nil {
		return err
	}
	return nil
}

func (M *Menu) addPlayers() error {
	var err error
	var x, y int32
	var b1, b2, b3, b4 *button.Button

	if M.data == nil {
		return errors.New(objects.ErrorData)
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

func (M *Menu) addButtonPlayer(x, y int32, p *database.Player) (*button.Button, error) {
	var t *text.Text
	var bl *block.Block
	var b *button.Button
	var err error

	// Create button
	b = button.New(M.SelectPlayer, p)
	if t, err = text.New(p.Name, conf.TxtMedium, conf.Font); err != nil {
		return nil, err
	}
	t.SetParams(x+conf.MenuElementPlayerWidth/2, y+conf.MenuElementPlayerHeight/2, conf.ColorTxtRed, conf.ColorTxtGreen, conf.ColorTxtBlue, conf.ColorTxtOpacity)
	t.SetUnderParams(conf.ColorUnderTxtRed, conf.ColorUnderTxtGreen, conf.ColorUnderTxtBlue, conf.ColorUnderTxtOpacity, text.PositionBotRight)
	// block sBasic && sFix
	if bl, err = block.New(block.Filled); err != nil {
		return nil, err
	}
	bl.SetParams(x, y, conf.MenuElementPlayerWidth, conf.MenuElementPlayerHeight, conf.ColorButtonRed, conf.ColorButtonGreen, conf.ColorButtonBlue, conf.ColorButtonOpacity)
	b.SetContentBasic(t, nil, bl)
	b.SetContentFix(t, nil, bl)
	// block sOver
	if bl, err = block.New(block.Filled); err != nil {
		return nil, err
	}
	bl.SetParams(x, y, conf.MenuElementPlayerWidth, conf.MenuElementPlayerHeight, conf.ColorOverButtonRed, conf.ColorOverButtonGreen, conf.ColorOverButtonBlue, conf.ColorOverButtonOpacity)
	b.SetContentOver(t, nil, bl)
	// block sClick
	if bl, err = block.New(block.Filled); err != nil {
		return nil, err
	}
	bl.SetParams(x, y, conf.MenuElementPlayerWidth, conf.MenuElementPlayerHeight, conf.ColorClickButtonRed, conf.ColorClickButtonGreen, conf.ColorClickButtonBlue, conf.ColorClickButtonOpacity)
	b.SetContentClick(t, nil, bl)

	return b, nil
}

func (M *Menu) addButtonDeletePlayer(x, y int32, p *database.Player) (*button.Button, error) {
	var i *image.Image
	var b *button.Button
	var err error

	x += conf.MenuElementPlayerWidth + conf.MenuElementPadding

	// Create button
	b = button.New(M.DeletePlayer, p)
	if i, err = image.New(conf.MenuIconDelete); err != nil {
		return nil, err
	}
	i.SetParams(x, y, conf.MenuIconWidth, conf.MenuElementPlayerHeight)
	b.SetContentBasic(nil, i, nil)
	b.SetContentFix(nil, i, nil)
	if i, err = image.New(conf.MenuIconOverDelete); err != nil {
		return nil, err
	}
	i.SetParams(x, y, conf.MenuIconWidth, conf.MenuElementPlayerHeight)
	b.SetContentOver(nil, i, nil)
	if i, err = image.New(conf.MenuIconOverDelete); err != nil {
		return nil, err
	}
	i.SetParams(x+1, y+1, conf.MenuIconWidth-2, conf.MenuElementPlayerHeight-2)
	b.SetContentClick(nil, i, nil)

	return b, nil
}

func (M *Menu) addLoadGame(x, y int32, p *database.Player) (*button.Button, error) {
	var i *image.Image
	var b *button.Button
	var err error

	x += conf.MenuElementPlayerWidth + (conf.MenuElementPadding * 2) + conf.MenuIconWidth

	// Create button
	b = button.New(M.LoadGame, p)
	if i, err = image.New(conf.MenuIconLoad); err != nil {
		return nil, err
	}
	i.SetParams(x, y, conf.MenuIconWidth, conf.MenuElementPlayerHeight)
	b.SetContentBasic(nil, i, nil)
	b.SetContentFix(nil, i, nil)
	if i, err = image.New(conf.MenuIconOverLoad); err != nil {
		return nil, err
	}
	i.SetParams(x, y, conf.MenuIconWidth, conf.MenuElementPlayerHeight)
	b.SetContentOver(nil, i, nil)
	if i, err = image.New(conf.MenuIconOverLoad); err != nil {
		return nil, err
	}
	i.SetParams(x+1, y+1, conf.MenuIconWidth-2, conf.MenuElementPlayerHeight-2)
	b.SetContentClick(nil, i, nil)

	return b, nil
}

func (M *Menu) addButtonStat(x, y int32, p *database.Player) (*button.Button, error) {
	var i *image.Image
	var b *button.Button
	var err error

	x += conf.MenuElementPlayerWidth + (conf.MenuElementPadding * 3) + (conf.MenuIconWidth * 2)

	// Create button
	b = button.New(M.DrawStat, p)
	if i, err = image.New(conf.MenuIconTrophy); err != nil {
		return nil, err
	}
	i.SetParams(x, y, conf.MenuIconWidth, conf.MenuElementPlayerHeight)
	b.SetContentBasic(nil, i, nil)
	b.SetContentFix(nil, i, nil)
	if i, err = image.New(conf.MenuIconOverTrophy); err != nil {
		return nil, err
	}
	i.SetParams(x, y, conf.MenuIconWidth, conf.MenuElementPlayerHeight)
	b.SetContentOver(nil, i, nil)
	if i, err = image.New(conf.MenuIconOverTrophy); err != nil {
		return nil, err
	}
	i.SetParams(x+1, y+1, conf.MenuIconWidth-2, conf.MenuElementPlayerHeight-2)
	b.SetContentClick(nil, i, nil)

	return b, nil
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

func (M *Menu) createInput() (*input.Input, error) {
	var x, y int32
	var bl *block.Block
	var i *input.Input
	var err error

	y = conf.MarginTop + conf.MenuHeaderHeight + conf.PaddingBlock + conf.MenuContentMediumBlockHeight/2 - (conf.PaddingBlock/2 + conf.ButtonHeight)
	interval := int32((conf.MenuContentBlockWidth - (conf.ButtonWidth*2 + conf.PaddingBlock)) / 2)
	x = conf.WindowWidth - conf.MarginRight - conf.MenuContentBlockWidth + interval

	// Create input
	i, err = input.New(objects.SBasic, conf.TxtMedium, conf.Font)
	if err != nil {
		return nil, err
	}
	// block sBasic && sFix
	if bl, err = block.New(block.Filled); err != nil {
		return nil, err
	}
	bl.SetParams(x, y, (conf.ButtonWidth*2)+conf.PaddingBlock, conf.ButtonHeight, conf.ColorInputRed, conf.ColorInputGreen, conf.ColorInputBlue, conf.ColorInputOpacity)
	i.SetBlockBasic(bl)
	i.SetBlockFix(bl)
	// block sOver
	if bl, err = block.New(block.Filled); err != nil {
		return nil, err
	}
	bl.SetParams(x, y, (conf.ButtonWidth*2)+conf.PaddingBlock, conf.ButtonHeight, conf.ColorOverInputRed, conf.ColorOverInputGreen, conf.ColorOverInputBlue, conf.ColorOverInputOpacity)
	i.SetBlockOver(bl)
	// block sClick
	if bl, err = block.New(block.Filled); err != nil {
		return nil, err
	}
	bl.SetParams(x, y, (conf.ButtonWidth*2)+conf.PaddingBlock, conf.ButtonHeight, conf.ColorClickInputRed, conf.ColorClickInputGreen, conf.ColorClickInputBlue, conf.ColorClickInputOpacity)
	i.SetBlockClick(bl)

	err = i.SetColorTxt(conf.ColorUnderTxtRed, conf.ColorUnderTxtGreen, conf.ColorUnderTxtBlue, conf.ColorUnderTxtOpacity)
	if err != nil {
		return nil, err
	}

	return i, nil
}

func (M *Menu) addVS() error {
	var t *text.Text
	var err error
	var y, x int32

	// add title
	p1 := M.data.Current.P1
	p2 := M.data.Current.P2
	if t, err = text.New(p1.Name+" VS "+p2.Name, conf.TxtMedium, conf.Font); err != nil {
		return err
	}
	y = conf.MarginTop + conf.MenuHeaderHeight + conf.PaddingBlock + conf.MenuContentMediumBlockHeight + conf.MenuContentMediumBlockHeight/2 - (conf.PaddingBlock / 2)
	x = conf.WindowWidth - conf.MarginRight - (conf.MenuContentBlockWidth / 2)
	t.SetParams(x, y, conf.ColorTxtRed, conf.ColorTxtGreen, conf.ColorTxtBlue, conf.ColorTxtOpacity)
	t.SetUnderParams(conf.ColorUnderTxtRed, conf.ColorUnderTxtGreen, conf.ColorUnderTxtBlue, conf.ColorUnderTxtOpacity, text.PositionBotRight)
	if err = t.Init(M.renderer); err != nil {
		return err
	}
	M.layers[layerVS] = append(M.layers[layerVS], t)
	return nil
}

func (M *Menu) addText() error {
	var t *text.Text
	var err error

	// add title
	if t, err = text.New("GOMOKU", conf.TxtLarge, conf.Font); err != nil {
		return err
	}
	t.SetParams(conf.WindowWidth/2, conf.MarginTop+(conf.MenuHeaderHeight/2), conf.ColorTxtRed, conf.ColorTxtGreen, conf.ColorTxtBlue, conf.ColorTxtOpacity)
	t.SetUnderParams(conf.ColorUnderTxtRed, conf.ColorUnderTxtGreen, conf.ColorUnderTxtBlue, conf.ColorUnderTxtOpacity, text.PositionBotRight)
	if err = t.Init(M.renderer); err != nil {
		return err
	}
	M.layers[layerText] = append(M.layers[layerText], t)

	signature := "Gomuku is present to you by Anis (agadhgad) and MrPiou (ymohl-cl), Enjoy !"
	if t, err = text.New(signature, conf.TxtLittle, conf.Font); err != nil {
		return err
	}
	t.SetParams(conf.WindowWidth/2, conf.WindowHeight-(conf.MarginBot/2), conf.ColorTxtRed, conf.ColorTxtGreen, conf.ColorTxtBlue, conf.ColorTxtOpacity)
	t.SetUnderParams(conf.ColorUnderTxtRed, conf.ColorUnderTxtGreen, conf.ColorUnderTxtBlue, conf.ColorUnderTxtOpacity, text.PositionBotRight)
	if err = t.Init(M.renderer); err != nil {
		return err
	}
	M.layers[layerText] = append(M.layers[layerText], t)

	return nil
}

func (M *Menu) addNotice() error {
	var t *text.Text
	var err error

	if t, err = text.New("", conf.TxtLittle, conf.Font); err != nil {
		return err
	}
	t.SetParams(conf.WindowWidth/2, conf.WindowHeight-conf.MarginBot-(conf.MenuFooterHeight/2), conf.ColorTxtRed, conf.ColorTxtGreen, conf.ColorTxtBlue, conf.ColorTxtOpacity)
	t.SetUnderParams(conf.ColorUnderTxtRed, conf.ColorUnderTxtGreen, conf.ColorUnderTxtBlue, conf.ColorUnderTxtOpacity, text.PositionBotRight)
	M.notice = t
	M.layers[layerNotice] = append(M.layers[layerNotice], M.notice)
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

func (M *Menu) getButtonDefaultPlayers() (*button.Button, error) {
	var x, y int32
	var t *text.Text
	var bl *block.Block
	var b *button.Button
	var err error

	y = conf.MarginTop + conf.MenuHeaderHeight + conf.PaddingBlock + conf.MenuContentMediumBlockHeight + conf.MenuContentMediumBlockHeight/2 + conf.PaddingBlock + conf.PaddingBlock/2
	interval := int32((conf.MenuContentBlockWidth - (conf.ButtonWidth*2 + conf.PaddingBlock)) / 2)
	x = conf.WindowWidth - conf.MarginRight - interval - conf.ButtonWidth

	// Create button
	b = button.New(M.DefaultPlayer)
	if t, err = text.New("DEFAULT PLAYERS", conf.TxtMedium, conf.Font); err != nil {
		return nil, err
	}
	t.SetParams(x+conf.ButtonWidth/2, y+conf.ButtonHeight/2, conf.ColorTxtRed, conf.ColorTxtGreen, conf.ColorTxtBlue, conf.ColorTxtOpacity)
	t.SetUnderParams(conf.ColorUnderTxtRed, conf.ColorUnderTxtGreen, conf.ColorUnderTxtBlue, conf.ColorUnderTxtOpacity, text.PositionBotRight)
	// block sBasic && sFix
	if bl, err = block.New(block.Filled); err != nil {
		return nil, err
	}
	bl.SetParams(x, y, conf.ButtonWidth, conf.ButtonHeight, conf.ColorButtonRed, conf.ColorButtonGreen, conf.ColorButtonBlue, conf.ColorButtonOpacity)
	b.SetContentBasic(t, nil, bl)
	b.SetContentFix(t, nil, bl)
	// block sOver
	if bl, err = block.New(block.Filled); err != nil {
		return nil, err
	}
	bl.SetParams(x, y, conf.ButtonWidth, conf.ButtonHeight, conf.ColorOverButtonRed, conf.ColorOverButtonGreen, conf.ColorOverButtonBlue, conf.ColorOverButtonOpacity)
	b.SetContentOver(t, nil, bl)
	// block sClick
	if bl, err = block.New(block.Filled); err != nil {
		return nil, err
	}
	bl.SetParams(x, y, conf.ButtonWidth, conf.ButtonHeight, conf.ColorClickButtonRed, conf.ColorClickButtonGreen, conf.ColorClickButtonBlue, conf.ColorClickButtonOpacity)
	b.SetContentClick(t, nil, bl)

	return b, nil
}

func (M *Menu) getButtonPlay() (*button.Button, error) {
	var x, y int32
	var t *text.Text
	var bl *block.Block
	var b *button.Button
	var err error

	y = conf.MarginTop + conf.MenuHeaderHeight + conf.PaddingBlock + conf.MenuContentMediumBlockHeight + conf.MenuContentMediumBlockHeight/2 + conf.PaddingBlock + conf.PaddingBlock/2
	interval := int32((conf.MenuContentBlockWidth - (conf.ButtonWidth*2 + conf.PaddingBlock)) / 2)
	x = conf.WindowWidth - conf.MarginRight - conf.MenuContentBlockWidth + interval

	// Create button
	b = button.New(M.Play)
	if t, err = text.New("PLAY !", conf.TxtMedium, conf.Font); err != nil {
		return nil, err
	}
	t.SetParams(x+conf.ButtonWidth/2, y+conf.ButtonHeight/2, conf.ColorTxtRed, conf.ColorTxtGreen, conf.ColorTxtBlue, conf.ColorTxtOpacity)
	t.SetUnderParams(conf.ColorUnderTxtRed, conf.ColorUnderTxtGreen, conf.ColorUnderTxtBlue, conf.ColorUnderTxtOpacity, text.PositionBotRight)
	// block sBasic && sFix
	if bl, err = block.New(block.Filled); err != nil {
		return nil, err
	}
	bl.SetParams(x, y, conf.ButtonWidth, conf.ButtonHeight, conf.ColorButtonRed, conf.ColorButtonGreen, conf.ColorButtonBlue, conf.ColorButtonOpacity)
	b.SetContentBasic(t, nil, bl)
	b.SetContentFix(t, nil, bl)
	// block sOver
	if bl, err = block.New(block.Filled); err != nil {
		return nil, err
	}
	bl.SetParams(x, y, conf.ButtonWidth, conf.ButtonHeight, conf.ColorOverButtonRed, conf.ColorOverButtonGreen, conf.ColorOverButtonBlue, conf.ColorOverButtonOpacity)
	b.SetContentOver(t, nil, bl)
	// block sClick
	if bl, err = block.New(block.Filled); err != nil {
		return nil, err
	}
	bl.SetParams(x, y, conf.ButtonWidth, conf.ButtonHeight, conf.ColorClickButtonRed, conf.ColorClickButtonGreen, conf.ColorClickButtonBlue, conf.ColorClickButtonOpacity)
	b.SetContentClick(t, nil, bl)

	return b, nil
}

func (M *Menu) getButtonResetName() (*button.Button, error) {
	var x, y int32
	var t *text.Text
	var bl *block.Block
	var b *button.Button
	var err error

	y = conf.MarginTop + conf.MenuHeaderHeight + conf.PaddingBlock + conf.MenuContentMediumBlockHeight/2 + conf.PaddingBlock/2
	interval := int32((conf.MenuContentBlockWidth - (conf.ButtonWidth*2 + conf.PaddingBlock)) / 2)
	x = conf.WindowWidth - conf.MarginRight - interval - conf.ButtonWidth

	// Create button
	b = button.New(M.ResetName)
	if t, err = text.New("RESET NAME", conf.TxtMedium, conf.Font); err != nil {
		return nil, err
	}
	t.SetParams(x+conf.ButtonWidth/2, y+conf.ButtonHeight/2, conf.ColorTxtRed, conf.ColorTxtGreen, conf.ColorTxtBlue, conf.ColorTxtOpacity)
	t.SetUnderParams(conf.ColorUnderTxtRed, conf.ColorUnderTxtGreen, conf.ColorUnderTxtBlue, conf.ColorUnderTxtOpacity, text.PositionBotRight)
	// block sBasic && sFix
	if bl, err = block.New(block.Filled); err != nil {
		return nil, err
	}
	bl.SetParams(x, y, conf.ButtonWidth, conf.ButtonHeight, conf.ColorButtonRed, conf.ColorButtonGreen, conf.ColorButtonBlue, conf.ColorButtonOpacity)
	b.SetContentBasic(t, nil, bl)
	b.SetContentFix(t, nil, bl)
	// block sOver
	if bl, err = block.New(block.Filled); err != nil {
		return nil, err
	}
	bl.SetParams(x, y, conf.ButtonWidth, conf.ButtonHeight, conf.ColorOverButtonRed, conf.ColorOverButtonGreen, conf.ColorOverButtonBlue, conf.ColorOverButtonOpacity)
	b.SetContentOver(t, nil, bl)
	// block sClick
	if bl, err = block.New(block.Filled); err != nil {
		return nil, err
	}
	bl.SetParams(x, y, conf.ButtonWidth, conf.ButtonHeight, conf.ColorClickButtonRed, conf.ColorClickButtonGreen, conf.ColorClickButtonBlue, conf.ColorClickButtonOpacity)
	b.SetContentClick(t, nil, bl)

	return b, nil
}

func (M *Menu) getButtonNewPlayer() (*button.Button, error) {
	var x, y int32
	var t *text.Text
	var bl *block.Block
	var b *button.Button
	var err error

	y = conf.MarginTop + conf.MenuHeaderHeight + conf.PaddingBlock + conf.MenuContentMediumBlockHeight/2 + conf.PaddingBlock/2
	interval := int32((conf.MenuContentBlockWidth - (conf.ButtonWidth*2 + conf.PaddingBlock)) / 2)
	x = conf.WindowWidth - conf.MarginRight - conf.MenuContentBlockWidth + interval

	// Create button
	b = button.New(M.NewPlayer)
	if t, err = text.New("NEW PLAYER", conf.TxtMedium, conf.Font); err != nil {
		return nil, err
	}
	t.SetParams(x+conf.ButtonWidth/2, y+conf.ButtonHeight/2, conf.ColorTxtRed, conf.ColorTxtGreen, conf.ColorTxtBlue, conf.ColorTxtOpacity)
	t.SetUnderParams(conf.ColorUnderTxtRed, conf.ColorUnderTxtGreen, conf.ColorUnderTxtBlue, conf.ColorUnderTxtOpacity, text.PositionBotRight)
	// block sBasic && sFix
	if bl, err = block.New(block.Filled); err != nil {
		return nil, err
	}
	bl.SetParams(x, y, conf.ButtonWidth, conf.ButtonHeight, conf.ColorButtonRed, conf.ColorButtonGreen, conf.ColorButtonBlue, conf.ColorButtonOpacity)
	b.SetContentBasic(t, nil, bl)
	b.SetContentFix(t, nil, bl)
	// block sOver
	if bl, err = block.New(block.Filled); err != nil {
		return nil, err
	}
	bl.SetParams(x, y, conf.ButtonWidth, conf.ButtonHeight, conf.ColorOverButtonRed, conf.ColorOverButtonGreen, conf.ColorOverButtonBlue, conf.ColorOverButtonOpacity)
	b.SetContentOver(t, nil, bl)
	// block sClick
	if bl, err = block.New(block.Filled); err != nil {
		return nil, err
	}
	bl.SetParams(x, y, conf.ButtonWidth, conf.ButtonHeight, conf.ColorClickButtonRed, conf.ColorClickButtonGreen, conf.ColorClickButtonBlue, conf.ColorClickButtonOpacity)
	b.SetContentClick(t, nil, bl)

	return b, nil
}

func (M *Menu) addStructuresPage() error {
	var b *block.Block
	var err error
	var x, y int32

	// Create blockheader
	if b, err = block.New(block.Filled); err != nil {
		return err
	}
	b.SetParams(conf.OriginX, conf.MarginTop, conf.WindowWidth, conf.MenuHeaderHeight, conf.ColorBlockRed, conf.ColorBlockGreen, conf.ColorBlockBlue, conf.ColorBlockOpacity)
	if err = b.Init(M.renderer); err != nil {
		return nil
	}
	M.layers[layerStructure] = append(M.layers[layerStructure], b)

	// Create blockLeft
	y = conf.MarginTop + conf.MenuHeaderHeight + conf.PaddingBlock
	if b, err = block.New(block.Filled); err != nil {
		return err
	}
	b.SetParams(conf.MarginLeft, y, conf.MenuContentBlockWidth, conf.MenuContentLargeBlockHeight, conf.ColorBlockRed, conf.ColorBlockGreen, conf.ColorBlockBlue, conf.ColorBlockOpacity)
	if err = b.Init(M.renderer); err != nil {
		return nil
	}
	M.layers[layerStructure] = append(M.layers[layerStructure], b)

	// Create blockTopRight
	x = conf.WindowWidth - conf.MarginRight - conf.MenuContentBlockWidth
	if b, err = block.New(block.Filled); err != nil {
		return err
	}
	b.SetParams(x, y, conf.MenuContentBlockWidth, conf.MenuContentMediumBlockHeight, conf.ColorBlockRed, conf.ColorBlockGreen, conf.ColorBlockBlue, conf.ColorBlockOpacity)
	if err = b.Init(M.renderer); err != nil {
		return nil
	}
	M.layers[layerStructure] = append(M.layers[layerStructure], b)

	// Create blockBottomRight
	y = conf.MarginTop + conf.MenuHeaderHeight + conf.PaddingBlock + conf.MenuContentMediumBlockHeight + conf.PaddingBlock
	if b, err = block.New(block.Filled); err != nil {
		return err
	}
	b.SetParams(x, y, conf.MenuContentBlockWidth, conf.MenuContentMediumBlockHeight, conf.ColorBlockRed, conf.ColorBlockGreen, conf.ColorBlockBlue, conf.ColorBlockOpacity)
	if err = b.Init(M.renderer); err != nil {
		return nil
	}
	M.layers[layerStructure] = append(M.layers[layerStructure], b)

	// Create blockFooter
	y = conf.WindowHeight - conf.MarginBot - conf.MenuFooterHeight
	if b, err = block.New(block.Filled); err != nil {
		return err
	}
	b.SetParams(conf.OriginX, y, conf.WindowWidth, conf.MenuHeaderHeight, conf.ColorBlockRed, conf.ColorBlockGreen, conf.ColorBlockBlue, conf.ColorBlockOpacity)
	if err = b.Init(M.renderer); err != nil {
		return nil
	}
	M.layers[layerStructure] = append(M.layers[layerStructure], b)

	return nil
}

func (M *Menu) addBackground() error {
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

	if err = i.Init(M.renderer); err != nil {
		return err
	}
	M.layers[layerBackground] = append(M.layers[layerBackground], i)
	return nil
}

func (M *Menu) addMusic() error {
	var err error

	M.music, err = audio.New(conf.MenuMusic)
	if err != nil {
		return err
	}

	if err = M.music.Init(M.renderer); err != nil {
		return err
	}
	return nil
}
