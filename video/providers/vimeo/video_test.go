package vimeo_test

import (
	"github.com/crit/critical-go/video/providers/vimeo"
	"testing"
)

func TestNew(t *testing.T) {

	b := "https://vimeo.com/76979871"

	_, err := vimeo.New(b)

	if err != nil {
		t.Errorf("unexpected error: %s", err.Error())
	}
}
