package service

import "net/http"

// NoteError is a generic error for failing to work with notes.
type NoteError struct {
	errorText string
}

// Error is an implementation of the Error interface.
func (e NoteError) Error() string {
	if e.errorText == "" {
		return "failed to store note"
	}
	return e.errorText
}

// StatusCode returns a http status code
func (NoteError) StatusCode() int {
	return http.StatusServiceUnavailable
}
