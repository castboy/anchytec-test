package main

import "fmt"

// for key, element = range aContainer {...}
// 被遍历的容器值是aContainer的一个副本。 注意，只有aContainer的直接部分被复制了。 此副本是一个匿名的值，所以它是不可被修改的

type Person3 struct {
	name string
	age  int
}

func main() {
	rangeArray()

	fmt.Println("-------------")

	rangeSlice()
}

func rangeArray() {
	persons := [2]Person3 {{"Alice", 28}, {"Bob", 25}}
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
	persons := []Person3 {{"Alice", 28}, {"Bob", 25}}
	for i, p := range persons {
		fmt.Println(i, p)
		// 这次，此修改将反映在此次遍历过程中。
		persons[1].name = "Jack"
		// 这个修改仍然不会体现在persons切片容器中。
		p.age = 31
	}
	fmt.Println("persons:", &persons)
}