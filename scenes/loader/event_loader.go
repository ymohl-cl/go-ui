package loader

import (
	"github.com/ymohl-cl/game-builder/conf"
	"github.com/ymohl-cl/game-builder/objects"
	"github.com/ymohl-cl/game-builder/objects/block"
)

/*
** Endpoint action from objects click
 */

func (L *Load) addLoadingBar() {
	var b *block.Block
	var err error
	var loop bool

	loop = true
	for loop {
		select {
		case <-L.closer:
			L.resetLoadingBlock()
			loop = false
		default:
			L.initialized = false
			if b, err = L.lastLoadBlock.Clone(L.renderer); err != nil {
				panic(err)
			}
			x, y := L.lastLoadBlock.GetPosition()
			if x+conf.LoadBlockWidth > conf.WindowWidth {
				L.resetLoadingBlock()
			} else {
				b.UpdatePosition(x+conf.LoadBlockWidth, y)
				L.layers[layerLoadingBar] = append(L.layers[layerLoadingBar], b)
				L.lastLoadBlock = b
			}
			L.initialized = true
		}
	}
}

func (L *Load) resetLoadingBlock() {
	L.lastLoadBlock = L.layers[layerLoadingBar][0].(*block.Block)
	del := L.layers[layerLoadingBar][1:]
	L.layers[layerLoadingBar] = L.layers[layerLoadingBar][:1]
	go clearLoadingBar(del)
}

func clearLoadingBar(sl []objects.Object) {
	var err error

	for _, v := range sl {
		if err = v.Close(); err != nil {
			panic(err)
		}
	}
}
