package main

import "fmt"

func main() {
	/* ====================
	1 - only with condition
	==================== */
	var iLoop1 = 0

	for iLoop1 < 3 {
		fmt.Println("only with condition | Angka", iLoop1)
		iLoop1++
	}

	/* ==============
	2 - only with FOR
	============== */
	var iLoop2 = 0

	for {
		fmt.Println("only with FOR | Angka", iLoop2)

		iLoop2++
		if iLoop2 == 4 {
			break
		}
	}

	/* ===================
	3 - Continue And Break
	=================== */
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}

		if i > 6 {
			break
		}

		fmt.Println("Continue And Break | Angka", i)
	}

	/* =====================
	4 - Label in nested loop
	===================== */
outerLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 2 {
				break outerLoop
			}
			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
	}

	/* ===========
	For - Range
	=========== */
	var fruits = [4]string{"apple", "banana", "manggo", "watermelon"}

	for i, fruit := range fruits {
		fmt.Printf("elemen %d : %s\n", i, fruit)
	}

	// underscore
	for _, fruit := range fruits {
		fmt.Printf("elemen : %s\n", fruit)
	}
}
