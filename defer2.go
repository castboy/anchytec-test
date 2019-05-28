package main

import "fmt"

// 一个延迟调用可以修改包含此延迟调用的最内层函数的返回值,
// 延迟调用在函数return执行，r = n + n => r += n => return
func Triple(n int) (r int) {
	defer func() {
		r += n // 修改返回值
	}()

	return n + n // <=> r = n + n; return
}

func main() {
	fmt.Println(Triple(5)) // 15
}
