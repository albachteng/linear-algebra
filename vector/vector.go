package vector

import "math"

type Vector struct {
	X, Y float64
}

/* returns the numerical value of a vector's length property */
func (v *Vector) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/* sums two vectors with vector addition and returns a pointer to the new
vector that results  */
func (v *Vector) SumVector(n Vector) *Vector {
	x := v.X + n.X
	y := v.Y + n.Y
	return &Vector{x, y}
}

/* multiplies the x and y values of a vector by a scalar (float)
and returns a pointer to the new vector */
func (v *Vector) Scale(s float64) *Vector {
	n := Vector{v.X * s, v.Y * s}
	return &n
}

/* returns the slope of a given vector */
func (v *Vector) Slope() float64 {
	return v.Y / v.X
}

/* most vectors have a span of all of 2 dimensional space, 
but when their slope is the same, their span is 1 dimension 
- a single line. If both vectors are zero, their span is only the origin. */ 
func GetSpan(a *Vector, b *Vector) uint8 {
	if a.Slope() == 0 && b.Slope() == 0 {
		return 0
	}
	if a.Slope() == b.Slope() {
		return 1
	}
	return 2
}