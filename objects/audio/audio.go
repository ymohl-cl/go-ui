package audio

import (
	"errors"
	"sync"

	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/ymohl-cl/game-builder/objects"
)

type Audio struct {
	// infos object
	status      uint8
	initialized bool

	// content object
	url string

	// sdl objects
	music    *mix.Music
	renderer *sdl.Renderer
}

// Init audio object
func (A *Audio) Init(r *sdl.Renderer) error {
	var err error

	if r == nil {
		return errors.New(objects.ErrorRenderer)
	}
	A.renderer = r

	A.music, err = mix.LoadMUS(A.url)
	if err != nil {
		return err
	}

	A.initialized = true
	return nil
}

// IsInit return status initialize
func (A Audio) IsInit() bool {
	return A.initialized
}

// Close sdl objects
func (A *Audio) Close() error {
	if A.music != nil {
		A.music.Free()
	}
	A.initialized = false
	return nil
}

func (A *Audio) GetStatus() uint8 {
	return A.status
}

func (A *Audio) IsOver(x, y int32) bool {
	return false
}

func (A *Audio) Click() {
	return
}

func (A *Audio) SetStatus(s uint8) {
	if A.status != objects.SFix {
		A.status = s
	}
}

// Draw the object audio.
func (A *Audio) Draw(wg *sync.WaitGroup) {
	defer wg.Done()

	sdl.Do(func() {
		if A.initialized == false {
			panic(errors.New(objects.ErrorNotInit))
		}
		A.music.Play(1)
	})
}
