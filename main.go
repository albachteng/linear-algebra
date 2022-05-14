package main

import (
	"fmt"
	"math"
	"reflect"
	"github.com/albachteng/linear-algebra/hello"
)

type vector struct {
	x, y float64
}

/* takes two floats and returns a reference to a new vector */
func newVector(x, y float64) *vector {
	v := vector{x, y}
	return &v
}

/* returns the numerical value of a vector's length property */
func (v *vector) length() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

/* sums two vectors with vector addition and returns a pointer to the new
vector that results  */
func (v *vector) sumVector(n vector) *vector {
	x := v.x + n.x
	y := v.y + n.y
	return newVector(x, y)
}

/* the same as above, but allows you to describe the vector on the fly */
func (v *vector) add(x, y float64) *vector {
	n := *newVector(x, y)
	return v.sumVector(n)
}

/* multiplies the x and y values of a vector by a scalar (float)
and returns a pointer to the new vector */
func (v *vector) scale(s float64) *vector {
	n := newVector(v.x*s, v.y*s)
	return n
}

func main() {
	unitVector := *newVector(1.0, 1.0)
	fmt.Println(reflect.TypeOf(unitVector)) // expect vector
	fmt.Println(unitVector.length())        // expect root 2
	anotherVector := unitVector.add(1, 0)   // add leaves unitVector untouched
	fmt.Println(anotherVector.length())     // expect NOT root 2
	fmt.Println(*unitVector.sumVector(*newVector(2, 3)))
	fmt.Println(anotherVector.scale(-2.1).length())
	fmt.Println(hello.Hello())
}
