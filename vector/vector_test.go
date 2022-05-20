package vector

import (
	"testing"
	"math"
)

func TestVector(t *testing.T) {
	unitVector := *NewVector(1, 1)
	rootTwo := unitVector.Length() // expect root 2
	if rootTwo != math.Sqrt(2) {
		t.Errorf("expected length of unit vector to be root 2, got %f", rootTwo)
	}
}
