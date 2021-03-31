package widget

import (
	"fmt"
	"os/exec"
	"runtime"
)

type Action interface {
	Run() error
}

type linkAction struct {
	url string
}

type funcAction struct {
	f func()
}

func NewLinkAction(link string) Action {
	return &linkAction{url: link}
}

func NewFuncAction(f func()) Action {
	return &funcAction{f: f}
}

func (a *linkAction) Run() error {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", a.url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", a.url).Start()
	case "darwin":
		err = exec.Command("open", a.url).Start()
	default:
		err = fmt.Errorf("unsupported platform to open link")
	}
	if err != nil {
		return err
	}
	return nil
}

func (a *funcAction) Run() error {
	a.f()
	return nil
}
