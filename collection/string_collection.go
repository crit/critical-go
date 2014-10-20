package collection

import (
	"encoding/json"
	"math/rand"
	"strings"
)

type stringCollection struct {
	items map[string]string
}

// NewStringCollection returns a new instance of a string collection
func NewStringCollection() stringCollection {
	return stringCollection{items: make(map[string]string)}
}

// MakeStringCollection returns a new string collection from an existing map
func MakeStringCollection(items map[string]string) stringCollection {
	a := NewStringCollection()
	a.items = items
	return a
}

func FromJSON(json string) (stringCollection, error) {
	
}

// All items in the collection are returned
func (c *stringCollection) All() map[string]string {
	return c.items
}

// Check the contents of the items for a specific value
func (c *stringCollection) Contains(value string) bool {
	for _, v := range c.items {
		if v == value {
			return true
		}
	}
	return false
}

// Count returns the number of items currently in the collection
func (c *stringCollection) Count() int {
	return len(c.items)
}

// Diff returns a new string collection with the keys that are not in both
// collections
func (c *stringCollection) Diff(input stringCollection) stringCollection {
	out := NewStringCollection()

	for k, v := range c.items {

		if !input.Has(k) {
			out.Put(k, v)
		}
	}

	for k, v := range input.items {
		if !c.Has(k) {
			out.Put(k, v)
		}
	}

	return out
}

// Each runs a function against all key/values in items
func (c *stringCollection) Each(fn func(string, string) (string, string)) {
	updates := make(map[string]string, len(c.items))

	for k, v := range c.items {
		a, b := fn(k, v)
		updates[a] = b
	}

	c.items = updates
}

// Empty returns whether there is nothing in items
func (c *stringCollection) Empty() bool {
	return len(c.items) == 0
}

// Filter returns a new string collection by filtered values
func (c *stringCollection) Filter(fn func(v string) bool) stringCollection {
	out := NewStringCollection()

	for k, v := range c.items {
		if fn(v) {
			out.Put(k, v)
		}
	}

	return out
}

// Flatten items into just values as a slice
func (c *stringCollection) Flatten() []string {
	out := make([]string, 0, len(c.items))

	for _, v := range c.items {
		out = append(out, v)
	}

	return out
}

// Flip item values to keys and keys to values
func (c *stringCollection) Flip() stringCollection {
	out := NewStringCollection()

	for k, v := range c.items {
		out.Put(v, k)
	}

	return out
}

// Forget removes an item from the stack by key
func (c *stringCollection) Forget(key string) {
	delete(c.items, key)

}

// Get the value of a key or return the passed in default
func (c *stringCollection) Get(key string, def string) string {
	v, ok := c.items[key]

	if !ok {
		return def
	}

	return v
}

// Has checks the items for a specific key
func (c *stringCollection) Has(key string) bool {
	for k := range c.items {
		if k == key {
			return true
		}
	}

	return false
}

// Implode items with a specific string
func (c *stringCollection) Implode(glue string) string {
	return strings.Join(c.Flatten(), glue)
}

// Intersect the items in the collection with another collection
func (c *stringCollection) Intersect(input stringCollection) stringCollection {
	out := NewStringCollection()

	for k, v := range input.items {
		if c.Has(k) {
			out.Put(k, v)
		}
	}

	for k, v := range c.items {
		if input.Has(k) {
			out.Put(k, v)
		}
	}

	return out
}

// Keys returns the item keys
func (c *stringCollection) Keys() []string {
	out := make([]string, 0, len(c.items))

	for k := range c.items {
		out = append(out, k)
	}

	return out
}

// Only returns a new collection for a set of keys
func (c *stringCollection) Only(keys []string) stringCollection {
	out := NewStringCollection()

	for _, k := range keys {
		v := c.Get(k, "")

		if v != "" {
			out.Put(k, v)
		}
	}

	return out
}

// Merge this collection with another while overwriting values of matching keys
// respecting the incoming collection's values
func (c *stringCollection) Merge(input stringCollection) {
	for k, v := range input.items {
		c.Put(k, v)
	}
}

// Pull removes an item from the collection and returns the item
func (c *stringCollection) Pull(key, def string) (value string, empty bool) {
	empty = c.Empty()

	if empty {
		return "", empty
	}

	v := c.Get(key, def)

	c.Forget(key)

	return v, empty
}

// Put a key and value onto the stack
func (c *stringCollection) Put(key, value string) {
	c.items[key] = value
}

// Random returns a random value from the collection
func (c *stringCollection) Random(amount int) stringCollection {
	out := NewStringCollection()

	if amount < 1 {
		return out
	}

	if amount > c.Count() {
		return out
	}

	keys := c.Keys()

	// shuffle the []string keys
	for i := range keys {
		j := rand.Intn(i + 1)
		keys[i], keys[j] = keys[j], keys[i]
	}

	for i := 0; i < amount; i++ {
		k := keys[i]
		v := c.Get(k, "")
		out.Put(k, v)
	}

	return out
}

// Reduce the items to a single value
func (c *stringCollection) Reduce(fn func(carry, item string) string, initial string) (value string) {
	a := c.Flatten()

	v := initial

	for i := range a {
		v = fn(v, a[i])
	}

	return v
}

// Reject returns a collection of all items that do not pass a given test
func (c *stringCollection) Reject(fn func(string) bool) stringCollection {
	out := NewStringCollection()

	for k, v := range c.items {
		if !fn(v) {
			out.Put(k, v)
		}
	}

	return out
}

// Search the collection for a value and return the key
func (c *stringCollection) Search(value string) (key string, found bool) {
	if c.Empty() {
		return "", false
	}

	for k, v := range c.items {
		if v == value {
			return k, true
		}
	}

	return "", false
}

// Shuffle the items in the collection
func (c *stringCollection) Shuffle() {
	a := c.Random(c.Count())

	c.items = a.items
}

// Unique returns a new collection of the unique values in current items
func (c *stringCollection) Unique() stringCollection {
	out := NewStringCollection()

	for k, v := range c.items {
		if !out.Contains(v) {
			out.Put(k, v)
		}
	}

	return out
}

// ToJSON serializes the current items into a JSON string
func (c *stringCollection) ToJSON() (string, error) {
	b, err := json.Marshal(c.items)

	if err != nil {
		return "", err
	}

	return string(b[:]), err
}