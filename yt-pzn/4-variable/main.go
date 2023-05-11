package main

import "fmt"

func main() {
	// way 1
	// =====
	var name string

	name = "bagas"
	fmt.Println(name)

	name = "bagas koro"
	fmt.Println(name)

	// way 2
	// =====
	var age = 10 // int
	fmt.Println(age)

	// way 3
	// =====
	country := "indonesia"
	fmt.Println(country)

	country = "USA"
	fmt.Println(country)

	// Multiple Declaration
	// ====================
	var (
		firstName = "bagas"
		lastName  = "koro"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

}
