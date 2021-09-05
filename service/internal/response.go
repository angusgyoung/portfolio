package internal

import "time"

type Response struct {
	Message string `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

func NewResponse(message string) *Response {
	return &Response{
		Message: message,
		Timestamp: time.Now(),
	}
}