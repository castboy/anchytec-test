package test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func direct() {
	var s string
	s = "hello " + "world!"
	_ = s
}

func fmtSprint() {
	var s string
	fmt.Sprintf("%s %s", "hello", "world!")
	_ = s
}

func stringsJoin() {
	var s string
	strings.Join([]string{"hello", "world!"}, " ")
	_ = s
}

func bytesBuffer() {
	buf := bytes.NewBuffer([]byte{'h', 'e', 'l', 'l', 'o', ' '})
	buf.Write([]byte{'w', 'o', 'r', 'l', 'd', '!'})
	_ = buf.String()
}

func stringsBuilder() {
	builder := strings.Builder{}
	builder.WriteString("hello ")
	builder.WriteString("world!")
	_ = builder.String()
}

func BenchmarkDirect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		direct()
	}
}

func BenchmarkFmtSprint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmtSprint()
	}
}

func BenchmarkStringsJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringsJoin()
	}
}

func BenchmarkBytesBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bytesBuffer()
	}
}

func BenchmarkStringsBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringsBuilder()
	}
}
