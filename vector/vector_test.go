package vector

import (
	"math"
	"testing"
)

func TestVector(t *testing.T) {
	unitVector := Vector{1, 1}
	rootTwo := unitVector.Length() // expect root 2
	if rootTwo != math.Sqrt(2) {
		t.Errorf("expected length of unit vector to be root 2, got %f", rootTwo)
	}
	scaledUnitVector := unitVector.Scale(2)
	span := GetSpan(scaledUnitVector, &unitVector)
	if span != 1 {
		t.Errorf("expected a span of 1, got %d", span)
	}
}
