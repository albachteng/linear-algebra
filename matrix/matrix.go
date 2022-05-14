package matrix

import (
	"github.com/albachteng/linear-algebra/vector"
)

type matrix struct {
	iHat vector.Vector
	jHat vector.Vector
}

func NewMatrix(v1, v2 vector.Vector) *matrix {
	return &matrix{v1, v2}
}

/* on a given matrix and given a transformation vector,
returns the new matrix */
func (m *matrix) LinearTransform(v vector.Vector) *vector.Vector {
	x := m.iHat.Scale(v.X)
	y := m.jHat.Scale(v.Y)
	return x.SumVector(*y) 
}
