package api

import "time"

type Response struct {
	Message string `json:"message"`
	Timestamp time.Time `json:"timestamp"`
	Version string `json:"version"`
}

func NewResponse(message string, version string) *Response {
	return &Response{
		Message: message,
		Timestamp: time.Now(),
		Version: version,
	}
}