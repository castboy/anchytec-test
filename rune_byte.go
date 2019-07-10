package main

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

func main() {
	// []byte 底层 []uint8, []rune 底层 []int32.
	s := "中国"
	fmt.Println([]byte(s), []rune(s))

	// []byte and []rune can not convert to each other.
	//t := []rune("中")
	//[]byte(t)

	// []byte and string convert to each other by deep replication.

	fmt.Println(len(s)) // without check if valid value to encode utf-8

	rs := []rune(s)
	Runes2Bytes(rs)

	bs := []byte(s)
	bytes.Runes(bs)

	//字符串和字节切片之间的转换的编译器优化，避免了深复制.
	//1.一个for-range循环中跟随range关键字的从字符串到字节切片的转换
	t := "hello"
	for i, b := range []byte(t) {
		fmt.Println(i, ":", b)
	}

	//2.一个在映射元素索引语法中被用做键值的从字节切片到字符串的转换；
	key := []byte{'k', 'e', 'y'}
	m := map[string]string{}
	m[string(key)] = "value"
	fmt.Println(m[string(key)])

	//3.一个字符串比较表达式中被用做比较值的从字节切片到字符串的转换
	var x = []byte{4: 'x'}
	var y = []byte{4: 'y'}
	if string(x) != string(y) {
		fmt.Println("x is not equal to y")
	}

	// range rune
	s = "éक्षिaπ囧"
	for i, rn := range s {
		fmt.Printf("%2v: 0x%x %v \n", i, rn, string(rn))
	}
	fmt.Println(len(s))

	// range byte lower
	for i := 0; i < len(s); i++ {
		fmt.Printf("第%v个字节为0x%x\n", i, s[i])
	}

	// range byte faster
	// 这里，[]byte(s)不需要深复制底层字节。
	for i, b := range []byte(s) {
		fmt.Printf("The byte at index %v: 0x%x \n", i, b)
	}
}

func Runes2Bytes(rs []rune) []byte {
	n := 0
	for _, r := range rs {
		n += utf8.RuneLen(r) // check if valid value to encode utf-8
	}

	n, bs := 0, make([]byte, n)

	for _, r := range rs {
		n = utf8.EncodeRune(bs[n:], r)
	}

	return bs
}
