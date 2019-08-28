package main

import (
	"bou.ke/monkey"
	"fmt"
	"github.com/agiledragon/gomonkey"
	"github.com/pkg/errors"
	"reflect"
)

// note: go build -gcflags=all=-l gomonkey_monkey.go

var num = 100

func main() {
	// monkey below.
	// Patch
	patchGuard := monkey.Patch(target, func(a string) string {
		return "gomonkey"
	})

	s := target("a")
	fmt.Println(s == "gomonkey")

	patchGuard.Unpatch()
	s = target("a")
	fmt.Println(s == "a")

	// PatchInstanceMethod
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
	// ApplyFunc
	patches := gomonkey.ApplyFunc(target, func(a string) string {
		return "monkey"
	})

	s = target("a")
	fmt.Println(s == "monkey")

	patches.Reset()
	s = target("a")
	fmt.Println(s == "a")
	// ApplyGlobalVar
	patches = gomonkey.ApplyGlobalVar(&num, 150)
	fmt.Println(num == 150)
	patches.Reset()
	fmt.Println(num == 100)

	// ApplyFuncSeq
	outputs := []gomonkey.OutputCell{
		{Values: gomonkey.Params{"output1", nil}, Times: 2},
		{Values: gomonkey.Params{"output2", errors.New("new func seq")}},
	}
	gomonkey.ApplyFuncSeq(funcSeq, outputs)
	s1, err := funcSeq()
	fmt.Println(s1 =="output1", err == nil)
	s1, err = funcSeq()
	fmt.Println(s1 =="output1", err == nil)
	s1, err = funcSeq()
	fmt.Println(s1 =="output2", err.Error())

	patchPairs := [][2]interface{}{
		{
			target,
			func(a string) string {return "gomonkey"},
		},
		{
			funcSeq,
			func() (string, error) { return "", errors.New("new func seq")},
		},
	}

	for i := range patchPairs {
		gomonkey.ApplyFunc(patchPairs[i][0], patchPairs[i][1])
	}
	s = target("")
	fmt.Println(s == "gomonkey")
	s, err = funcSeq()
	fmt.Println(s == "", err)

	var myInt2 MyInt = 10
	gomonkey.ApplyMethod(reflect.TypeOf(myInt2), "Method", func(myI MyInt) int {
		return int(myI) + 1
	})

	v = myInt2.Method()
	fmt.Println(v == 11)
}

func target(a string) string {
	return a
}

type MyInt int
func (my MyInt) Method() int {
	return int(my)
}

func funcSeq() (string, error) {
	return "", errors.New("func seq")
}
