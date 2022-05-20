package main

import (
	"testing"
	"github.com/albachteng/linear-algebra/matrix"
	"github.com/albachteng/linear-algebra/vector"
)

func TestVector(t *testing.T) {
	someVector := vector.Vector{-1, 2}
	i := vector.Vector{1, -2}
	j := vector.Vector{3, 0}
	matrix := matrix.Matrix{i, j}
	got := *matrix.LinearTransform(someVector) // expect 5,2
	if got.X != 5 || got.Y != 2 {
		t.Errorf("expected transform {5, 2}, got {%f, %f}", got.X, got.Y)
	}
}
