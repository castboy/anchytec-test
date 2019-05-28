package main

import "fmt"

//若干`包级`变量在声明时刻的依赖关系影响初始化顺序, ay = 5, ac = ay, ab = ac+1, aa = ab+1, ax = aa+1
var ax, ay = aa + 1, 5
var aa, ab, ac = ab + 1, ac + 1, ay

func main() {
	fmt.Println(15 == 0xf)
	fmt.Println(15 == 017)
	fmt.Println('a' == 97)
	fmt.Println('a' == '\141') // 八进制
	fmt.Println('a' == '\x61') // 十六进制
	fmt.Println("A" == string(65), "A" == string('A'))
	fmt.Println(98 == 'a'+1)

	const (
		YES = true
		NO  = !YES
	)
	fmt.Println(YES, NO)

	const (
		X float64 = 3.14
		Y
		Z

		A, B = "Go", "Language"
		C, _
	)
	fmt.Println(X, Y, Z, A, B, C) //3.14 3.14 3.14 Go Language Go

	const (
		k = 3 // 在此处，iota == 0

		m float32 = iota + .5 // m float32 = 1 + .5
		n                     // n float32 = 2 + .5

		p    = 9          // 在此处，iota == 3
		q    = iota * 2   // q = 4 * 2
		_                 // _ = 5 * 2
		r                 // r = 6 * 2
		s, t = iota, iota // s, t = 7, 7
		u, v              // u, v = 8, 8
		_, w              // _, w = 9, 9
	)

	const x = iota // x = 0 （iota == 0）
	const (
		y = iota // y = 0 （iota == 0）
		z        // z = 1
	)

	// use
	const (
		Failed    = iota - 1 // 0-1
		Unknow               // 1-1
		Successed            // 2-1
	)

	const (
		Readable   = 1 << iota // 1
		Writable               // 2
		Executable             // 4
	)

	di, da := "difa", 100
	_, _ = di, da //di, da做为源值使用一次

}
