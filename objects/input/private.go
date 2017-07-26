package input

import (
	"strings"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/ymohl-cl/game-builder/objects"
)

/*
** Private method
 */
func (I *Input) setParamsByStatus() {
	switch I.status {
	case objects.SFix:
		I.txt.SetParams(I.fix.txtRect.X, I.fix.txtRect.Y, I.fix.txtColor.R, I.fix.txtColor.G, I.fix.txtColor.B, I.fix.txtColor.A)
		I.block.SetParams(I.fix.blockRect.X, I.fix.blockRect.Y, I.fix.blockRect.W, I.fix.blockRect.H, I.fix.blockColor.R, I.fix.blockColor.G, I.fix.blockColor.B, I.fix.blockColor.A)
	case objects.SBasic:
		I.txt.SetParams(I.basic.txtRect.X, I.basic.txtRect.Y, I.basic.txtColor.R, I.basic.txtColor.G, I.basic.txtColor.B, I.basic.txtColor.A)
		I.block.SetParams(I.basic.blockRect.X, I.basic.blockRect.Y, I.basic.blockRect.W, I.basic.blockRect.H, I.basic.blockColor.R, I.basic.blockColor.G, I.basic.blockColor.B, I.basic.blockColor.A)
	case objects.SOver:
		I.txt.SetParams(I.over.txtRect.X, I.over.txtRect.Y, I.over.txtColor.R, I.over.txtColor.G, I.over.txtColor.B, I.over.txtColor.A)
		I.block.SetParams(I.over.blockRect.X, I.over.blockRect.Y, I.over.blockRect.W, I.over.blockRect.H, I.over.blockColor.R, I.over.blockColor.G, I.over.blockColor.B, I.over.blockColor.A)
	case objects.SClick:
		I.txt.SetParams(I.click.txtRect.X, I.click.txtRect.Y, I.click.txtColor.R, I.click.txtColor.G, I.click.txtColor.B, I.click.txtColor.A)
		I.block.SetParams(I.click.blockRect.X, I.click.blockRect.Y, I.click.blockRect.W, I.click.blockRect.H, I.click.blockColor.R, I.click.blockColor.G, I.click.blockColor.B, I.click.blockColor.A)
	}
}

func (I *Input) updateParamsByStatus() {
	switch I.status {
	case objects.SFix:
		I.txt.UpdateColor(I.fix.txtColor.R, I.fix.txtColor.G, I.fix.txtColor.B, I.fix.txtColor.A)
		I.block.UpdateColor(I.fix.blockColor.R, I.fix.blockColor.G, I.fix.blockColor.B, I.fix.blockColor.A)
		I.txt.UpdatePosition(I.fix.txtRect.X, I.fix.txtRect.W)
		I.block.UpdatePosition(I.fix.blockRect.X, I.fix.blockRect.W)
	case objects.SBasic:
		I.txt.UpdateColor(I.basic.txtColor.R, I.basic.txtColor.G, I.basic.txtColor.B, I.basic.txtColor.A)
		I.block.UpdateColor(I.basic.blockColor.R, I.basic.blockColor.G, I.basic.blockColor.B, I.basic.blockColor.A)
		I.txt.UpdatePosition(I.basic.txtRect.X, I.basic.txtRect.W)
		I.block.UpdatePosition(I.basic.blockRect.X, I.basic.blockRect.W)
	case objects.SOver:
		I.txt.UpdateColor(I.over.txtColor.R, I.over.txtColor.G, I.over.txtColor.B, I.over.txtColor.A)
		I.block.UpdateColor(I.over.blockColor.R, I.over.blockColor.G, I.over.blockColor.B, I.over.blockColor.A)
		I.txt.UpdatePosition(I.over.txtRect.X, I.over.txtRect.W)
		I.block.UpdatePosition(I.over.blockRect.X, I.over.blockRect.W)
	case objects.SClick:
		I.txt.UpdateColor(I.click.txtColor.R, I.click.txtColor.G, I.click.txtColor.B, I.click.txtColor.A)
		I.block.UpdateColor(I.click.blockColor.R, I.click.blockColor.G, I.click.blockColor.B, I.click.blockColor.A)
		I.txt.UpdatePosition(I.click.txtRect.X, I.click.txtRect.W)
		I.block.UpdatePosition(I.click.blockRect.X, I.click.blockRect.W)
	}
}

func (I *Input) updatePositionByStatus() {
	switch I.status {
	case objects.SFix:
		I.txt.UpdatePosition(I.fix.txtRect.X, I.fix.txtRect.W)
		I.block.UpdatePosition(I.fix.blockRect.X, I.fix.blockRect.W)
	case objects.SBasic:
		I.txt.UpdatePosition(I.basic.txtRect.X, I.basic.txtRect.W)
		I.block.UpdatePosition(I.basic.blockRect.X, I.basic.blockRect.W)
	case objects.SOver:
		I.txt.UpdatePosition(I.over.txtRect.X, I.over.txtRect.W)
		I.block.UpdatePosition(I.over.blockRect.X, I.over.blockRect.W)
	case objects.SClick:
		I.txt.UpdatePosition(I.click.txtRect.X, I.click.txtRect.W)
		I.block.UpdatePosition(I.click.blockRect.X, I.click.blockRect.W)
	}
}

func (I Input) checkSizeTxt(newStr, oldStr string) bool {
	var widthBlock, widthText int32

	widthBlock, _ = I.block.GetSize()
	widthText, _ = I.txt.GetSize()

	if (widthText / int32(len(oldStr)) * int32(len(newStr)+paddingSizeTxt)) > widthBlock {
		return false
	}
	return true
}

func (I *Input) addKeyCode(s string, key sdl.Keycode) string {
	idx := strings.Index(s, caret)
	newStr := s[:idx] + string(key) + s[idx:]
	return newStr
}

func (I *Input) caretRight(s string) string {
	idx := strings.Index(s, caret)
	if idx == len(s)-1 {
		return s
	}
	newStr := s[:idx] + string(s[idx+1]) + string(caret) + s[idx+2:]
	return newStr
}

func (I *Input) caretLeft(s string) string {
	idx := strings.Index(s, caret)
	if idx == 0 {
		return s
	}
	newStr := s[:idx-1] + string(caret) + s[idx-1:idx] + s[idx+1:]
	return newStr
}

func (I *Input) caretEnd(s string) string {
	idx := strings.Index(s, caret)

	if idx == len(s)-1 {
		return s
	}
	newStr := s[:idx] + s[idx+1:] + string(caret)
	return newStr
}

func (I *Input) caretBegin(s string) string {
	idx := strings.Index(s, caret)
	if idx == 0 {
		return s
	}
	newStr := string(caret) + s[:idx] + s[idx+1:]
	return newStr
}

func (I *Input) removeKeyBackspace(s string) string {
	idx := strings.Index(s, caret)
	if idx == 0 {
		return s
	}
	newStr := s[:idx-1] + s[idx:]
	return newStr
}

func (I *Input) removeKeyDelete(s string) string {
	idx := strings.Index(s, caret)
	if idx == len(s)-1 {
		return s
	}
	newStr := s[:idx+1] + s[idx+2:]
	return newStr
}
