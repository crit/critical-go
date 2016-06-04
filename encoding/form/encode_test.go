package form

import (
	"testing"
)

type Optionals struct {
	Sr string `form:"sr"`
	So string `form:"so,omitempty"`
	Sw string `form:"-"`

	Ir int `form:"omitempty"` // actually named omitempty
	Io int `form:"io,omitempty"`

	Fr float64 `form:"fr"`
	Fo float64 `form:"fo,omitempty"`

	Br bool `form:"br"`
	Bo bool `form:"bo,omitempty"`

	Ur uint `form:"ur"`
	Uo uint `form:"uo,omitempty"`

	// unsupported
	Str struct{}               `form:"str"`
	Mr  map[string]interface{} `form:"mr"`
	Slr []string               `form:"slr"`
}

func optionalsExpected() string {
	out := `<form>`
	out += `<input type="text" name="sr" value="">`
	out += `<input type="text" name="omitempty" value="0">`
	out += `<input type="text" name="fr" value="0.0000000000">`
	out += `<input type="checkbox" name="br">`
	out += `<input type="text" name="ur" value="0">`
	out += `</form>`

	return out
}

func TestOmitEmpty(t *testing.T) {
	var o Optionals
	o.Sw = "something"
	o.Mr = map[string]interface{}{}

	got, err := Marshal(o)

	if err != nil {
		t.Fatal(err)
	}

	if got != optionalsExpected() {
		t.Errorf("\ngot: %s\nwant:%s\n", got, optionalsExpected())
	}
}

func TestFormEncoding(t *testing.T) {
	test := struct {
		ID     uint `form:"id,hidden"`
		Name   string
		Active bool
		Omit   int `form:"-"`
	}{
		1,
		"Test",
		true,
		22,
	}

	expected := `<form>`
	expected += `<input type="hidden" name="id" value="1">`
	expected += `<input type="text" name="name" value="Test">`
	expected += `<input type="checkbox" name="active" checked>`
	expected += `</form>`

	got, err := Marshal(test)

	if err != nil {
		t.Errorf("unexpected error '%s'", err.Error())
		return
	}

	if got != expected {
		t.Errorf("expected '%s' got '%s'", expected, got)
		return
	}
}
