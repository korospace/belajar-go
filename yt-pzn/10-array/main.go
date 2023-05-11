package main

import "fmt"

func main() {
	// way 1
	var names [3]string

	names[0] = "jhon"
	names[1] = "salka"
	names[2] = "tian"
	fmt.Println(names)

	// way 2
	var cars = [...]string{
		"nissan gtr",
		"toyota supra",
		"mazda rx7",
	}
	fmt.Println(cars)

	// count array length
	fmt.Println("total cars:", len(cars))

	var countries [12]string

	fmt.Println("total countries:", len(countries))

}
