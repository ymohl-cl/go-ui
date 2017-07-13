package audio

import (
	"errors"
	"sync"

	"github.com/42MrPiou42/game-builder/objects"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

type Audio struct {
	// infos object
	status      uint8
	initialized bool

	// content object
	url string

	// sdl objects
	music *mix.Music
}

/*
** Functions audio specifications
 */

// New create a new audio object
func New(url string) (*Audio, error) {
	a := new(Audio)

	if url == "" {
		return errors.New("Audio url is empty")
	}

	a.url = url
	a.status = objects.SFix
	return a, nil
}

/*
** Interface objects functions
 */

// IsInit return status initialize
func (A Audio) IsInit() bool {
	return A.initialized
}

// Init audio object
func (A *Audio) Init(r *sdl.Renderer) error {
	var err error

	A.music, err = mix.LoadMUS(A.url)
	if err != nil {
		return err
	}

	A.initialized = true
	return nil
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
func (A *Audio) Draw(r *sdl.Renderer, wg *sync.WaitGroup) error {
	if r == nil {
		return errors.New("Can't draw image because renderer is nil")
	}
	if wg == nil {
		return errors.New("Can't draw image because sync WaitGroup not define")
	}
	if A.initialized == false {
		return errors.New("Can't draw image object is not initialized")
	}

	wg.Add(1)
	defer wg.Done()

	sdl.Do(func() {
		A.music.Play(1)
	})
	return nil
}
