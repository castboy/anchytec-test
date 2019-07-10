package main

import "fmt"

func main() {
	//下面是（使用==比较运算符）比较两个接口值的步骤：

	//1.如果其中一个接口值是一个nil接口值，则比较结果为另一个接口值是否也为一个nil接口值。
	//2.如果这两个接口值的动态类型不一样，则比较结果为false。
	//3.对于这两个接口值的动态类型一样的情形，
	//	1)如果它们的动态类型为一个不可比较类型，则将产生一个恐慌。
	//	2)否则，比较结果为它们的动态值的比较结果。
	//简而言之，两个接口值的比较结果只有在下面两种任一情况下才为true：
	//1.这两个接口值都为nil接口值。
	//2.这两个接口值的动态类型相同、动态类型为可比较类型、并且动态值相等。

	var a, b, c interface{} = "abc", 123, "a" + "b" + "c"
	fmt.Println(a == b) // 第二步的情形。输出"false"。
	fmt.Println(a == c) // 第三步的情形。输出"true"。

	var x *int = nil
	var y *bool = nil
	var ix, iy interface{} = x, y
	var i interface{} = nil
	fmt.Println(ix == iy) // 第二步的情形。输出"false"。
	fmt.Println(ix == i)  // 第一步的情形。输出"false"。
	fmt.Println(iy == i)  // 第一步的情形。输出"false"。

	var s []int = nil // []int为一个不可比较类型。
	i = s
	fmt.Println(i == nil) // 第一步的情形。输出"false"。
	//fmt.Println(i == i)   // 第三步的情形。将产生一个恐慌。

	//一个[]T类型的值不能直接被转换为类型[]I，即使类型T实现了接口类型I

	//比如，我们不能直接将一个[]string值转换为类型[]interface{}。 我们必须使用一个循环来实现此转换：
	words := []string{"C", "Go", "Python"}

	iw := make([]interface{}, 0, len(words))
	for _, w := range words {
		iw = append(iw, w)
	}

	//fmt.Println(words...) error
	fmt.Println(iw...)
	fmt.Println(words)

	//一个接口类型每个指定的每一个方法都对应着一个隐式声明的函数
	//
	//如果接口类型I指定了一个名为m的方法原型，则编译器将隐式声明一个与之对应的函数名为I.m的函数。
	//此函数比m的方法原型中的参数多一个。此多出来的参数为函数I.m的第一个参数，它的类型为I。
	//对于一个类型为I的值i，方法调用i.m(...)和函数调用I.m(i, ...)是等价的。

	var ii II = T("gopher")
	fmt.Println(ii.m(5))                             // true
	fmt.Println(II.m(ii, 5))                         // true
	fmt.Println((interface{ m(int) bool }).m(ii, 5)) // true

}

type II interface {
	m(int) bool
}

type T string

func (t T) m(n int) bool {
	return len(t) > n
}
