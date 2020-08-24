package logger

import (
	"encoding/json"
	"fmt"
	"io"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"
)

// Loggable is any structure that can be passed directly to the logger instance.
//
// logger.With(user, field, device)
// vs
// logger.With(log.Data{"user": user.id, "field": field.value, "device": device.id})
type Loggable interface {
	Log() map[string]interface{}
}

// Data creates a shortcut to the Loggable interface that can be used
// by the client by reducing the amount of noise in the log call:
//
// logger.With(map[string]interface{"key": "someValue"})
// vs
// logger.With(log.Data{"key": "someValue"})
type Data map[string]interface{}

// Log satisfies the Loggable interface
func (d Data) Log() map[string]interface{} {
	return d
}

// New creates a new Log instance with a specific name and writer. The writer
// being io.Writer is the key to making this system extensible.
func New(appName string, out io.Writer) *Log {
	return &Log{app: appName, Out: out}
}

// Log is the core structure that holds all methods used by clients. The `Out`
// property is exported so that we can reuse the io.Writer in other places
// which has been required to make some middleware useful.
type Log struct {
	app  string
	data map[string]interface{}
	mu   sync.Mutex
	Out  io.Writer
}

// writeLog is the structure of the log used when writing data out. This could be
// customized to anything needed by a logging service. This also assumes JSON
// formatting (which seems to be popular with logging services) but could be any format.
type writeLog struct {
	Time  time.Time `json:"time"`
	App   string    `json:"app"`
	Level string    `json:"level"`
	Msg   string    `json:"msg"`
	Data  Data      `json:"data,omitempty"`
	Src   src       `json:"src"`
}

// src is part of writeLog that details log call location in the client. This is
// super useful for detecting problems in live code.
type src struct {
	File string `json:"file"`
	Line int    `json:"line"`
}

// AppName retrieves the log instance's current app name. Useful in middleware.
func (l *Log) AppName() string {
	return l.app
}

// With saves specific data to be written out when Info, Error, or Warn is called.
func (l *Log) With(data ...Loggable) *Log {
	// set will allow us to detect when a key has already been created with a value
	// and change the value to a slice of values if the key is presented again for logging.
	// logger := log.New(...).With(log.Data{"key": "v1"}) => {... "key":"v1" ...}
	// logger.With(log.Data{"key": "v2"})                 => {... "key":["v1","v2"] ...}
	set := make(map[string]interface{})

	// Right now this implementation is memory hungry. There is much room for improvement.
	for key, value := range l.data {
		set[key] = value
	}

	for _, node := range data {
		for key, value := range node.Log() {
			// do we have a current key already?
			if current, ok := set[key]; ok {
				// is the current value already a slice?
				if s, ok := current.([]interface{}); ok {
					// append to old slice
					s = append(s, value)
					set[key] = s
					continue
				}

				// create a new slice
				set[key] = []interface{}{current, value}
				continue
			}

			// create the new entry
			set[key] = value
		}
	}

	return &Log{
		app:  l.app,
		data: Data(set),
		Out:  l.Out,
	}
}

// Warn sends the log data to the writer with appropriate level.
func (l *Log) Warn(msg string) {
	l.output(2, "warn", msg)
}

// Error sends the log data to the writer with appropriate level.
func (l *Log) Error(errorOrMsg interface{}) {
	l.output(2, "error", errorParse(errorOrMsg))
}

// Info sends the log data to the writer with appropriate level.
func (l *Log) Info(msg string) {
	l.output(2, "info", msg)
}

// output creates the structure log and sends it to the writer.
func (l *Log) output(callDepth int, level, msg string) {
	var out writeLog
	var ok bool

	// do some "expensive" processing here like getting UTC and
	// caller info (which really _is_ expensive)

	out.Time = time.Now().UTC()

	_, out.Src.File, out.Src.Line, ok = runtime.Caller(callDepth)

	if !ok {
		out.Src.File = "???"
		out.Src.Line = 0
	} else {
		out.Src.File = srcFileParse(out.Src.File)
	}

	out.Level = level
	out.Msg = msg
	out.Data = map[string]interface{}{}

	// lock to protect from mutations while copying data from logger to log entry
	l.mu.Lock()

	// avoiding `defer l.mu.Unlock()` so we can unlock as soon as possible to allow other
	// channels to continue their processes since there is no reason to lock for the
	// JSON marshalling and writing.

	out.App = l.app

	for key, value := range l.data {
		out.Data[key] = value
	}

	// clear existing data to prevent repeated writes
	l.data = make(map[string]interface{})

	l.mu.Unlock() // unlock for more "expensive" processing

	data, err := json.Marshal(out)

	if err != nil {
		// This should "never" happen, but "never" is only a matter of time in distributed
		// systems so lets handle it.
		// I highly encourage making this message something that your logging service
		// can easily detect and attaching a mission critical alert to it.
		data = []byte("Logger unable to marshal log output to JSON: " + err.Error())
	}

	l.Out.Write(data)
}

// errorParse handles transforming the message into an appropriate string.
func errorParse(errorOrMsg interface{}) string {
	if msg, ok := errorOrMsg.(string); ok {
		return msg
	}

	if err, ok := errorOrMsg.(error); ok {
		return err.Error()
	}

	return fmt.Sprintf("%v", errorOrMsg)
}

// srcFileParse returns either the filename and extension, or the last directory
// (which is also usually the package name in Go) with the filename and extension.
// "project/src/model/user.go" => "model/user.go"
// "main.go" => "main.go
func srcFileParse(filename string) string {
	// "project/src/model/user.go" => "project/src/model", "user.go"
	dir, file := filepath.Split(filename)

	// "project/src/model" => ["project", "src", "model"]
	parts := strings.FieldsFunc(dir, func(c rune) bool {
		return c == filepath.Separator
	})

	if len(parts) > 0 {
		// => "model/user.go"
		return filepath.Join(parts[len(parts)-1], file)
	}

	// "user.go"
	return file
}
