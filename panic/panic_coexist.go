package main

import "fmt"
//一个协程中可以有多个活动的恐慌共存
func main() { // 调用深度为0
	defer fmt.Println("程序崩溃了，因为退出时恐慌3依然未恢复")

	defer func() { // 调用深度为1
		defer func() { // 调用深度为2
			// 恐慌6被消除了。
			fmt.Println( recover() ) // 6
		}()

		// 恐慌3的深度为0，恐慌6的深度为1。
		defer fmt.Println("现在，恐慌3和恐慌6共存")
		defer panic(6) // 将压制恐慌5
		defer panic(5) // 将压制恐慌4
		panic(4) // 不会压制恐慌3，因为恐慌4和恐慌3的深度
		// 不同。恐慌3为0，而恐慌4的深度为1。
	}()

	defer fmt.Println("现在，只存在恐慌3")
	defer panic(3) // 将压制恐慌2
	defer panic(2) // 将压制恐慌1
	panic(1)
}
