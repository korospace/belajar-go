package main

import "fmt"

func main() {
	/* Non Floating Point */
	var onlyPositive uint = 89
	var positiveNegative int = -89

	fmt.Printf("onlyPositive: %d | positiveNegative: %d\n", onlyPositive, positiveNegative)

	/* Floating Point */
	var decimalNumber float32 = 2.26

	fmt.Printf("desimal 1: %f\n", decimalNumber)   // 2.260000
	fmt.Printf("desimal 2: %.3f\n", decimalNumber) // 2.260

	/* Bool */
	var isExist bool = true

	fmt.Printf("apakah ada ? %t\n", isExist)

	/* String */
	var message string = "Bagaskoro"

	fmt.Printf("nama: %s\n", message)

	var greeting string = `Selamat datang di golang
	mari belajar bersama`

	fmt.Printf("greeting: %s\n", greeting)

	/* Constant */
	const PI int = 22 / 7

	fmt.Println("PI: ", PI)

	const (
		today    string = "senin" // manifest typing
		sekarang                  // manifest typing | value: senin
		isToday2 = true           // type interface
	)

	fmt.Println("Today: ", today, " sekarang: ", sekarang, " isToday: ", isToday2)

}
