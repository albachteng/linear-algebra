package matrix

import (
	"errors"
	"fmt"

	"github.com/albachteng/linear-algebra/vector"
)

/* a 2-dimensional matrix that describes
a linear transformation */
type Matrix struct {
	IHat vector.Vector
	JHat vector.Vector
}

/* on a given matrix and given a vector to transform,
returns the transformed vector */
func (m *Matrix) LinearTransform(v vector.Vector) *vector.Vector {
	x := m.IHat.Scale(v.X)
	y := m.JHat.Scale(v.Y)
	return x.SumVector(*y)
}

/* multiplies the matrix by another matrix to return a new one */
func (m *Matrix) Composition(s *Matrix) *Matrix {
	i := m.LinearTransform(s.IHat)
	j := m.LinearTransform(s.JHat)
	return &Matrix{*i, *j}
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
	iHat := vector.Vector{X: d, Y: c}
	jHat := vector.Vector{X: b, Y: a}
	return &Matrix{*iHat.Scale(determinant), *jHat.Scale(determinant)}, nil
}

func (m *Matrix) SolveSystem(v *vector.Vector) (*vector.Vector, error) {
	inverse, err := m.Inverse()
	return inverse.LinearTransform(*v), err
}
