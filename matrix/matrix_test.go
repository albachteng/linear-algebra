package matrix

import (
	"testing"
	// "github.com/albachteng/linear-algebra/matrix"
	"github.com/albachteng/linear-algebra/vector"
)

func TestMatrix(t *testing.T) {
	vector1 := *vector.NewVector(1, 1)
	vector2 := *vector.NewVector(-2, 0)
	vector3 := *vector.NewVector(0, 1)
	vector4 := *vector.NewVector(2, 0)
	rotation := NewMatrix(vector3, vector4)
	shear := NewMatrix(vector1, vector2)
	got := int(rotation.Composition(shear).iHat.X)
	if uint(got) != 2 {
		t.Errorf("Composition iHat.X = %d, want 2", got)
	}
}
