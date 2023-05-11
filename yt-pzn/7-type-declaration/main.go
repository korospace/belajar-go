package main

import "fmt"

func main() {
	type marriage bool

	var married marriage = true
	var notMarried marriage = false

	fmt.Println("menikah ==", married)
	fmt.Println("belum menikah ==", notMarried)
}
