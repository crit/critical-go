package bootstrap_test

import (
	"github.com/crit/critical-go/bootstrap"
	"testing"
)

func TestH1(t *testing.T) {
	a := "<h1>Test</h1>"

	b := bootstrap.H1("Test")

	if a != b {
		t.Errorf("expected %s; got %s", a, b)
	}
}

func TestH2(t *testing.T) {
	a := "<h2>Test</h2>"

	b := bootstrap.H2("Test")

	if a != b {
		t.Errorf("expected %s; got %s", a, b)
	}
}

func TestH3(t *testing.T) {
	a := "<h3>Test</h3>"

	b := bootstrap.H3("Test")

	if a != b {
		t.Errorf("expected %s; got %s", a, b)
	}
}

func TestH4(t *testing.T) {
	a := "<h4>Test</h4>"

	b := bootstrap.H4("Test")

	if a != b {
		t.Errorf("expected %s; got %s", a, b)
	}
}

func TestH5(t *testing.T) {
	a := "<h5>Test</h5>"

	b := bootstrap.H5("Test")

	if a != b {
		t.Errorf("expected %s; got %s", a, b)
	}
}

func TestH6(t *testing.T) {
	a := "<h6>Test</h6>"

	b := bootstrap.H6("Test")

	if a != b {
		t.Errorf("expected %s; got %s", a, b)
	}
}

func TestSmall(t *testing.T) {
	a := "<small>Test</small>"

	b := bootstrap.Small("Test")

	if a != b {
		t.Errorf("expected %s; got %s", a, b)
	}
}

func TestP(t *testing.T) {
	a := "<p>Test</p>"

	b := bootstrap.P("Test")

	if a != b {
		t.Errorf("expected %s; got %s", a, b)
	}
}

func TestLead(t *testing.T) {
	a := "<p class=\"lead\">Test</p>"

	b := bootstrap.Lead("Test")

	if a != b {
		t.Errorf("expected %s; got %s", a, b)
	}
}

func TestMark(t *testing.T) {
	a := "<mark>Test</mark>"

	b := bootstrap.Mark("Test")

	if a != b {
		t.Errorf("expected %s; got %s", a, b)
	}
}

func TestDeleted(t *testing.T) {
	a := "<del>Test</del>"

	b := bootstrap.Deleted("Test")

	if a != b {
		t.Errorf("expected %s; got %s", a, b)
	}
}

func TestStrikethrough(t *testing.T) {
	a := "<s>Test</s>"

	b := bootstrap.Strikethrough("Test")

	if a != b {
		t.Errorf("expected %s; got %s", a, b)
	}
}

func TestInserted(t *testing.T) {
	a := "<ins>Test</ins>"

	b := bootstrap.Inserted("Test")

	if a != b {
		t.Errorf("expected %s; got %s", a, b)
	}
}

func TestUnderlined(t *testing.T) {
	a := "<u>Test</u>"

	b := bootstrap.Underlined("Test")

	if a != b {
		t.Errorf("expected %s; got %s", a, b)
	}
}

func TestBold(t *testing.T) {
	a := "<b>Test</b>"

	b := bootstrap.Bold("Test")

	if a != b {
		t.Errorf("expected %s; got %s", a, b)
	}
}

func TestItalics(t *testing.T) {
	a := "<em>Test</em>"

	b := bootstrap.Italics("Test")

	if a != b {
		t.Errorf("expected %s; got %s", a, b)
	}
}

func TestRight(t *testing.T) {
	a := "<span class=\"text-right\">Test</span>"

	b := bootstrap.Right("Test")

	if a != b {
		t.Errorf("expected %s; got %s", a, b)
	}
}

func TestLeft(t *testing.T) {
	a := "<span class=\"text-left\">Test</span>"

	b := bootstrap.Left("Test")

	if a != b {
		t.Errorf("expected %s; got %s", a, b)
	}
}

func TestCenter(t *testing.T) {
	a := "<span class=\"text-center\">Test</span>"

	b := bootstrap.Center("Test")

	if a != b {
		t.Errorf("expected %s; got %s", a, b)
	}
}

func TestJustify(t *testing.T) {
	a := "<span class=\"text-justify\">Test</span>"

	b := bootstrap.Justify("Test")

	if a != b {
		t.Errorf("expected %s; got %s", a, b)
	}
}

func TestLowercase(t *testing.T) {
	a := "<span class=\"text-lowercase\">Test</span>"

	b := bootstrap.Lowercase("Test")

	if a != b {
		t.Errorf("expected %s; got %s", a, b)
	}
}

func TestUppercase(t *testing.T) {
	a := "<span class=\"text-uppercase\">Test</span>"

	b := bootstrap.Uppercase("Test")

	if a != b {
		t.Errorf("expected %s; got %s", a, b)
	}
}

func TestCapitalize(t *testing.T) {
	a := "<span class=\"text-capitalize\">Test</span>"

	b := bootstrap.Capitalize("Test")

	if a != b {
		t.Errorf("expected %s; got %s", a, b)
	}
}

func TestAbbreviation(t *testing.T) {
	a := "<abbr title=\"Testing\">T</abbr>"

	b := bootstrap.Abbreviation("Testing", "T")

	if a != b {
		t.Errorf("expected %s; got %s", a, b)
	}
}

func TestInitialism(t *testing.T) {
	a := "<abbr title=\"Testing\" class=\"initialism\">T</abbr>"

	b := bootstrap.Initialism("Testing", "T")

	if a != b {
		t.Errorf("expected %s; got %s", a, b)
	}
}

func TestAddress(t *testing.T) {
	a := "<address>Test</address>"

	b := bootstrap.Address("Test")

	if a != b {
		t.Errorf("expected %s; got %s", a, b)
	}
}

func TestBlockQuote(t *testing.T) {
	a := "<blockquote>Test</blockquote>"

	b := bootstrap.BlockQuote("Test")

	if a != b {
		t.Errorf("expected %s; got %s", a, b)
	}
}
