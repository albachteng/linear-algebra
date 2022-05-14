package vector

import "math"

type Vector struct {
	x, y float64
}

/* takes two floats and returns a reference to a new vector */
func NewVector(x, y float64) *Vector {
	v := Vector{x, y}
	return &v
}

/* returns the numerical value of a vector's length property */
func (v *Vector) Length() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

/* sums two vectors with vector addition and returns a pointer to the new
vector that results  */
func (v *Vector) SumVector(n Vector) *Vector {
	x := v.x + n.x
	y := v.y + n.y
	return NewVector(x, y)
}

/* the same as above, but allows you to describe the vector on the fly */
func (v *Vector) Add(x, y float64) *Vector {
	n := *NewVector(x, y)
	return v.SumVector(n)
}

/* multiplies the x and y values of a vector by a scalar (float)
and returns a pointer to the new vector */
func (v *Vector) Scale(s float64) *Vector {
	n := NewVector(v.x*s, v.y*s)
	return n
}