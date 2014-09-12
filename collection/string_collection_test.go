package collection

import (
	"strings"
	"testing"
)

func Test_String_Put(t *testing.T) {
	a := NewStringCollection()

	a.Put("key", "value")

	v := a.items["key"]

	if v != "value" {
		t.Error("Unable to get value for placed key")
	}
}

func Test_String_Get(t *testing.T) {
	a := NewStringCollection()

	a.Put("key", "value")

	v := a.Get("key", "def")

	if v != "value" {
		t.Error("Unable to get preset key")
	}
}

func Test_String_Def(t *testing.T) {
	a := NewStringCollection()

	v := a.Get("key", "def")

	if v != "def" {
		t.Error("Unable to return default value for unset key")
	}
}

func Test_String_Make(t *testing.T) {
	items := make(map[string]string)

	items["key"] = "value"

	a := MakeStringCollection(items)

	v := a.Get("key", "def")

	if v != "value" {
		t.Error("Unable to pass in premade items")
	}
}

func Test_String_Contains(t *testing.T) {
	a := NewStringCollection()

	a.Put("key", "value")

	b := a.Contains("value")

	if b != true {
		t.Error("Unable to find value in items")
	}

	c := a.Contains("nope")

	if c != false {
		t.Error("Unable to fail searching for an unset value")
	}
}

func Test_String_Forget(t *testing.T) {
	a := NewStringCollection()

	a.Put("key", "value")
	a.Forget("key")

	v := a.Get("key", "def")

	if v != "def" {
		t.Error("Unable to forget a key in items")
	}
}

func Test_String_Has(t *testing.T) {
	a := NewStringCollection()

	a.Put("key", "value")

	if !a.Has("key") {
		t.Error("Unable to find preset key")
	}

	if a.Has("nope") {
		t.Error("Unable to fail searching for an unset key")
	}
}

func Test_String_Diff(t *testing.T) {
	a := NewStringCollection()
	b := NewStringCollection()

	a.Put("a", "alpha")
	b.Put("a", "alpha")
	b.Put("b", "beta")

	diff := a.Diff(b)

	if !diff.Contains("beta") {
		t.Error("Unable to find the differnce in items")
	}
}

func Test_String_Each(t *testing.T) {
	a := NewStringCollection()

	a.Put("a", "alpha")
	a.Put("b", "beta")
	a.Put("c", "gamma")
	a.Put("d", "delta")

	a.Each(func(k, v string) (string, string) {
		return k, strings.ToUpper(v)
	})

	if !a.Contains("ALPHA") {
		t.Error("Unable to uppercase alpha")
	}

	if !a.Contains("BETA") {
		t.Error("Unable to uppercase beta")
	}

	if !a.Contains("GAMMA") {
		t.Error("Unable to uppercase gamma")
	}
}

func Test_String_Filter(t *testing.T) {
	a := NewStringCollection()

	a.Put("a", "foo")
	a.Put("b", "bar")
	a.Put("c", "foo")
	a.Put("d", "foo")
	a.Put("e", "bar")

	b := a.Filter(func(v string) bool {
		return v == "bar"
	})

	if !b.Has("b") {
		t.Error("Unable to filter to only specific values")
	}
}

func Test_String_Empty(t *testing.T) {
	a := NewStringCollection()

	if !a.Empty() {
		t.Error("Empty should return true")
	}
}

func Test_String_Flatten(t *testing.T) {
	a := NewStringCollection()

	a.Put("a", "alpha")
	a.Put("b", "beta")

	c := a.Flatten()

	v := c[0]

	if !a.Contains(v) {
		t.Errorf("Expected %v to be in %v", v, a)
	}

}

func Test_String_Flip(t *testing.T) {
	a := NewStringCollection()

	a.Put("a", "alpha")
	a.Put("b", "beta")

	b := a.Flip()

	if !b.Has("alpha") {
		t.Error("Unable to flip items")
	}
}

func Test_String_Implode(t *testing.T) {
	b := NewStringCollection()

	b.Put("a", "alpha")
	b.Put("b", "beta")

	c := b.Implode(",")

	if !strings.Contains(c, "alpha") {
		t.Errorf("Expecting 'alpha' in %v", c)
	}
}

func Test_String_Intersect(t *testing.T) {
	a := NewStringCollection()
	b := NewStringCollection()

	a.Put("a", "alpha")
	a.Put("b", "beta")
	a.Put("d", "delta")

	b.Put("a", "alpha")
	b.Put("b", "nope")
	b.Put("c", "gamma")
	b.Put("d", "delta")

	c := a.Intersect(b)

	if c.Has("c") {
		t.Error("Unable to calculate intersection of items")
	}

	if c.Contains("nope") {
		t.Error("Unable to respect local item value over comparison items")
	}
}

