package main

import (
	"fmt"
	"math"
	"reflect"
)

type vector struct {
	x, y float64
}

func newVector(x, y float64) *vector {
	v := vector{x, y}
	return &v
}

func (v *vector) length() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func getLength(v vector) float64 {
	return v.length()
}

func main() {
	fmt.Println("hello world")
	unitVector := *newVector(1.0, 1.0)
	fmt.Println(math.Round(1.5))
	fmt.Println(unitVector)
	fmt.Println(reflect.TypeOf(unitVector))
	fmt.Println(getLength(unitVector)) // expect root 2
	fmt.Println(unitVector.length())   // expect root 2
}
