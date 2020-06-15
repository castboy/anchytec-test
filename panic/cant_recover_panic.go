package main

import "fmt"

// 5,6 recover函数并未直接在一个延迟函数调用中调用

func main() {
	defer func() {
		fmt.Println("7:", recover())
	}()

	defer func() {
		func() {
			fmt.Println("6:", recover())
		}()
	}()

	defer fmt.Println("5:", recover())

	panic(789)
}
