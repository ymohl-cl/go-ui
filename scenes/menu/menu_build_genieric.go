package menu

import (
	"github.com/ymohl-cl/game-builder/conf"
	"github.com/ymohl-cl/game-builder/database"
	"github.com/ymohl-cl/game-builder/objects"
	"github.com/ymohl-cl/game-builder/objects/block"
	"github.com/ymohl-cl/game-builder/objects/button"
	"github.com/ymohl-cl/game-builder/objects/image"
	"github.com/ymohl-cl/game-builder/objects/input"
	"github.com/ymohl-cl/game-builder/objects/text"
)

func (M *Menu) createBlockToDefaultButton(x, y, w, h int32) (*block.Block, error) {
	var b *block.Block
	var err error

	if b, err = block.New(block.Filled); err != nil {
		return nil, err
	}
	// Set style fix and basic
	if err = b.SetVariantStyle(conf.ColorButtonRed, conf.ColorButtonGreen, conf.ColorButtonBlue, conf.ColorButtonOpacity, objects.SFix, objects.SBasic); err != nil {
		return nil, err
	}
	// Set style over
	if err = b.SetVariantStyle(conf.ColorOverButtonRed, conf.ColorOverButtonGreen, conf.ColorOverButtonBlue, conf.ColorOverButtonOpacity, objects.SOver); err != nil {
		return nil, err
	}
	// set style click
	if err = b.SetVariantStyle(conf.ColorClickButtonRed, conf.ColorClickButtonGreen, conf.ColorClickButtonBlue, conf.ColorClickButtonOpacity, objects.SClick); err != nil {
		return nil, err
	}
	// Set position
	b.UpdatePosition(x, y)
	// Set size
	b.UpdateSize(w, h)
	return b, nil
}

func (M *Menu) createTxtToButton(x, y int32, txt string) (*text.Text, error) {
	var t *text.Text
	var err error

	if t, err = text.New(txt, conf.TxtMedium, conf.Font, x, y); err != nil {
		return nil, err
	}
	t.SetVariantStyle(conf.ColorTxtRed, conf.ColorTxtGreen, conf.ColorTxtBlue, conf.ColorTxtOpacity, objects.SFix)
	t.SetVariantUnderStyle(conf.ColorUnderTxtRed, conf.ColorUnderTxtGreen, conf.ColorUnderTxtBlue, conf.ColorUnderTxtOpacity, objects.SFix)
	t.SetUnderPosition(x-conf.TxtUnderPadding, y-conf.TxtUnderPadding)
	return t, nil
}

func (M *Menu) addButtonPlayer(x, y int32, p *database.Player) (*button.Button, error) {
	var t *text.Text
	var bl *block.Block
	var b *button.Button
	var err error

	if bl, err = M.createBlockToDefaultButton(x, y, conf.MenuElementPlayerWidth, conf.MenuElementPlayerHeight); err != nil {
		return nil, err
	}
	w, h := bl.GetSize()
	if t, err = M.createTxtToButton(x+(w/2), y+(h/2), p.Name); err != nil {
		return nil, err
	}
	b = button.New(bl, nil, t)
	// SetAction
	b.SetAction(M.SelectPlayer, p)

	return b, nil
}

func (M *Menu) addButtonDeletePlayer(x, y int32, p *database.Player) (*button.Button, error) {
	var i *image.Image
	var b *button.Button

	x += conf.MenuElementPlayerWidth + conf.MenuElementPadding

	i = image.New(conf.MenuIconDelete, x, y, conf.MenuIconWidth, conf.MenuElementPlayerHeight)
	i.SetVariantStyle(conf.MenuIconDelete, conf.MenuIconOverDelete, conf.MenuIconOverDelete)
	b = button.New(nil, i, nil)
	b.SetAction(M.DeletePlayer, p)

	return b, nil
}

func (M *Menu) addLoadGame(x, y int32, p *database.Player) (*button.Button, error) {
	var i *image.Image
	var b *button.Button

	x += conf.MenuElementPlayerWidth + (conf.MenuElementPadding * 2) + conf.MenuIconWidth

	i = image.New(conf.MenuIconLoad, x, y, conf.MenuIconWidth, conf.MenuElementPlayerHeight)
	i.SetVariantStyle(conf.MenuIconLoad, conf.MenuIconOverLoad, conf.MenuIconOverLoad)
	b = button.New(nil, i, nil)
	b.SetAction(M.LoadGame, p)

	return b, nil
}

