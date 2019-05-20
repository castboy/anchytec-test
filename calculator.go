package main

import (
	"fmt"
	"math"
)

type Calculator struct {
	acc float64
}

func NewCalculator(acc float64) *Calculator {
	return &Calculator{acc:acc}
}

type opFunc func(float64, float64) float64

func (cal *Calculator) Do(op opFunc, a float64) {
	cal.acc = op(cal.acc, a)
}

func Add(a, b float64) float64 {
	return a + b
}

func Sqrt(a, _ float64) float64 {
	return math.Sqrt(a)
}

func main() {
	cal := NewCalculator(3)
	cal.Do(Add, 1)
	fmt.Println(cal)

	cal.Do(Sqrt, 0)
	fmt.Println(cal)
}


