package logger

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"sync"
	"time"
)

// transport with reasonable timeouts for log transfer over the network
var transport = &http.Transport{
	Dial: (&net.Dialer{
		Timeout: 1 * time.Second,
	}).Dial,
	TLSHandshakeTimeout: 1 * time.Second,
}

// client with a reasonable timeout and our transport
var client = &http.Client{
	Timeout:   time.Second * 3,
	Transport: transport,
}

// NewPostWriter gives us a new Writer that we can pass to a logger instance
func NewPostWriter(targetURL string, writer io.Writer) io.Writer {
	if _, err := url.ParseRequestURI(targetURL); err != nil {
		log.Printf("NewPostWriter targetURL '%s' invalid: %s\n", targetURL, err.Error())
		targetURL = ""
	}

	return &PostWriter{target: targetURL, writer: writer}
}

// PostWriter is the standard data structure for handling HTTP POST requests for a
// logging instance
type PostWriter struct {
	target    string
	targetMux sync.Mutex
	writer    io.Writer
}

// SetURL lets a client change the target server address for logging after
// writer initialization.
func (w *PostWriter) SetURL(target string) error {
	_, err := url.ParseRequestURI(target)

	if err != nil {
		return fmt.Errorf("write target url invalid: %v", err)
	}

	w.targetMux.Lock()
	defer w.targetMux.Unlock()

	w.target = target

	return nil
}

// Write satisfies the io.Writer interface
func (w *PostWriter) Write(p []byte) (int, error) {
	if w.target != "" {
		// We have a target address, lets send an HTTP Post with the payload to that server.
		// This is done with a new Go channel and any errors will be dumped to standard out.
		// It's possible we could make this more robust with a backoff, retry,
		// circuit breaker, and a fixed buffer queue.
		// This works for now.
		go func(target string, p []byte) {
			res, err := client.Post(target, "application/json", bytes.NewReader(p))

			if err != nil {
				log.Printf("internal/log PostWriter error: %s", err.Error())
				return
			}

			defer res.Body.Close()

			if res.StatusCode >= 300 {
				log.Printf("internal/log PostWriter error: %s", res.Status)
			}
		}(w.target, p)
	}

	// while that HTTP Post is happening in the background, lets immediately write
	// out to whatever writer the client has as backup. Most likely log.Writer{}.
	return w.writer.Write(p)
}
