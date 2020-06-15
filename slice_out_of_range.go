package main

import "fmt"

func main() {
	fmt.Println("service beginning...")

	s := []int{1, 2, 3, 4, 5, 6}

	for i := range s {
		if i == 2 {
			s = append(s[:2], s[3:]...)
		}
		fmt.Println(i, "index")
		fmt.Println(s[i])
	}
}

//service beginning...
//0 index
//1
//1 index
//2
//2 index
//4
//3 index
//5
//4 index
//6
//5 index
//panic: runtime error: index out of range [5] with length 5
