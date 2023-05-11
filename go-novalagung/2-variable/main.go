package main

import "fmt"

func main() {
	// manifest typing
	var one, two string = "one", "two"

	fmt.Println(one, two)

	// interface typing
	three, four := 3, "four"

	fmt.Println(three, four)

	/* reserved variable
	====================
	semua variabel yang dideklarasikan harus digunakan
	kecuali variable underscore
	*/
	five, _ := "five", 6

	fmt.Println(five)

}
