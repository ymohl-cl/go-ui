package image

import (
	"github.com/ymohl-cl/go-ui/objects"
)

/*
** Privates method to Image object
 */

func (I *Image) cloneStatus(source *Image) {
	if source.status == objects.SFix {
		return
	}

	for id, v := range source.urls {
		if id != objects.SFix {
			I.urls[id] = v
		}
	}
}