func Test_String_Keys(t *testing.T) {
	a := NewStringCollection()

	a.Put("a", "alpha")
	a.Put("b", "beta")

	b := a.Keys()

	v := b[0]

	if !a.Has(v) {
		t.Errorf("Expected %v to be in %v", v, a)
	}
}

func Test_String_Only(t *testing.T) {
	a := NewStringCollection()
	b := []string{"a", "d"}

	a.Put("a", "alpha")
	a.Put("b", "beta")
	a.Put("c", "gamma")
	a.Put("d", "delta")
	a.Put("e", "epsilon")

	d := a.Only(b)

	if !d.Has("a") {
		t.Error("Unable to list entries")
	}

}

func Test_String_Merge(t *testing.T) {
	a := NewStringCollection()
	b := NewStringCollection()

	a.Put("a", "alpha")
	a.Put("b", "nope")

	b.Put("b", "beta")
	b.Put("c", "gamma")
	b.Put("d", "delta")

	a.Merge(b)

	if !a.Has("d") {
		t.Errorf("Expecting 'd' in %v", a)
	}
}

func Test_String_Pull(t *testing.T) {
	a := NewStringCollection()

	a.Put("a", "alpha")
	a.Put("b", "beta")
	a.Put("c", "gamma")

	v, _ := a.Pull("b", "default")

	if v != "beta" {
		t.Errorf("Expected 'beta' and got %v", v)
	}

	if a.Contains("beta") {
		t.Errorf("Expected to not find 'beta' in %v", a)
	}
}

func Test_String_Reduce(t *testing.T) {
	a := NewStringCollection()

	a.Put("a", "alpha")
	a.Put("b", "beta")
	a.Put("c", "gamma")

	v := a.Reduce(func(carry, item string) string {
		return carry + " > " + item
	}, "Start: ")

	if !strings.Contains(v, "> alpha") {
		t.Errorf("Expecting to find '> alpha' in %v", v)
	}
}

func Test_String_Reject(t *testing.T) {
	a := NewStringCollection()

	a.Put("a", "foo")
	a.Put("b", "food")
	a.Put("c", "flood")

	b := a.Reject(func(v string) bool {
		return strings.Contains(v, "foo")
	})

	if !b.Has("c") {
		t.Errorf("Expected 'c' in %v", b)
	}
}

func Test_String_Search(t *testing.T) {
	a := NewStringCollection()

	a.Put("a", "alpha")
	a.Put("b", "beta")
	a.Put("c", "gamma")

	v, f := a.Search("gamma")

	if !f {
		t.Errorf("Did not find 'gamma' in %v", a)
	}

	if v != "c" {
		t.Errorf("Expected 'gamma' to find 'c' in %v", a)
	}
}

func Test_String_Random(t *testing.T) {
	a := NewStringCollection()

	a.Put("a", "alpha")
	a.Put("b", "beta")
	a.Put("c", "gamma")

	b := a.Random(1)

	if b.Count() != 1 {
		t.Errorf("Expecting count to be 1 for %v", b)
	}
}

func Test_String_Shuffle(t *testing.T) {
	a := NewStringCollection()

	a.Put("a", "alpha")
	a.Put("b", "beta")
	a.Put("c", "gamma")

	a.Shuffle()

	if a.Count() != 3 {
		t.Errorf("Expecting count to be 3 for %v", a)
	}
}

func Test_String_Count(t *testing.T) {
	a := NewStringCollection()

	a.Put("a", "alpha")
	a.Put("b", "beta")
	a.Put("c", "gamma")

	if a.Count() != 3 {
		t.Errorf("Expected %v to have 3 items", a)
	}
}

func Test_String_Unique(t *testing.T) {
	a := NewStringCollection()

	a.Put("a", "alpha")
	a.Put("b", "beta")
	a.Put("c", "beta")

	b := a.Unique()

	if b.Count() == a.Count() {
		t.Errorf("Expected %v to have %v items", b, a.Count())
	}
}

func Test_String_ToJSON(t *testing.T) {
	a := NewStringCollection()

	a.Put("a", "alpha")
	a.Put("b", "beta")
	a.Put("c", "gamma")

	v, err := a.ToJSON()

	if err != nil {
		t.Error(err.Error())
	}

	if !strings.Contains(v, "\"a\":\"alpha\"") {
		t.Errorf("Expected JSON of %v and got %v", a, v)
	}
}
