package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("moby.txt")
	if err != nil {
		fmt.Println(err)
	}

	f.SetDeadline()
	var f2 io.Reader = f

	p := make([]byte, 10)
	for i := 0; i < 10; i++ {
		n, err := f2.Read(p)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(n, string(p))
	}

	writer := bufio.NewReader(f)
	writer.WriteTo(os.Stdout)

}