func (M *Menu) addButtonStat(x, y int32, p *database.Player) (*button.Button, error) {
	var i *image.Image
	var b *button.Button

	x += conf.MenuElementPlayerWidth + (conf.MenuElementPadding * 3) + (conf.MenuIconWidth * 2)

	i = image.New(conf.MenuIconTrophy, x, y, conf.MenuIconWidth, conf.MenuElementPlayerHeight)
	i.SetVariantStyle(conf.MenuIconTrophy, conf.MenuIconOverTrophy, conf.MenuIconOverTrophy)
	b = button.New(nil, i, nil)
	b.SetAction(M.DrawStat, p)

	return b, nil
}

func (M *Menu) createBlockToInput(x, y int32) (*block.Block, error) {
	var b *block.Block
	var err error

	if b, err = block.New(block.Filled); err != nil {
		return nil, err
	}
	// Set style fix and basic
	if err = b.SetVariantStyle(conf.ColorInputRed, conf.ColorInputGreen, conf.ColorInputBlue, conf.ColorInputOpacity, objects.SFix, objects.SBasic); err != nil {
		return nil, err
	}
	// Set style over
	if err = b.SetVariantStyle(conf.ColorOverInputRed, conf.ColorOverInputGreen, conf.ColorOverInputBlue, conf.ColorOverInputOpacity, objects.SOver); err != nil {
		return nil, err
	}
	// set style click
	if err = b.SetVariantStyle(conf.ColorClickInputRed, conf.ColorClickInputGreen, conf.ColorClickInputBlue, conf.ColorClickInputOpacity, objects.SClick); err != nil {
		return nil, err
	}
	// Set position
	b.UpdatePosition(x, y)
	// Set size
	b.UpdateSize((conf.ButtonWidth*2)+conf.PaddingBlock, conf.ButtonHeight)
	return b, nil
}

func (M *Menu) createInput() (*input.Input, error) {
	var x, y int32
	var w, h int32
	var b *block.Block
	var i *input.Input
	var err error

	y = conf.MarginTop + conf.MenuHeaderHeight + conf.PaddingBlock + conf.MenuContentMediumBlockHeight/2 - (conf.PaddingBlock/2 + conf.ButtonHeight)
	interval := int32((conf.MenuContentBlockWidth - (conf.ButtonWidth*2 + conf.PaddingBlock)) / 2)
	x = conf.WindowWidth - conf.MarginRight - conf.MenuContentBlockWidth + interval

	if b, err = M.createBlockToInput(x, y); err != nil {
		return nil, err
	}

	if i, err = input.New(conf.TxtMedium, conf.Font, b); err != nil {
		return nil, err
	}
	i.Txt.SetVariantStyle(conf.ColorUnderTxtRed, conf.ColorUnderTxtGreen, conf.ColorUnderTxtBlue, conf.ColorUnderTxtOpacity, objects.SFix)
	w = (conf.ButtonWidth * 2) + conf.PaddingBlock
	h = (conf.ButtonHeight)
	i.Txt.UpdatePosition(x+(w/2), y+(h/2))

	return i, nil
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

	// create block
	if bl, err = M.createBlockToDefaultButton(x, y, conf.ButtonWidth, conf.ButtonHeight); err != nil {
		return nil, err
	}

	// create txt
	x += conf.ButtonWidth / 2
	y += conf.ButtonHeight / 2
	if t, err = M.createTxtToButton(x, y, "DEFAULT PLAYERS"); err != nil {
		return nil, err
	}

	// create button
	b = button.New(bl, nil, t)
	b.SetAction(M.DefaultPlayer)

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

	// create block
	if bl, err = M.createBlockToDefaultButton(x, y, conf.ButtonWidth, conf.ButtonHeight); err != nil {
		return nil, err
	}

	// create txt
	if t, err = M.createTxtToButton(x+conf.ButtonWidth/2, y+conf.ButtonHeight/2, "PLAY !"); err != nil {
		return nil, err
	}

	b = button.New(bl, nil, t)
	b.SetAction(M.Play)

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

	// create block
	if bl, err = M.createBlockToDefaultButton(x, y, conf.ButtonWidth, conf.ButtonHeight); err != nil {
		return nil, err
	}

	// create txt
	if t, err = M.createTxtToButton(x+conf.ButtonWidth/2, y+conf.ButtonHeight/2, "RESET NAME"); err != nil {
		return nil, err
	}

	b = button.New(bl, nil, t)
	b.SetAction(M.ResetName)
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

	// create block
	if bl, err = M.createBlockToDefaultButton(x, y, conf.ButtonWidth, conf.ButtonHeight); err != nil {
		return nil, err
	}

	// create txt
	if t, err = M.createTxtToButton(x+conf.ButtonWidth/2, y+conf.ButtonHeight/2, "NEW PLAYER"); err != nil {
		return nil, err
	}

	b = button.New(bl, nil, t)
	b.SetAction(M.NewPlayer)

	return b, nil
}
