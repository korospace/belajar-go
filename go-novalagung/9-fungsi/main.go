package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
)

var randomizer = rand.New(rand.NewSource(123456789))

func main() {
	// void
	arrName := []string{"bagas", "zia"}
	printMsg("Hai", arrName)

	// return function
	randomValue := randomWithRange(1, 4)
	fmt.Println(randomValue)

	// multiple return
	area, circumference := calculateSquare(2)
	fmt.Println("area:", area)
	fmt.Println("circumference:", circumference)

	// pre-defined return
	area2, circumference2 := calculateCircle(3)
	fmt.Println("area2:", area2)
	fmt.Println("circumference2:", circumference2)

	// variadic function
	fmt.Println("avg 1:", avgNumber(2, 4, 4))

	// variadic function with slice data
	numbers := []int{4, 5, 6, 7}
	strAvg := fmt.Sprintln(avgNumber(numbers...))
	fmt.Printf("avg 2:" + strAvg)

	// Variadic with ohter param
	fmt.Println(makeHobies("bagas", "sleeping", "gaming"))

}

/*
============
Void Fungction
============
*/
func printMsg(msg string, arr []string) {
	names := strings.Join(arr, " dan ")

	fmt.Println(msg, names)
}

/*
================
Return Fungction
================
*/
func randomWithRange(min, max int) int {
	return rand.Int()%(max-min+1) + min
}

/*
================
Multiple Return
================
*/
func calculateSquare(l int) (int, int) {
	area := 4 * l

	circumference := l * l

	return area, circumference
}

/*
==================
Pre-defined Return
==================
*/
func calculateCircle(d float64) (area float64, circumference float64) {
	area = math.Pi * math.Pow(d/2, 2)
	circumference = math.Pi * d

	return
}

/*
==================
Variadic Function
==================
*/
func avgNumber(numbers ...int) float64 {
	var total float64 = 0

	for _, val := range numbers {
		total = total + float64(val)
	}

	return total / float64(len(numbers))
}

/*
=========================
Variadic with ohter param
=========================
*/
func makeHobies(name string, hobies ...string) string {
	strHobies := strings.Join(hobies, " ")

	return name + "'s hobies are " + strHobies
}
