package youtube_test

import (
	"github.com/crit/critical-go/video/providers/youtube"
	"testing"
)

func TestNew(t *testing.T) {
	a := "dQw4w9WgXcQ"

	b := "https://www.youtube.com/watch?v=" + a + "&feature=youtu.be"

	_, err := youtube.New(b)

	if err != nil {
		t.Errorf("unexpected error: %s", err.Error())
	}
}
