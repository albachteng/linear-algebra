package main

import (
	"fmt"
	"github.com/albachteng/linear-algebra/matrix"
	"github.com/albachteng/linear-algebra/vector"
)

func main() {
	anotherVector := vector.Vector{-1, 2}
	i := vector.Vector{1, -2}
	j := vector.Vector{3, 0}
	matrix := *matrix.NewMatrix(i, j)
	fmt.Println(matrix)
	fmt.Println(*matrix.LinearTransform(anotherVector)) // expect 5,2
}
