package main

import "fmt"

func main() {
	/* ========
	Declaration
	======== */
	var fruits [4]string

	// Horizontal declaration
	fruits = [4]string{"apple", "banana", "manggo", "watermelon"}

	// Vertical declaration
	fruits = [4]string{
		"apple",
		"banana",
		"manggo",
	}

	fmt.Println("Fruits length:", len(fruits))
	fmt.Println("Fruits values:", fruits)
	fmt.Println("Fruits index-3:", fruits[3])

	/* =========================
	Declaration with make kyword
	========================= */
	var cars = make([]string, 2)
	cars[0] = "mazda"

	fmt.Println("Car length:", len(cars))
	fmt.Println("Car values:", cars)
	fmt.Println("Car index-1:", cars[1])

	/* =======================
	Declaration without length
	======================= */
	var foods = [...]string{"tes1", "tes2"}

	foods[0] = "pizza"
	foods[1] = "burger"

	fmt.Println("Food values:", foods)

	/* ===========
	Multidimension
	=========== */
	var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}

	fmt.Println("numbers1", numbers1)
	fmt.Println("numbers2", numbers2)
}
