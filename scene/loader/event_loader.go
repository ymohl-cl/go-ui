package loader

/*
** Endpoint action from scene
 */

func (l *DefaultLoader) addLoadingBar() {
	x, _ := l.loadBlock.GetPosition()
	w, h := l.loadBlock.GetSize()
	// check next width [padding x + blockWidt] + [width current] + [width next block]
	if x+l.widthBlock+w+l.widthBlock > l.widthScreen-l.widthBlock {
		l.resetLoadingBlock()
	} else {
		l.loadBlock.UpdateSize(w+l.widthBlock, h)
	}
}

func (l *DefaultLoader) resetLoadingBlock() {
	_, h := l.loadBlock.GetSize()
	l.loadBlock.UpdateSize(l.widthBlock, h)
}
