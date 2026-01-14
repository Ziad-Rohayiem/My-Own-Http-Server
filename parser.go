package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// HTTPRequest represents a parsed HTTP request
type HTTPRequest struct {
	Method  string
	Path    string
	Version string
	Headers map[string]string
	Body    []byte
}

// ParseRequest parses an HTTP request from a reader
func ParseRequest(reader *bufio.Reader) (*HTTPRequest, error) {
	// Read the request line (e.g., "GET /path HTTP/1.1")
	requestLine, err := reader.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("failed to read request line: %w", err)
	}

	// Remove trailing \r\n
	requestLine = strings.TrimSpace(requestLine)

	// Parse the request line
	parts := strings.Split(requestLine, " ")
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid request line: %s", requestLine)
	}

	req := &HTTPRequest{
		Method:  parts[0],
		Path:    parts[1],
		Version: parts[2],
		Headers: make(map[string]string),
	}

	// Parse headers
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return nil, fmt.Errorf("failed to read header: %w", err)
		}

		line = strings.TrimSpace(line)

		// Empty line indicates end of headers
		if line == "" {
			break
		}

		// Parse header (e.g., "Content-Type: text/plain")
		colonIdx := strings.Index(line, ":")
		if colonIdx == -1 {
			continue
		}

		headerName := strings.TrimSpace(line[:colonIdx])
		headerValue := strings.TrimSpace(line[colonIdx+1:])
		req.Headers[headerName] = headerValue
	}

	// Read body if Content-Length is specified
	if contentLengthStr, ok := req.Headers["Content-Length"]; ok {
		contentLength, err := strconv.Atoi(contentLengthStr)
		if err != nil {
			return nil, fmt.Errorf("invalid Content-Length: %w", err)
		}

		if contentLength > 0 {
			req.Body = make([]byte, contentLength)
			_, err = io.ReadFull(reader, req.Body)
			if err != nil {
				return nil, fmt.Errorf("failed to read body: %w", err)
			}
		}
	}

	return req, nil
}

// GetHeader retrieves a header value (case-insensitive)
func (r *HTTPRequest) GetHeader(name string) string {
	// HTTP headers are case-insensitive
	for k, v := range r.Headers {
		if strings.EqualFold(k, name) {
			return v
		}
	}
	return ""
}

// ParseAcceptEncoding parses the Accept-Encoding header and returns supported schemes
func (r *HTTPRequest) ParseAcceptEncoding() []string {
	acceptEncoding := r.GetHeader("Accept-Encoding")
	if acceptEncoding == "" {
		return []string{}
	}

	schemes := []string{}
	parts := strings.Split(acceptEncoding, ",")
	for _, part := range parts {
		scheme := strings.TrimSpace(part)
		// Remove quality values if present (e.g., "gzip;q=0.8")
		if idx := strings.Index(scheme, ";"); idx != -1 {
			scheme = scheme[:idx]
		}
		schemes = append(schemes, strings.TrimSpace(scheme))
	}

	return schemes
}
