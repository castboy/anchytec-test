package main

import (
	"context"
	"fmt"
	"github.com/prometheus/common/log"
	"io/ioutil"
	"net/http"
	"net/http/httptrace"
	"time"
)

// refer: http://www.voidcn.com/article/p-qnpinueq-bny.html
type transport struct {
	current *http.Request
}

func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	t.current = req

	return http.DefaultTransport.RoundTrip(req)
}

func main() {
	ctx := context.Background()

	trace := &httptrace.ClientTrace{
		GetConn: func(hostPort string) {
			fmt.Println("GetConn", hostPort)
		},

		GotConn: func(info httptrace.GotConnInfo) {
			fmt.Printf("GotConn, %+v\n", info)

			info.Conn.SetDeadline(time.Now().Add(15 * time.Second))
		},

		PutIdleConn: func(err error) {
			fmt.Println("PutIdleConn", err)
		},

		GotFirstResponseByte: func() {
			fmt.Println("GotFirstResponseByte")
		},

		Got100Continue: func() {
			fmt.Println("Got100Continue")
		},

		DNSStart: func(info httptrace.DNSStartInfo) {
			fmt.Printf("DNSStart, %+v\n", info)
		},

		DNSDone: func(info httptrace.DNSDoneInfo) {
			fmt.Printf("DNSDone, %+v\n", info)
		},

		ConnectStart: func(network, addr string) {
			fmt.Printf("ConnectStart, %+v, %+v\n", network, addr)
		},

		ConnectDone: func(network, addr string, err error) {
			fmt.Printf("ConnectDone, %+v, %+v, %+v\n", network, addr, err)
		},

		WroteHeaderField: func(key string, value []string) {
			fmt.Println("WroteHeaderField, ", key, value)
		},

		WroteHeaders: func() {
			fmt.Println("WroteHeaders")
		},

		WroteRequest: func(info httptrace.WroteRequestInfo) {
			fmt.Println("WroteRequest", info.Err)
		},
	}

	ctx = httptrace.WithClientTrace(ctx, trace)

	req, err := http.NewRequestWithContext(ctx, "GET", "http://127.0.0.1:8080", nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Accept-Encoding", "gzip, deflate")

	client := http.Client{Transport: &transport{}}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res.Header.Get("Accept-Encoding"))

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))
}
