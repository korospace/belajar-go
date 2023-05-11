package main

import (
	"fmt"
	"math"
)

func main() {

	// Augmented Assigment
	var a = 10
	a += 10
	fmt.Println(a) // 20

	var b = 10
	b -= 10
	fmt.Println(b) // 0

	// Unary Operator
	a++
	fmt.Println(a) // 21
	a--
	fmt.Println(a) // 20

	var toNegative = 2
	toNegative = -toNegative
	fmt.Println(toNegative) // -2

	var toPositive = math.Abs(float64(toNegative))
	fmt.Println(toPositive) // 2

	var benar = true
	fmt.Println(!benar)
}
