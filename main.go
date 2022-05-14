package main

import (
	"fmt"
	"reflect"
	"github.com/albachteng/linear-algebra/vector"
)

func main() {
	unitVector := *vector.NewVector(1.0, 1.0)
	fmt.Println(reflect.TypeOf(unitVector)) // expect vector
	fmt.Println(unitVector.Length())        // expect root 2
	anotherVector := unitVector.Add(1, 0)   // add leaves unitVector untouched
	fmt.Println(anotherVector.Length())     // expect NOT root 2
	fmt.Println(*unitVector.SumVector(*vector.NewVector(2, 3)))
	fmt.Println(anotherVector.Scale(-2.1).Length())
}
