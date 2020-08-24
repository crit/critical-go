package logger

import "fmt"

// Writer satisfies io.Write interface
type Writer struct{}

func (Writer) Write(p []byte) (int, error) {
	return fmt.Println(string(p))
}
