package main

//在某些情形下，内嵌在其它组合字面值中的组合字面形式可以简化为{...}（即类型部分被省略掉了）。 内嵌组合字面值前的取地址操作符&有时也可以被省略。
type language struct {
	year int
	name string
}

type languageCategory struct {
	dynamic,
	strong bool
}

func main() {
	var _ = []*[4]byte{
		&[4]byte{'A', 'B', 'C', 'D'},
		&[4]byte{'E', 'F', 'G', 'H'},
	}
	var _ = []*[4]byte{
		{'A', 'B', 'C', 'D'},
		{'E', 'F', 'G', 'H'},
	}

	var _ = [...]language{
		language{year: 1972, name: "C"},
		language{year: 2009, name: "GO"},
	}
	var _ = [...]language{
		{year: 1972, name: "C"},
		{year: 2009, name: "GO"},
	}

	var _ = map[languageCategory]map[string]int{
		languageCategory{true, true}: map[string]int{
			"Python": 1991,
			"Erlang": 1986,
		},
		languageCategory{true, false}: map[string]int{
			"JavaScript": 1995,
		},
		languageCategory{false, true}: map[string]int{
			"Go":   2009,
			"Rust": 2010,
		},
		languageCategory{false, false}: map[string]int{
			"C": 1972,
		},
	}
	var _ = map[languageCategory]map[string]int{
		{true, true}: {
			"Python": 1991,
			"Erlang": 1986,
		},
		{true, false}: {
			"JavaScript": 1995,
		},
		{false, true}: {
			"Go":   2009,
			"Rust": 2010,
		},
		{false, false}: {
			"C": 1972,
		},
	}

}
