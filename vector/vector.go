package vector

import "math"

type vector struct {
	x, y float64
}

/* takes two floats and returns a reference to a new vector */
func NewVector(x, y float64) *vector {
	v := vector{x, y}
	return &v
}

/* returns the numerical value of a vector's length property */
func (v *vector) Length() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

/* sums two vectors with vector addition and returns a pointer to the new
vector that results  */
func (v *vector) sumVector(n vector) *vector {
	x := v.x + n.x
	y := v.y + n.y
	return NewVector(x, y)
}

/* the same as above, but allows you to describe the vector on the fly */
func (v *vector) Add(x, y float64) *vector {
	n := *NewVector(x, y)
	return v.sumVector(n)
}

/* multiplies the x and y values of a vector by a scalar (float)
and returns a pointer to the new vector */
func (v *vector) Scale(s float64) *vector {
	n := NewVector(v.x*s, v.y*s)
	return n
}
