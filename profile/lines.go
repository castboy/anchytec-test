package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func CountLines(r io.Reader) (int, error) {
	var (
		br    = bufio.NewReader(r)
		lines int
		err   error
	)

	for {
		_, err = br.ReadString('\n')
		lines++
		if err != nil {
			break
		}
	}

	if err != io.EOF {
		return 0, err
	}
	return lines, nil
}

func CountLines2(r io.Reader) (int, error) {
	sc := bufio.NewScanner(r)
	lines := 0

	for sc.Scan() {
		lines++
	}
	return lines, sc.Err()
}

func main() {
	fd, err := os.Open("moby.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()

	num, err := CountLines(fd)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(num, "1")

	fd.Seek(0, io.SeekStart)
	num, err = CountLines2(fd)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(num, "2")
}
