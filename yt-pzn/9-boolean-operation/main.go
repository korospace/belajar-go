package main

import "fmt"

func main() {
	var benar bool = true
	var salah bool = false

	fmt.Println(benar && salah) // false
	fmt.Println(benar || salah) // true
	fmt.Println(!benar)         // false

	var midleTest int32 = 90
	var finalTest int32 = 80

	var passedMidelTest bool = midleTest >= 90
	var passedFinalTest bool = finalTest >= 80

	var passed = passedFinalTest && passedMidelTest

	fmt.Println("passed test:", passed)

}
