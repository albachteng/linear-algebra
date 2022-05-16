package main

import (
	"fmt"
	"reflect"
	"github.com/albachteng/linear-algebra/matrix"
	"github.com/albachteng/linear-algebra/vector"
)

func main() {
	anotherVector := *vector.NewVector(-1, 2)
	i := *vector.NewVector(1, -2)
	j := *vector.NewVector(3, 0)
	matrix := *matrix.NewMatrix(i, j)
	fmt.Println(matrix)
	fmt.Println(*matrix.LinearTransform(anotherVector)) // expect 5,2
}
