package main

import (
	"net"
	"io"
)

// mux not mutex.
//type Mux struct {
//	mtx sync.Mutex
//	conns map[net.Addr]net.Conn
//}

type Mux struct {
	join    chan net.Conn
	remove  chan net.Addr
	sendMsg chan string
}

func NewMux() *Mux {
	return &Mux{
		join:    make(chan net.Conn),
		remove:  make(chan net.Addr),
		sendMsg: make(chan string),
	}
}

func (mux *Mux) Join(conn net.Conn) {
	mux.join <- conn
}

func (mux *Mux) Remove(addr net.Addr) {
	mux.remove <- addr
}

func (mux *Mux) SendMsg(msg string) {
	mux.sendMsg <- msg
}

func (mux *Mux) Loop() {
	conns := make(map[net.Addr]net.Conn)
	for {
		select {
		case conn := <-mux.join:
			conns[conn.RemoteAddr()] = conn
		case addr := <-mux.remove:
			delete(conns, addr)
		case s := <-mux.sendMsg:
			for _, conn := range conns {
				_, err := io.WriteString(conn, s)
				if err != nil {
					//
				}
			}
		}
	}
}
