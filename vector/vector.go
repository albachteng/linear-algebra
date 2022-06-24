package vector

import (
	"errors"
	"math"
)

type Vector []float64

/* returns the numerical value of a vector's length property */
func (v Vector) Length() float64 {
	sum := 0.0
	for _, n := range v {
		sum += n * n
	}
	return math.Sqrt(sum)
}

/* sums two vectors with vector addition and returns a pointer to the new
vector that results  */
func (v Vector) SumVector(w Vector) Vector {
	var sums = make(Vector, 2)
	for i, num := range v {
		sums[i] =  num + w[i]
	}
	return sums 
}

/* multiplies the values of a vector by a scalar (float)
and returns a pointer to the new vector */
func (v Vector) Scale(s float64) Vector {
	var scaled = make(Vector, 2)
	for i, num := range v {
		scaled[i] = num * s	
	} 
	return scaled
}

/* returns the slope of a given vector */
// func (v Vector) Slope() float64 {
// 	return v.Y / v.X
// }

/* most vectors have a span of all of 2 dimensional space,
but when their slope is the same, their span is 1 dimension
- a single line. If both vectors are zero, their span is only the origin. */
// func GetSpan(a *Vector, b *Vector) uint8 {
// 	if a.Slope() == 0 && b.Slope() == 0 {
// 		return 0
// 	}
// 	if a.Slope() == b.Slope() {
// 		return 1
// 	}
// 	return 2
// }

/* sums the products of each value in two vectors of the same size */ 
func DotProduct(v Vector, w Vector) (float64, error) {
	if v.Length() != w.Length() {
		return 0, errors.New("DotProduct called with vectors of unequal length")
	}
	var products = make([]float64, 2)
	for i, num := range v {
		products[i] = num * w[i]
	}
	var sum float64
	for _, num := range products {
		sum += num
	}
	return sum, nil
}
