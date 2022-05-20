package matrix

import (
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

/* multiplies the matrix by another matrix to return a new one
applies a rotation and then a shear, where s is the shear */
func (m *Matrix) Composition(s *Matrix) *Matrix {
	i := m.LinearTransform(s.IHat)
	j := m.LinearTransform(s.JHat)
	return &Matrix{*i, *j}
}
