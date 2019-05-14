package main

import (
	"crypto/tls"
	"fmt"
	"net"
	"time"
)

type Server struct {
	timeout time.Duration
	listener net.Listener
}

func NewServer(addr string, options ...func(server *Server)) (*Server, error) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}

	srv := Server{listener:l}

	for _, option := range options {
		option(&srv)
	}

	return &srv, nil
}

func main() {
	srv, err := NewServer("localhost:8080") // default

	timeout := func(srv *Server) {
		srv.timeout = 60 * time.Second
	}

	tls := func(srv *Server) {
		srv.listener = tls.NewListener(srv.listener, &tls.Config{})
	}

	srv2, err := NewServer("localhost:8090", timeout, tls) // listen securely with a 60 second timeout

	fmt.Println(srv, srv2, err)
}


