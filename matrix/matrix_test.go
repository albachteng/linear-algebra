package matrix

import (
	"testing"

	"github.com/albachteng/linear-algebra/vector"
)

func TestMatrix(t *testing.T) {
	vector1 := vector.Vector{X: 1, Y: 1}
	vector2 := vector.Vector{X: -2, Y: 0}
	vector3 := vector.Vector{X: 0, Y: 1}
	vector4 := vector.Vector{X: 2, Y: 0}
	rotation := Matrix{vector3, vector4}
	shear := &Matrix{vector1, vector2}
	got := rotation.Composition(shear)
	if got.IHat.X != 2.0 || got.IHat.Y != 1.0 {
		t.Errorf("Composition iHat.X = %f, iHat.Y = %f, want 2", got.IHat.X, got.IHat.Y)
	}
	det := Matrix{vector.Vector{X: 1, Y: 1}, vector.Vector{X: 2, Y: -1}}
	dot := det.DotProduct()
	if dot != -3.0 {
		t.Errorf("Determinant is %f, expected 3", dot)
	}
}
