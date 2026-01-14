package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Config holds server configuration
type Config struct {
	Directory string
}

// HandleRequest processes an HTTP request and returns a response
func HandleRequest(req *HTTPRequest, config *Config) *HTTPResponse {
	path := req.Path

	// Root path
	if path == "/" {
		return Response200()
	}

	// Echo endpoint: /echo/{message}
	if strings.HasPrefix(path, "/echo/") {
		message := strings.TrimPrefix(path, "/echo/")
		resp := Response200()
		resp.SetHeader("Content-Type", "text/plain")

		// Check if client supports gzip
		schemes := req.ParseAcceptEncoding()
		if SupportsGzip(schemes) {
			// Compress the message
			compressed, err := CompressGzip([]byte(message))
			if err == nil {
				resp.SetBody(compressed)
				resp.SetHeader("Content-Encoding", "gzip")
				return resp
			}
		}

		resp.SetBody([]byte(message))
		return resp
	}

	// User-Agent endpoint
	if path == "/user-agent" {
		userAgent := req.GetHeader("User-Agent")
		resp := Response200()
		resp.SetHeader("Content-Type", "text/plain")
		resp.SetBody([]byte(userAgent))
		return resp
	}

	// Files endpoint: /files/{filename}
	if strings.HasPrefix(path, "/files/") {
		filename := strings.TrimPrefix(path, "/files/")

		if req.Method == "GET" {
			// Read and return file
			return handleFileGet(filename, config)
		} else if req.Method == "POST" {
			// Save file
			return handleFilePost(filename, req, config)
		}
	}

	// Unknown path - return 404
	return Response404()
}

// handleFileGet handles GET requests for files
func handleFileGet(filename string, config *Config) *HTTPResponse {
	if config.Directory == "" {
		return Response404()
	}

	filePath := filepath.Join(config.Directory, filename)

	// Read file contents
	contents, err := os.ReadFile(filePath)
	if err != nil {
		// File doesn't exist or can't be read
		return Response404()
	}

	resp := Response200()
	resp.SetHeader("Content-Type", "application/octet-stream")
	resp.SetBody(contents)
	return resp
}

// handleFilePost handles POST requests to save files
func handleFilePost(filename string, req *HTTPRequest, config *Config) *HTTPResponse {
	if config.Directory == "" {
		return Response500()
	}

	filePath := filepath.Join(config.Directory, filename)

	// Ensure directory exists
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		fmt.Printf("Failed to create directory: %v\n", err)
		return Response500()
	}

	// Write file
	err := os.WriteFile(filePath, req.Body, 0644)
	if err != nil {
		fmt.Printf("Failed to write file: %v\n", err)
		return Response500()
	}

	return Response201()
}
