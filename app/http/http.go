package http

import (
	"fmt"
	"net"
)

// Response represents an HTTP response message.
type Response struct {
	c       net.Conn // client connection
	Status  uint16   // status code
	Version string   // HTTP version
}

// StatusText returns the status text corresponding to status.
func StatusText(status uint16) string {
	switch status {
	case 200:
		return "OK"
	}
	return ""
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
