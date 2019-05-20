package main

import "fmt"

type One struct {
	Name string
	Age  int
	Male bool
}

func Male(t *One) error {
	t.Male = true
	return nil
}

func Age(age int) func(t *One) error {
	return func(t *One) error {
		return t.setAge(age)
	}
}

func (t *One) setAge(age int) error {
	t.Age = age
	return nil
}

func NewOne(name string, options ...func(t *One) error) (*One, error) {
	one := One{Name: name}

	for _, option := range options {
		option(&one)
	}

	return &one, nil
}

func main() {
	one, _ := NewOne("wmg")
	fmt.Println(one)

	one2, _ := NewOne("wmg", Male, Age(28))
	fmt.Println(one2)
}
