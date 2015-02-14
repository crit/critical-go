package bootstrap_test

import (
	"github.com/crit/critical-go/bootstrap"
	"testing"
)

func TestCode(t *testing.T) {
	a := `<code>Test</code>`

	b := bootstrap.Code("Test")

	if a != b {
		t.Errorf("expected %s; got %s", a, b)
	}
}

func TestKeyboard(t *testing.T) {
	a := `<kbd>Test</kbd>`

	b := bootstrap.Keyboard("Test")

	if a != b {
		t.Errorf("expected %s; got %s", a, b)
	}
}

func TestPre(t *testing.T) {
	a := `<pre>Test</pre>`

	b := bootstrap.Pre("Test")

	if a != b {
		t.Errorf("expected %s; got %s", a, b)
	}
}

func TestVar(t *testing.T) {
	a := `<var>Test</var>`

	b := bootstrap.Var("Test")

	if a != b {
		t.Errorf("expected %s; got %s", a, b)
	}
}
