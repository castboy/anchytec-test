package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"unicode"

	"github.com/pkg/profile"
)

var pool sync.Pool

func init() {
	pool = sync.Pool{New: func() interface{} {
		return [1]byte{}
	}}
}

func readbyte2(r io.Reader) (rune, error) {
	//var buf [1]byte
	//_, err := r.Read(buf[:])
	//return rune(buf[0]), err

	buf := pool.Get().([1]byte)

	_, err := r.Read(buf[:])

	pool.Put(buf)

	return rune(buf[0]), err
}

func main() {
	defer profile.Start(profile.MemProfile).Stop()

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("could not open file %q: %v", os.Args[1], err)
	}

	b := bufio.NewReader(f)

	words := 0
	inword := false
	for {
		r, err := readbyte2(b)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("could not read file %q: %v", os.Args[1], err)
		}
		if unicode.IsSpace(r) && inword {
			words++
			inword = false
		}
		inword = unicode.IsLetter(r)
	}
	fmt.Printf("%q: %d words\n", os.Args[1], words)
}
