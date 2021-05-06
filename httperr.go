// Package httperr implements simple http errors
package httperr

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// DefaultError defines a simple DefaultError.
type DefaultError struct {
	Message    string `json:"message"`
	ErrorCode  string `json:"code"`
	StatusCode int    `json:"-"`
}

// Renderer implements the Render function.
// Its goal is to prepare the JSON rendering
type Renderer interface {
	Render(http.ResponseWriter, *http.Request) error
}

// JSON encodes the payload as JSON and sends it back.
// It calls the Render function of the renderer, so that you can set-up
// Anything you want to inside it, such as logging the error locally or such
func JSON(w http.ResponseWriter, r *http.Request, ren Renderer) {
	err := ren.Render(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(true)

	if err := enc.Encode(ren); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(buf.Bytes())
}

// Render implements Renderer for DefaultError
func (e *DefaultError) Render(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(e.StatusCode)
	return nil
}

func (e *DefaultError) Error() string {
	return fmt.Sprintf(
		"httperr:error=%s;error_code=%s;status_code=%d",
		e.Message,
		e.ErrorCode,
		e.StatusCode,
	)
}
