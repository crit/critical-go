package video

import (
	"github.com/crit/critical-go/video/providers/vimeo"
	"github.com/crit/critical-go/video/providers/youtube"
)

type VideoProvider interface {
	ID() string
	Player() string
	URI() string
}

func Youtube(src string) (VideoProvider, error) {
	return youtube.New(src)
}

func Vimeo(src string) (VideoProvider, error) {
	return vimeo.New(src)
}
