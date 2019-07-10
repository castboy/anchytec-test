package main

import (
	"fmt"
	"reflect"
)

// for key, element = range aContainer {...}
// 被遍历的容器值是aContainer的一个副本。 注意，只有aContainer的直接部分被复制了。 此副本是一个匿名的值，所以它是不可被修改的

type Person3 struct {
	name string
	age  int
}

func rangeArray() {
	persons := [2]Person3{{"Alice", 28}, {"Bob", 25}}
	for i, p := range persons {
		fmt.Println(i, p)
		// 此修改将不会体现在这个遍历过程中，
		// 因为被遍历的数组是persons的一个副本。
		persons[1].name = "Jack"
		// 此修改不会反映到persons数组中，因为p
		// 是persons数组的副本中的一个元素的副本。
		p.age = 31
	}
	fmt.Println("persons:", &persons)
}

func rangeSlice() {
	persons := []Person3{{"Alice", 28}, {"Bob", 25}}
	for i, p := range persons {
		fmt.Println(i, p)
		// 这次，此修改将反映在此次遍历过程中。
		persons[1].name = "Jack"
		// 这个修改仍然不会体现在persons切片容器中。
		p.age = 31
	}
	fmt.Println("persons:", &persons)
}

// 我们可以通过在range关键字后跟随一个数组的指针来遍历此数组中的元素。
// 对于大尺寸的数组，这种方法比较高效，因为复制一个指针比复制一个大尺寸数组的代价低得多。
// 下面的例子中的两个循环是等价的，它们的效率也基本相同

func pointerSlice() {
	var a [100]int

	for i, n := range &a { // 复制一个指针的开销很小
		fmt.Println(i, n)
	}

	for i, n := range a[:] { // 复制一个切片的开销很小
		fmt.Println(i, n)
	}
}

//如果一个for-range循环中的第二个循环变量既没有被忽略，也没有被舍弃，
// 并且range关键字后跟随一个nil数组指针，则此循环将造成一个恐慌。
// 在下面这个例子中，前两个循环都将打印出5个下标，但最后一个循环将导致一个恐慌。

func panic() {
	var p *[5]int // nil

	for i, _ := range p { // okay
		fmt.Println(i)
	}

	for i := range p { // okay
		fmt.Println(i)
	}

	for i, n := range p { // panic
		fmt.Println(i, n)
	}
}

//我们可以通过数组的指针来访问和修改此数组中的元素

func modify() {
	a := [5]int{2, 3, 5, 7, 11}
	p := &a
	p[0], p[1] = 17, 19
	fmt.Println(a) // [17 19 5 7 11]
}

//单独修改一个切片的长度或者容量

func reflectLenCap() {
	s := make([]int, 2, 6)
	fmt.Println(len(s), cap(s)) // 2 6

	reflect.ValueOf(&s).Elem().SetLen(3)
	fmt.Println(len(s), cap(s)) // 3 6

	reflect.ValueOf(&s).Elem().SetCap(5)
	fmt.Println(len(s), cap(s)) // 3 5
}

/////////////////////
func main() {
	reflectLenCap()

	rangeArray()

	fmt.Println("-------------")

	rangeSlice()
}
