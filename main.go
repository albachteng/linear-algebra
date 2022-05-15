package main

import (
	"fmt"
	"reflect"

	"github.com/albachteng/linear-algebra/matrix"
	"github.com/albachteng/linear-algebra/vector"
)

func main() {
	unitVector := *vector.NewVector(1, 1)
	fmt.Println(reflect.TypeOf(unitVector)) // expect vector
	fmt.Println(unitVector.Length())        // expect root 2
	// anotherVector := unitVector.Add(1, 0)   // add leaves unitVector untouched
	anotherVector := *vector.NewVector(-1, 2)
	// fmt.Println(anotherVector.Length())     // expect NOT root 2
	// fmt.Println(*unitVector.Add(2, 3))
	// fmt.Println(anotherVector.Scale(-2.1).Length())
	i := *vector.NewVector(1, -2)
	j := *vector.NewVector(3, 0)
	matrix := *matrix.NewMatrix(i, j)
	fmt.Println(matrix)
	fmt.Println(*matrix.LinearTransform(anotherVector)) // expect 5,2
}
