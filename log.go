package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	Trace *log.Logger
	Info  *log.Logger
	Warn  *log.Logger
	Error *log.Logger
)

func main() {
	Trace = log.New(ioutil.Discard, "Trace ", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(os.Stdout, "Info ", log.Ldate|log.Ltime|log.Lshortfile)
	Warn = log.New(os.Stdout, "Warn ", log.Ldate|log.Ltime|log.Lshortfile)

	f, _ := os.Create("log")
	Error = log.New(io.MultiWriter(f, os.Stdout), "Error ", log.Ltime|log.Ltime|log.Lshortfile)

	Trace.Println("I have something standard to say")
	Info.Println("Special Information")
	Warn.Println("There is something you need to know about")
	Error.Println("Something has failed")
}
