package main

import (
	"fmt"

	petname "github.com/dustinkirkland/golang-petname"
)

func main() {
	fmt.Println("Hello", "world!", "how", "are", "you")

	/* Generate Random Path Name
	use package */
	fmt.Println(petname.Name())
}
