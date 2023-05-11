package main

import "fmt"

func main() {

	// number to number
	var data16 int16 = 128
	var data32 int32 = int32(data16)
	var data8 int8 = int8(data32)

	fmt.Println(data32)
	fmt.Println(data8)

	// number to string
	var name string = "bagas"

	var getWord0 byte = name[0]
	var getWord1 byte = name[1]

	var word0 string = string(getWord0)
	var word1 string = string(getWord1)

	fmt.Println(word0)
	fmt.Println(word1)

}
