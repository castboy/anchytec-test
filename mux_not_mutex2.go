package main

import (
	"io"
	"net"
)

type Mux2 struct {
	ops chan func(map[net.Addr]net.Conn)
}

func NewMux2() *Mux2 {
	return &Mux2{
		ops: make(chan func(map[net.Addr]net.Conn)),
	}
}

func (mux *Mux2) Join2(conn net.Conn) {
	mux.ops <- func(conns map[net.Addr]net.Conn) {
		conns[conn.RemoteAddr()] = conn
	}
}

func (mux *Mux2) Remove(addr net.Addr) {
	mux.ops <- func(conns map[net.Addr]net.Conn) {
		delete(conns, addr)
	}
}

func (mux *Mux2) SendMsg(msg string) error {
	result := make(chan error, 1)

	mux.ops <- func(conns map[net.Addr]net.Conn) {
		for _, conn := range conns {
			_, err := io.WriteString(conn, msg)
			if err != nil {
				result <- err
				return
			}
		}
	}

	return <-result
}

func (mux *Mux2) Loop() {
	conns := make(map[net.Addr]net.Conn)
	for ops := range mux.ops {
		ops(conns)
	}
}
