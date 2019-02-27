package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 10000)
	go func() {
		for {
			fmt.Println(len(ch))
		}
	}()

	for {
		go func() {
			<- ch
		}()
	}

	go func() {
		for j := 0; j < 10000; j++ {
			ch <- j
		}
	}()




	time.Sleep(10*time.Minute)
}
