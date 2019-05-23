package main

import "fmt"

func main() {
	m := make(map[string]string)
	v, ok := m["a"]
	fmt.Println(v == "", ok)

	m[""] = ""

	v, ok = m[""]
	fmt.Println(v == "", ok, len(m))
}
