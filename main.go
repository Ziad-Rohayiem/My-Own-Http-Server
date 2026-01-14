package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"strings"
	"time"
)

func main() {
	// Parse command-line flags
	directory := flag.String("directory", "", "Directory to serve files from")
	flag.Parse()

	config := &Config{
		Directory: *directory,
	}

	fmt.Println("Starting HTTP server on 0.0.0.0:4221...")
	if config.Directory != "" {
		fmt.Printf("Serving files from: %s\n", config.Directory)
	}

	// Bind to port 4221
	listener, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Printf("Failed to bind to port 4221: %v\n", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening...")

	// Accept connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error accepting connection: %v\n", err)
			continue
		}

		// Handle each connection in a separate goroutine for concurrency
		go handleConnection(conn, config)
	}
}

// handleConnection handles a single client connection
// Supports persistent connections (HTTP/1.1 keep-alive)
func handleConnection(conn net.Conn, config *Config) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	// Handle multiple requests on the same connection (persistent connections)
	for {
		// Set read timeout to avoid hanging on persistent connections
		conn.SetReadDeadline(time.Now().Add(5 * time.Second))

		// Parse the HTTP request
		req, err := ParseRequest(reader)
		if err != nil {
			// Connection closed or timeout - this is normal for persistent connections
			return
		}

		// Handle the request
		resp := HandleRequest(req, config)

		// Check Connection header to determine if we should keep the connection alive
		connectionHeader := req.GetHeader("Connection")
		keepAlive := true

		if strings.EqualFold(connectionHeader, "close") {
			keepAlive = false
		}

		// Write the response
		_, err = conn.Write(resp.ToBytes())
		if err != nil {
			fmt.Printf("Error writing response: %v\n", err)
			return
		}

		// If Connection: close, close the connection after response
		if !keepAlive {
			return
		}

		// Continue to next request (persistent connection)
	}
}
