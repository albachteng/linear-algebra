package matrix

import (
	"math"
	"testing"

	"github.com/albachteng/linear-algebra/vector"
)

func withinTolerance(a, b, e float64) bool {
	if a == b {
		return true
	}
	d := math.Abs(a - b)
	if b == 0 {
		return d < e
	}
	return (d / math.Abs(b)) < e
}

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
	m := Matrix{vector.Vector{X: 40, Y: 20}, vector.Vector{X: 70, Y: 60}}
	inverse, err := m.Inverse()
	if err != nil || withinTolerance(.6, inverse.IHat.X, 1e-12) || withinTolerance(-.2, inverse.IHat.Y, 1e-12) {
		t.Errorf("expect inverse.IHat to be {0.6, -0.2, got %f, %f", inverse.IHat.X, inverse.IHat.Y)
	}
}
