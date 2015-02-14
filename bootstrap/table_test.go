package bootstrap_test

import (
	"github.com/crit/critical-go/bootstrap"
	"testing"
)

func TestTable(t *testing.T) {
	a := `<table class="table">Test</table>`

	b := bootstrap.Table("Test")

	if a != b {
		t.Errorf("expected %s; got %s", a, b)
	}
}

func TestTableStripped(t *testing.T) {
	a := `<table class="table table-stripped">Test</table>`

	b := bootstrap.TableStripped("Test")

	if a != b {
		t.Errorf("expected %s; got %s", a, b)
	}
}

func TestTableBordered(t *testing.T) {
	a := `<table class="table table-bordered">Test</table>`

	b := bootstrap.TableBordered("Test")

	if a != b {
		t.Errorf("expected %s; got %s", a, b)
	}
}

func TestTableCondensed(t *testing.T) {
	a := `<table class="table table-condensed">Test</table>`

	b := bootstrap.TableCondensed("Test")

	if a != b {
		t.Errorf("expected %s; got %s", a, b)
	}
}
