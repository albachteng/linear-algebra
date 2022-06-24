package main

import (
	"fmt"

	"github.com/albachteng/linear-algebra/matrix"
	. "github.com/albachteng/linear-algebra/vector"
)

func main() {

	anotherVector := Vector{-1, 2}
	i := Vector{1, -2}
	j := Vector{3, 0}
	matrix := matrix.Matrix{i, j}
	fmt.Println(matrix)
	fmt.Println(matrix.LinearTransform(anotherVector)) // expect 5,2
}
