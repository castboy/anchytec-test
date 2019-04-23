package main

import "fmt"

type Student struct {
	Study
	Eat
}

type Eat interface {
	EatFood()
}

type Food struct {}

func (f *Food) EatFood() {
	fmt.Println("eat food")
}

type Study interface {
	StudyMath()
}

type Math struct{}

func (m *Math) StudyMath() {
	fmt.Println("study math")
}


func NewStuent() *Student {
	return &Student{
		&Math{},
		&Food{},
	}
}

func main() {
	stu := NewStuent()
	stu.StudyMath()
	stu.EatFood()
}