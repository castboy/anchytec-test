package main

import (
	"bytes"
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("service beginning...")

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(printStackTrace(err))
		}
	}()

	a()

}

func printStackTrace(err interface{}) string {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "%v\n", err)
	for i := 1; ; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		fmt.Fprintf(buf, "%s:%d (0x%x)\n", file, line, pc)
	}
	return buf.String()
}

func a() {
	b()
}

func b() {
	panic("aaaa")
}
