package main

import "fmt"

func main() {
	var chicken = map[string]int{}

	chicken["satu"] = 10
	chicken["dua"] = 20

	fmt.Println(chicken)         // [dua:20 satu:10]
	fmt.Println(chicken["satu"]) // 10

	/*================
	Initialization Map
	================*/
	// var months = map[string]int       // ERROR
	// var months = make(map[string]int) // No Error
	// var months = map[string]int{
	// 	"january":  10,
	// 	"february": 20,
	// }

	var months = map[string]int{"january": 10, "february": 20}

	fmt.Println(months)

	for key, value := range months {
		fmt.Println(key, ":", value)
	}

	/*================
	Delete Map
	================*/
	delete(months, "january")

	fmt.Println(months)

	/*================
	Is Exist
	================*/
	value, isExist := months["january"]
	value2, isExist2 := months["february"]

	fmt.Println(value, ":", isExist)
	fmt.Println(value2, ":", isExist2)

	/*================
	Slice & Map
	================*/
	students := []map[string]string{
		map[string]string{"name": "bagas", "class": "one"},
		map[string]string{"name": "john", "class": "two"},
		map[string]string{"name": "aji", "class": "three"},
	}

	for _, student := range students {
		fmt.Println("name: ", student["name"], " class: ", student["class"])
	}

	students = []map[string]string{
		{"name": "bagas", "class": "one"},
		{"name": "john", "class": "two"},
		{"name": "aji", "class": "three"},
	}

	for _, student := range students {
		fmt.Println("name: ", student["name"], " class: ", student["class"])
	}

}
