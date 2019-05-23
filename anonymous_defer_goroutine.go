package main

import (
	"fmt"
	"time"
)

//一个匿名函数体内的表达式是在此函数被执行的时候才会被逐个估值的，不管此函数是被普通调用还是延迟/协程调用
func main() {
	for i := 0; i < 3; i++ {
		go func() {
			fmt.Println("b:", i)
		}()

		defer func() {
			fmt.Println("c:", i)
		}()
	}

	time.Sleep(time.Second)
}
