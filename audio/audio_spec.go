package audio

import (
	"errors"

	"github.com/ymohl-cl/game-builder/objects"
)

// New create a new object
func New(url string) (*Audio, error) {
	a := new(Audio)

	if url == "" {
		return nil, errors.New("Audio url is empty")
	}

	a.url = url
	a.status = objects.SFix
	return a, nil
}
