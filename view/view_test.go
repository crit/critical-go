package view

import (
	"strings"
	"testing"
)

func Test_NewView(t *testing.T) {
	view := NewView("test_tpl.html")

	r := view.Render()

	if !strings.Contains(r, "test") {
		t.Errorf("Could not find 'test' in %v", r)
	}
}
