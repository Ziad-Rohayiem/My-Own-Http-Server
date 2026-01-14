package main

import (
	"fmt"
	"strings"
)

// HTTPResponse represents an HTTP response
type HTTPResponse struct {
	StatusCode int
	StatusText string
	Headers    map[string]string
	Body       []byte
}

// NewResponse creates a new HTTP response
func NewResponse(statusCode int, statusText string) *HTTPResponse {
	return &HTTPResponse{
		StatusCode: statusCode,
		StatusText: statusText,
		Headers:    make(map[string]string),
		Body:       []byte{},
	}
}

// SetHeader sets a response header
func (r *HTTPResponse) SetHeader(name, value string) {
	r.Headers[name] = value
}

// SetBody sets the response body
func (r *HTTPResponse) SetBody(body []byte) {
	r.Body = body
	r.SetHeader("Content-Length", fmt.Sprintf("%d", len(body)))
}

// ToBytes converts the response to bytes following HTTP/1.1 format
func (r *HTTPResponse) ToBytes() []byte {
	var builder strings.Builder

	// Status line
	builder.WriteString(fmt.Sprintf("HTTP/1.1 %d %s\r\n", r.StatusCode, r.StatusText))

	// Headers
	for name, value := range r.Headers {
		builder.WriteString(fmt.Sprintf("%s: %s\r\n", name, value))
	}

	// Empty line separating headers from body
	builder.WriteString("\r\n")

	// Convert to bytes and append body
	result := []byte(builder.String())
	result = append(result, r.Body...)

	return result
}

// Response200 creates a 200 OK response
func Response200() *HTTPResponse {
	return NewResponse(200, "OK")
}

// Response404 creates a 404 Not Found response
func Response404() *HTTPResponse {
	return NewResponse(404, "Not Found")
}

// Response201 creates a 201 Created response
func Response201() *HTTPResponse {
	return NewResponse(201, "Created")
}

// Response500 creates a 500 Internal Server Error response
func Response500() *HTTPResponse {
	return NewResponse(500, "Internal Server Error")
}
