package matrix

import (
	"testing"

	"github.com/albachteng/linear-algebra/vector"
)

func TestMatrix(t *testing.T) {
	vector1 := *vector.NewVector(1, 1)
	vector2 := *vector.NewVector(-2, 0)
	vector3 := *vector.NewVector(0, 1)
	vector4 := *vector.NewVector(2, 0)
	rotation := NewMatrix(vector3, vector4)
	shear := NewMatrix(vector1, vector2)
	got := rotation.Composition(shear)
	if got.iHat.X != 2.0 || got.iHat.Y != 1.0 {
		t.Errorf("Composition iHat.X = %f, iHat.Y = %f, want 2", got.iHat.X, got.iHat.Y)
	}
}
