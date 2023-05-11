package main

import "fmt"

func main() {
	var cars = [4]string{"mazda", "honda", "bmw", "audi"} // array

	/*==========================
	Slice Is Reference Data Type
	==========================*/
	var newCars = cars[0:2] // slice

	newCars[0] = "mazda xx" // will change cars index-0

	fmt.Println(cars[0:2]) // "mazda xx", "honda"

	/*==========================
	Slice Operation
	==========================*/
	fmt.Println(cars[0:2]) // mazda xx, honda
	fmt.Println(cars[0:4]) // mazda xx, honda, bmw, audi
	fmt.Println(cars[:])   // mazda xx, honda, bmw, audi
	fmt.Println(cars[2:])  // bmw, audi
	fmt.Println(cars[:2])  // mazda xx, honda

	/*==========================
	Len & Cap
	==========================*/
	var brand1 = cars[0:2]   // mazda xx, honda
	fmt.Println(len(brand1)) // 2
	fmt.Println(cap(brand1)) // 4

	var brand2 = cars[2:4]   // bmw, audi
	fmt.Println(len(brand2)) // 2
	fmt.Println(cap(brand2)) // 2

	/*==========================
	Append
	==========================*/
	var brand3 = append(brand1, "bmw xx")

	fmt.Println(cars)   // mazda xx, honda, bmw xx, audi
	fmt.Println(brand1) // mazda xx, honda
	fmt.Println(brand3) // mazda xx, honda, bmw xx

	var carsxx []string

	var carsyy = append(carsxx, "he")

	var carszz = append(carsyy, "hehe")

	fmt.Println(carsxx)
	fmt.Println(carsyy)
	fmt.Println(carszz)

	/*==========================
	Copy
	==========================*/
	var dst = make([]string, 3)
	var src = []string{"bmw", "mazda", "honda", "nissan"}
	var n int = copy(dst, src)

	fmt.Println("dst: ", dst) // "bmw", "mazda", "honda"
	fmt.Println("src: ", src)
	fmt.Println("n: ", n)

	var src2 = []string{"honda xx", "nissan xx"}
	copy(dst, src2)

	fmt.Println("dst: ", dst) // "honda xx", "nissan xx", "honda"

	/*==========================
	3 Index Slicing
	==========================*/
	fruits := []string{"apple", "manggo", "banana"}
	fruitsA := fruits[0:2]
	fruitsB := fruits[0:2:2]

	fmt.Println("fruitsA:", fruitsA) // "apple", "manggo" | len:2 | cap:3
	fmt.Println("fruitsB:", fruitsB) // "apple", "manggo" | len:2 | cap:2

}
