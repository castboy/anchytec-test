package main

import (
	"bou.ke/monkey"
	"fmt"
	"github.com/agiledragon/gomonkey"
	"reflect"
)

// note: go build -gcflags=all=-l gomonkey_monkey.go

func main() {
	// monkey below.
	patchGuard := monkey.Patch(target, func(a string) string {
		return "gomonkey"
	})

	s := target("a")
	fmt.Println(s == "gomonkey")

	patchGuard.Unpatch()
	s = target("a")
	fmt.Println(s == "a")


	var myInt MyInt = 10
	patchGuard = monkey.PatchInstanceMethod(reflect.TypeOf(myInt), "Method", func(myI MyInt) int {
		return int(myI) + 1
	})

	v := myInt.Method()
	fmt.Println(v == 11)

	monkey.UnpatchInstanceMethod(reflect.TypeOf(myInt), "Method")
	v = myInt.Method()
	fmt.Println(v == 10)

	// gomonkey below.
	patches := gomonkey.ApplyFunc(target, func(a string) string {
		return "monkey"
	})

	s = target("a")
	fmt.Println(s == "monkey")

	patches.Reset()
	s = target("a")
	fmt.Println(s == "a")

}

func target(a string) string {
	return a
}

type MyInt int
func (my MyInt) Method() int {
	return int(my)
}

