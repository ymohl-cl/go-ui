package audio

import (
	"errors"
	"sync"

	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/ymohl-cl/go-ui/objects"
)

// Audio structure
type Audio struct {
	// infos object
	status      uint8
	initialized bool
	channel     int

	// content object
	url string

	// sdl objects
	music *mix.Chunk
}

// New create a new object
func New(url string, channel int) (*Audio, error) {
	a := new(Audio)

	if url == "" {
		return nil, errors.New("Audio url is empty")
	}

	a.url = url
	a.status = objects.SFix
	return a, nil
}

// Init audio object
func (A *Audio) Init(r *sdl.Renderer) error {
	var err error

	if r == nil {
		return errors.New(objects.ErrorRenderer)
	}
	A.music, err = mix.LoadWAV(A.url)
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
	A.initialized = false
	if A.music != nil {
		mix.HaltChannel(A.channel)
		A.music.Free()
	}
	return nil
}

// Play the audio source.
func (A *Audio) Play(wg *sync.WaitGroup, r *sdl.Renderer) {
	defer wg.Done()

	sdl.Do(func() {
		if A.initialized == false {
			return
		}
		if r == nil {
			panic(errors.New(objects.ErrorRenderer))
		}
		A.music.Play(A.channel, -1)
	})
}
