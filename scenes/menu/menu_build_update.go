package menu

import (
	"errors"

	"github.com/ymohl-cl/game-builder/conf"
	"github.com/ymohl-cl/game-builder/database"
	"github.com/ymohl-cl/game-builder/objects/button"
)

func (M *Menu) addUIPlayer(nb int, p *database.Player) error {
	var err error
	var x, y int32
	var b1, b2, b3, b4 *button.Button

	x = conf.MarginLeft
	y = conf.MarginTop + conf.MenuHeaderHeight + conf.PaddingBlock
	y += (conf.MenuElementPlayerHeight + conf.MenuElementPadding) * int32(nb)

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

	return nil
}

func (M Menu) closeUIPlayer(idx int) error {
	var err error
	var size int

	size = len(M.layers[layerPlayers])
	if size <= idx {
		return errors.New("id object not found")
	}
	button := M.layers[layerPlayers][idx]
	if err = button.Close(); err != nil {
		return err
	}
	return nil
}

func (M *Menu) removeUIPlayer(idData int) error {
	var err error
	var id int

	id = idData * buttonByPlayer
	if err = M.closeUIPlayer(id); err != nil {
		return err
	}
	if err = M.closeUIPlayer(id + 1); err != nil {
		return err
	}
	if err = M.closeUIPlayer(id + 2); err != nil {
		return err
	}
	if err = M.closeUIPlayer(id + 3); err != nil {
		return err
	}

	// Update position next elements
	for _, b := range M.layers[layerPlayers][id+4:] {
		x, y := b.GetPosition()
		y -= (conf.MenuElementPlayerHeight + conf.MenuElementPadding)
		b.UpdatePosition(x, y)
	}

	M.layers[layerPlayers] = append(M.layers[layerPlayers][:id], M.layers[layerPlayers][id+4:]...)
	return nil
}

func (M *Menu) updateVS() {
	var err error
	p1 := M.data.Current.P1
	p2 := M.data.Current.P2

	if p1 == nil || p2 == nil {
		panic(errors.New("Players is nil"))
	}
	if M.vs.IsInit() {
		if err = M.vs.Close(); err != nil {
			panic(err)
		}
	}
	if err = M.vs.UpdateText(p1.Name+" VS "+p2.Name, M.renderer); err != nil {
		panic(err)
	}
}
