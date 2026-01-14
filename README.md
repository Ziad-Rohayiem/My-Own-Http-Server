# My Own HTTP Server

A fully functional HTTP/1.1 server built from scratch in Go using TCP primitives. This project was created as part of the [CodeCrafters](https://codecrafters.io) "Build your own HTTP server" challenge.

## Features

✅ **Stage 1**: TCP server binding to port 4221  
✅ **Stage 2**: Basic HTTP 200 OK responses  
✅ **Stage 3**: URL path extraction and routing  
✅ **Stage 4**: Response bodies with proper headers  
✅ **Stage 5**: HTTP header parsing  
✅ **Stage 6**: Concurrent connections using goroutines  
✅ **Stage 7**: File serving from filesystem  
✅ **Stage 8**: POST requests and file uploads  
✅ **Stage 9**: Accept-Encoding header parsing  
✅ **Stage 10**: Multiple compression scheme support  
✅ **Stage 11**: Gzip compression  
✅ **Stage 12**: Persistent connections (keep-alive)  
✅ **Stage 13**: Concurrent persistent connections  
✅ **Stage 14**: Proper connection closure handling  

## Architecture

The server is organized into several modules:

- **main.go**: Entry point, TCP server setup, and connection handling
- **parser.go**: HTTP request parsing (request line, headers, body)
- **response.go**: HTTP response generation
- **handler.go**: Request routing and endpoint logic
- **compression.go**: Gzip compression utilities

## Prerequisites

- Go 1.21 or higher

## Installation

```bash
# Clone the repository
git clone https://github.com/Ziad-Rohayiem/My-Own-Http-Server.git
cd My-Own-Http-Server

# Build the server
go build -o my-http-server
```

## Usage

### Basic Usage

Start the server on port 4221:

```bash
go run .
```

Or run the compiled binary:

```bash
./my-http-server
```

### Serving Files

To enable file serving, use the `--directory` flag:

```bash
go run . --directory /path/to/files
```

## Supported Endpoints

### GET /
Returns a simple 200 OK response.

**Example:**
```bash
curl http://localhost:4221/
```

### GET /echo/{message}
Echoes back the provided message. Supports gzip compression if the client sends `Accept-Encoding: gzip`.

**Example:**
```bash
curl http://localhost:4221/echo/hello
# Response: hello

curl -H "Accept-Encoding: gzip" http://localhost:4221/echo/hello --compressed
# Response: hello (gzip compressed)
```

### GET /user-agent
Returns the User-Agent header from the request.

**Example:**
```bash
curl http://localhost:4221/user-agent
# Response: curl/7.68.0
```

### GET /files/{filename}
Serves a file from the directory specified with `--directory` flag.

**Example:**
```bash
go run . --directory /tmp
curl http://localhost:4221/files/test.txt
```

### POST /files/{filename}
Saves the request body to a file in the directory specified with `--directory` flag.

**Example:**
```bash
go run . --directory /tmp
curl -X POST http://localhost:4221/files/test.txt -d "Hello, World!"
```

## HTTP/1.1 Features

### Concurrent Connections
The server handles multiple simultaneous connections using goroutines. Each connection is handled in a separate goroutine, allowing non-blocking concurrent request processing.

### Persistent Connections (Keep-Alive)
The server supports HTTP/1.1 persistent connections. Multiple requests can be sent over the same TCP connection, reducing latency and overhead.

- Connections are kept alive by default
- Clients can send `Connection: close` to close after response
- Automatic timeout after 5 seconds of inactivity

### Gzip Compression
The server automatically compresses response bodies when:
1. Client sends `Accept-Encoding: gzip` header
2. The endpoint supports compression (currently `/echo/*`)

## Implementation Details

### Goroutines for Concurrency
```go
// Each connection is handled in its own goroutine
for {
    conn, err := listener.Accept()
    if err != nil {
        continue
    }
    go handleConnection(conn, config)
}
```

### HTTP Request Parsing
- Uses `bufio.Reader` for efficient reading
- Parses request line, headers, and body
- Supports `Content-Length` for body reading
- Case-insensitive header matching

### Response Format
All responses follow the HTTP/1.1 specification:
```
HTTP/1.1 200 OK\r\n
Content-Type: text/plain\r\n
Content-Length: 13\r\n
\r\n
Hello, World!
```

### Error Handling
- Returns 404 for unknown paths
- Returns 404 for non-existent files
- Returns 500 for server errors (file write failures, etc.)
- Gracefully handles malformed requests
- Implements connection timeouts for persistent connections

## Testing

You can test the server using curl, telnet, or any HTTP client:

```bash
# Test basic response
curl -v http://localhost:4221/

# Test echo endpoint
curl http://localhost:4221/echo/testing

# Test with gzip compression
curl -H "Accept-Encoding: gzip" http://localhost:4221/echo/compressed --compressed

# Test User-Agent
curl http://localhost:4221/user-agent

# Test file operations
go run . --directory /tmp
echo "test content" > /tmp/test.txt
curl http://localhost:4221/files/test.txt
curl -X POST http://localhost:4221/files/new.txt -d "new content"

# Test persistent connections
telnet localhost 4221
GET / HTTP/1.1
Host: localhost

GET /echo/hello HTTP/1.1
Host: localhost

Connection: close
GET /user-agent HTTP/1.1
Host: localhost
User-Agent: Telnet

```

## Stage Completion Checklist

- [x] Stage 1: Bind to port 4221
- [x] Stage 2: Respond with 200 OK
- [x] Stage 3: Extract URL path
- [x] Stage 4: Respond with body
- [x] Stage 5: Read headers
- [x] Stage 6: Concurrent connections
- [x] Stage 7: Return a file
- [x] Stage 8: Read request body (POST)
- [x] Stage 9: Parse Accept-Encoding header
- [x] Stage 10: Handle multiple compression schemes
- [x] Stage 11: Implement gzip compression
- [x] Stage 12: Persistent connections
- [x] Stage 13: Concurrent persistent connections
- [x] Stage 14: Connection closure handling

## Learning Outcomes

This project demonstrates:
- Low-level TCP/IP networking with Go's `net` package
- HTTP/1.1 protocol implementation (RFC 7230)
- Concurrent programming with goroutines
- File system operations
- Data compression with gzip
- Proper error handling in Go
- Clean code organization and separation of concerns

## License

This project is created for educational purposes as part of the CodeCrafters challenge.

## Author

Ziad Rohayiem

## Repository

https://github.com/Ziad-Rohayiem/My-Own-Http-Server
