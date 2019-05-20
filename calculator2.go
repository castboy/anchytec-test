package main

import (
	"fmt"
	"math"
)

type Calculator2 struct {
	acc float64
}

func NewCalculator2(acc float64) *Calculator2 {
	return &Calculator2{acc:acc}
}

func Add2(b float64) func(float64) float64 {
	return func(a float64) float64 {
		return a + b
	}
}

func Sqrt2() func(float64) float64 {
	return func(a float64) float64 {
		return math.Sqrt(a)
	}
}

func (cal *Calculator2) Do(f func(float64) float64) {
	cal.acc = f(cal.acc)
}

func main() {
	cal := NewCalculator2(2)
	cal.Do(Add2(2))
	fmt.Println(cal)

	cal.Do(Sqrt2())
	fmt.Println(cal)
}

