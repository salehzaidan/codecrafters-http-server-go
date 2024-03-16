package main

import (
	"log"
	"net"
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/app/http"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	conn, err := l.Accept()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	req, err := http.NewRequest(conn)
	if err != nil {
		log.Fatal(err)
	}
	res := http.NewResponse(conn)
	if s, ok := strings.CutPrefix(req.Path, "/echo/"); ok {
		res.SetBody("text/plain", []byte(s))
	} else if req.Path == "/user-agent" {
		res.SetBody("text/plain", []byte(req.Headers["User-Agent"]))
	} else if req.Path != "/" {
		res.Status = 404
	}
	res.Send()
}
