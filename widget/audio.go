package widget

import (
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

type Audio struct {
	widget

	channel     int
	chunk       *mix.Chunk
	stateToPlay StateWidget
	repeat      bool
	playing     bool
}

func NewAudio(path string, repeat bool) (*Audio, error) {
	var a Audio
	var err error

	if a.chunk, err = mix.LoadWAV(path); err != nil {
		return nil, err
	}
	a.repeat = repeat
	// set position and state to default behavior with background music unhoverable and unactionable
	a.stateToPlay = StateBase
	a.SetPosition(-1, -1)
	return &a, nil
}

func (a *Audio) Close() {
	a.chunk.Free()
}

func (a *Audio) Render(r *sdl.Renderer) error {
	var err error

	if a.state == a.stateToPlay && !a.playing {
		repeat := 1
		if a.repeat {
			repeat = -1
		}
		if a.channel, err = a.chunk.Play(-1, repeat); err != nil {
			return err
		}
		a.playing = true
	} else if a.state != a.stateToPlay && a.playing {
		mix.HaltChannel(a.channel)
		a.playing = false
	}

	return nil
}

func (a *Audio) SetStateToPlay(s StateWidget) {
	a.stateToPlay = s
}
