package http

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

// HTTP methods.
const (
	MethodGet = "GET"
)

// Request represents an HTTP request message.
type Request struct {
	Method  string // HTTP method
	Path    string // request path
	Version string // HTTP version
}

// parseRequest parses the request message from the client connection.
func parseRequest(c net.Conn) (Request, error) {
	r := Request{}
	// Wrap client connection in a bufio.Reader.
	rd := bufio.NewReader(c)
	// Parse request line.
	line, err := rd.ReadString('\n')
	if err != nil {
		return r, err
	}
	lineFields := strings.Fields(line)
	r.Method = lineFields[0]
	r.Path = lineFields[1]
	r.Version = lineFields[2]
	return r, nil
}

// NewRequest initializes a new HTTP request.
func NewRequest(c net.Conn) (Request, error) {
	return parseRequest(c)
}

// StatusText returns the status text corresponding to status.
func StatusText(status uint16) string {
	switch status {
	case 200:
		return "OK"
	}
	return ""
}

// Response represents an HTTP response message.
type Response struct {
	c       net.Conn // client connection
	Status  uint16   // status code
	Version string   // HTTP version
}

// NewResponse initializes a new HTTP response. The response version is HTTP/1.1 and
// the status code is 200.
func NewResponse(c net.Conn) Response {
	return Response{c: c, Status: 200, Version: "HTTP/1.1"}
}

// Send sends the response to the client.
func (r Response) Send() {
	r.c.Write([]byte(fmt.Sprintf("%s %d %s\r\n\r\n", r.Version, r.Status, StatusText(r.Status))))
}
