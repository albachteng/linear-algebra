package matrix

import (
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

/* returns the determinant of a transformation, represented by a Matrix */
func (m *Matrix) DotProduct() float64 {
	return m.IHat.X*m.JHat.Y - m.IHat.Y*m.JHat.X
}

func (m *Matrix) Inverse() *Matrix {
	determinant := 1/m.DotProduct()
	a := m.IHat.X
	c := -1 * m.IHat.Y
	b := -1 * m.JHat.X
	d := m.JHat.Y
	fmt.Println(a, b, c, d)
	iHat := vector.Vector{X: d, Y: c}
	jHat := vector.Vector{X: b, Y: a}
	return &Matrix{*iHat.Scale(determinant), *jHat.Scale(determinant)}

}
