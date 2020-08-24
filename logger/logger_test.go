package logger

import (
	"bytes"
	"errors"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestInfo(t *testing.T) {
	var out bytes.Buffer

	New("info-test", &out).Info("hello")

	expected := `"app":"info-test","level":"info","msg":"hello","src":{"file":"logger/logger_test.go"`

	assert.True(t, strings.Contains(out.String(), expected), out.String())
}

func TestWarn(t *testing.T) {
	var out bytes.Buffer

	New("warn-test", &out).Warn("look out")

	expected := `"app":"warn-test","level":"warn","msg":"look out","src":{"file":"logger/logger_test.go"`

	assert.True(t, strings.Contains(out.String(), expected), out.String())
}

func TestErrorString(t *testing.T) {
	var out bytes.Buffer

	New("error-string-test", &out).Error("oh no")

	expected := `"app":"error-string-test","level":"error","msg":"oh no","src":{"file":"logger/logger_test.go"`

	assert.True(t, strings.Contains(out.String(), expected), out.String())
}

func TestError(t *testing.T) {
	var out bytes.Buffer

	New("error-test", &out).Error(errors.New("oh no"))

	expected := `"app":"error-test","level":"error","msg":"oh no","src":{"file":"logger/logger_test.go"`

	assert.True(t, strings.Contains(out.String(), expected), out.String())
}

func TestWith(t *testing.T) {
	var out bytes.Buffer

	l := New("with-test", &out)

	l.With(Data{"country": "us", "limit": 10}).Info("hello")

	expected := `"app":"with-test","level":"info","msg":"hello","data":{"country":"us","limit":10}`

	assert.True(t, strings.Contains(out.String(), expected), out.String())
}

type model struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (m model) Log() map[string]interface{} {
	return Data{"model": m}
}

func TestWithStruct(t *testing.T) {
	var out bytes.Buffer
	var a model

	l := New("with-struct-test", &out)

	a.ID = 1
	a.Name = "Example"

	l.With(a).Info("hello")

	needle := `"data":{"model":{"id":1,"name":"Example"}}`

	assert.Equal(t, 1, strings.Count(out.String(), needle), out.String())
}

func TestWithStructCollision(t *testing.T) {
	var out bytes.Buffer
	var a model
	var b model

	l := New("with-struct-collision-test", &out)

	a.ID = 1
	a.Name = "Example A"

	b.ID = 2
	b.Name = "Example B"

	l.With(a, b).Info("hello")

	needle := `"data":{"model":[{"id":1,"name":"Example A"},{"id":2,"name":"Example B"}]}`

	assert.Equal(t, 1, strings.Count(out.String(), needle), out.String())
}
