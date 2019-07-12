package mock

import "fmt"

type MyOperation interface {
	Add(a, b int) int
	Put(a int)
	RtnErr() error
	RtnChan() <-chan bool
	Ptr(a int, b *int) int
}

func AddOper(myOp MyOperation) {
	if myOp.Add(1, 1) == 2 {
		fmt.Println("meet expectation")
		// do other exec.
		myOp.Put(3)
	}

	myOp.Add(1, 2)

	var b int
	fmt.Println(myOp.Ptr(1, &b))

	err := myOp.RtnErr()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(myOp.RtnChan())
	fmt.Println(myOp.RtnChan())
	fmt.Println(myOp.RtnChan())
}
