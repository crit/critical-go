package bootstrap_test

import (
	"github.com/crit/critical-go/bootstrap"
	"testing"
)

func TestRow(t *testing.T) {
	a := `<div class="row">Test</div>`

	b := bootstrap.Row("Test")

	if a != b {
		t.Errorf("expected %s; got %s", a, b)
	}
}

func TestMultiSizes(t *testing.T) {
	a := `<div class="col-md-1 col-lg-1">Test</div>`

	b := bootstrap.Col("Test", "md-1", "lg-1")

	if a != b {
		t.Errorf("expected %s; got %s", a, b)
	}
}
