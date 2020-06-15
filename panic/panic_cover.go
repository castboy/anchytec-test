package panic

import "fmt"
//新生成的恐慌将压制同一深度的老的恐慌
func main() {
	defer fmt.Println("程序退出时未崩溃")

	defer func() {
		fmt.Println( recover() ) // 3
	}()

	defer fmt.Println("恐慌3将压制恐慌2")
	defer panic(3)
	defer fmt.Println("恐慌2将压制恐慌1")
	defer panic(2)
	panic(1)
}