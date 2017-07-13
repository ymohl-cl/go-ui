package input

import (
	"strings"

	"github.com/veandco/go-sdl2/sdl"
)

/*
** Private method
 */

func (I Input) checkSizeTxt(newStr, oldStr string) bool {
	var widthBlock, widthText int32

	widthBlock, _ = I.block.GetSize()
	widthText, _ = I.Txt.GetSize()

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
