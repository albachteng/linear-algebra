package matrix

import (
	"errors"
	"fmt"

	. "github.com/albachteng/linear-algebra/vector"
)

/* a n-dimensional matrix that describes
a linear transformation */
type Matrix []Vector

/* on a given matrix and given a vector to transform,
returns the transformed vector */
func (m Matrix) LinearTransform(v Vector) Vector {
	scaledVectors := make([]Vector, 2)
	for i, vec := range m {
		if 
		scaledVectors[i] = vec.Scale(v[i])
	}
	sum := scaledVectors[0]
	for i := 1; i < len(scaledVectors); i++ {
		sum.SumVector(scaledVectors[i])
	}
	return sum
}

/* multiplies the matrix by another matrix to return a new one */
func (m Matrix) Composition(s Matrix) Matrix {
	vecs := make([]Vector)
	for i, v := range m {
		vecs[i] = m.LinearTransform(v[i])
	}
	return Matrix{vecs}
}

/* returns the determinant of a transformation */
func (m *Matrix) DotProduct() float64 {
	return m.IHat.X*m.JHat.Y - m.IHat.Y*m.JHat.X
}

/* returns the inverse of a matrix
of the form:
	| a b |
	| c d |
if the determinant is zero, there is no inverse
*/
func (m *Matrix) Inverse() (*Matrix, error) {
	determinant := m.DotProduct()
	if determinant == 0 {
		return nil, errors.New("determinant is zero")
	}
	a := m.IHat.X
	b := -1 * m.JHat.X
	c := -1 * m.IHat.Y
	d := m.JHat.Y
	fmt.Println(a, b, c, d)
	iHat := Vector{X: d, Y: c}
	jHat := Vector{X: b, Y: a}
	return &Matrix{iHat.Scale(determinant), jHat.Scale(determinant)}, nil
}

func (m *Matrix) SolveSystem(v Vector) (Vector, error) {
	inverse, err := m.Inverse()
	return inverse.LinearTransform(v), err
}
