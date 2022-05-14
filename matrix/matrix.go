package matrix

import (
	"github.com/albachteng/linear-algebra/vector"
)

/* a 2-dimensional matrix that describes 
a linear transformation */
type matrix struct {
	iHat vector.Vector
	jHat vector.Vector
}

func NewMatrix(v1, v2 vector.Vector) *matrix {
	return &matrix{v1, v2}
}

/* on a given matrix and given a vector to transform,
returns the transformed vector */
func (m *matrix) LinearTransform(v vector.Vector) *vector.Vector {
	x := m.iHat.Scale(v.X)
	y := m.jHat.Scale(v.Y)
	return x.SumVector(*y) 
}
