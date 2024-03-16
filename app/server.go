package main

import (
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"net"
	"os"
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/app/http"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	dir := flag.String("directory", wd, "file directory")
	flag.Parse()

	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go func(c net.Conn) {
			defer c.Close()
			req, err := http.NewRequest(c)
			if err != nil {
				fmt.Println(err)
				return
			}
			res := http.NewResponse(c)
			switch req.Method {
			case http.MethodGet:
				if s, ok := strings.CutPrefix(req.Path, "/echo/"); ok {
					res.SetBody("text/plain", []byte(s))
				} else if req.Path == "/user-agent" {
					res.SetBody("text/plain", []byte(req.Headers["User-Agent"]))
				} else if name, ok := strings.CutPrefix(req.Path, "/files/"); ok {
					path := strings.Join([]string{*dir, name}, string(os.PathSeparator))
					content, err := os.ReadFile(path)
					if err != nil {
						if errors.Is(err, fs.ErrNotExist) {
							res.Status = 404
						} else {
							fmt.Println(err)
							return
						}
					} else {
						res.SetBody("application/octet-stream", content)
					}
				} else if req.Path != "/" {
					res.Status = 404
				}
			case http.MethodPost:
				if name, ok := strings.CutPrefix(req.Path, "/files/"); ok {
					path := strings.Join([]string{*dir, name}, string(os.PathSeparator))
					if err := os.WriteFile(path, req.Body, 0666); err != nil {
						fmt.Println(err)
						return
					}
					res.Status = 201
				}
			}
			res.Send()
		}(conn)
	}
}
