package loader

import (
	"os"
	"runtime"

	"github.com/ymohl-cl/game-builder/objects"
	"github.com/ymohl-cl/game-builder/objects/block"
	"github.com/ymohl-cl/game-builder/objects/text"
)

func (L *DefaultLoader) addBlockLoading() error {
	var b *block.Block
	var err error

	if b, err = block.New(block.Filled); err != nil {
		return err
	}
	b.SetVariantStyle(colorRed, colorGreen, colorBlue, colorOpacity, objects.SFix)
	b.UpdateSize(L.widthBlock, L.heightBlock)
	b.UpdatePosition(originX+L.widthBlock, originY+(L.heightScreen-L.heightScreen/4))

	if err = b.Init(L.renderer); err != nil {
		return err
	}
	L.layers[layerLoadingBar] = append(L.layers[layerLoadingBar], b)
	L.loadBlock = b
	return nil
}

func (L *DefaultLoader) addTxt() error {
	var t *text.Text
	var err error
	var x, y int32
	var fontPath string

	x = L.widthScreen / 2
	y = L.heightScreen / 4

	// get good path to open font
	fontPath = os.Getenv("GOPATH") + "/pkg/" + runtime.GOOS + "_" + runtime.GOARCH + font

	// add title
	if t, err = text.New(mainTxt, sizeMainTxt, fontPath, x, y); err != nil {
		return err
	}
	t.SetVariantStyle(colorRed, colorGreen, colorBlue, colorOpacity, objects.SFix)
	if err = t.Init(L.renderer); err != nil {
		return err
	}
	L.layers[layerText] = append(L.layers[layerText], t)

	// add signature
	y = L.heightScreen / 2
	if t, err = text.New(txtBottomToLoad, sizeTxtBottomToLoad, fontPath, x, y); err != nil {
		return err
	}
	t.SetVariantStyle(colorRed, colorGreen, colorBlue, colorOpacity, objects.SFix)
	if err = t.Init(L.renderer); err != nil {
		return err
	}
	L.layers[layerText] = append(L.layers[layerText], t)

	return nil
}
