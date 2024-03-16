package main

import (
	"log"
	"net"

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

	res := http.NewResponse(conn)
	res.Send()
}
