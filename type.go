package main

type I int
type AliasI = int

func main() {
	v := 100
	var i I = v      // 报错
	var d AliasI = v // 不报错
	var z int = d    // 不报错
	TypeAlias(d)     // 不报错

	Type(100)  // 不报错
	Type(v)    // 报错
	Type(I(v)) // 不报错

	var mv I = 10 // 不报错
	Type(mv)      // 不报错
}

func Type(i I) {}

func TypeAlias(i int) {}
