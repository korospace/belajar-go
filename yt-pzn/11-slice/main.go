package main

import "fmt"

func main() {
	var months = [...]string{
		"januari",
		"februari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}

	/**
	Create Slice From Existing Array
	================================
	*/
	var slice1 = months[4:8] // [mei juni juli agustus]

	fmt.Println(slice1)
	fmt.Println(len(slice1)) // length   = 4
	fmt.Println(cap(slice1)) // capacity = 8 --> from index 4 to last index in array

	slice1[0] = "meixx"
	fmt.Println(months) // [... ... ... meixx ... ... ... ... ... ... ... ...]

	var slice2 = months[8:]
	fmt.Println(slice2) // index 8 to last index 		 --> [september oktober november desember]

	var slice3 = months[:8]
	fmt.Println(slice3) // first index to before index 8 --> [januari februari maret april meixx juni juli agustus]

	var slice4 = months[:]
	fmt.Println(slice4) // first index to last index 	 --> [januari februari maret april meixx juni juli agustus september oktober november desember]

	/**
	Append Slice
	============
	*/
	var slice5 = months[9:11] // [oktober november]

	var append1 = append(slice5, "bulanBaru1")
	fmt.Println(append1)      // [oktober november bulanBaru1]
	fmt.Println(len(append1)) // 3
	fmt.Println(cap(append1)) // 3

	append1[0] = "oktoberxx"
	fmt.Println(months) // [... ... ... ... meixx ... ... ... ... oktoberxx ... bulanBaru1]

	var days = [...]string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"}
	var slice6 = days[5:]                     // [saturday sunday]
	var append2 = append(slice6, "hariBaru1") // will create new array because capacity is full

	append2[0] = "saturdayxx"
	fmt.Println(append2) // [saturdayxx sunday hariBaru1]
	fmt.Println(days)    // [monday tuesday wednesday thursday friday saturday sunday]

	/**
	Create Slice From Scratch
	=========================
	*/
	var slice7 = make([]string, 2, 4) // type,length,capacity

	slice7[0] = "toyota-supra"
	slice7[1] = "nissan-gtr"

	fmt.Println(slice7)      //[toyota-supra nissan-gtr]
	fmt.Println(len(slice7)) // 2
	fmt.Println(cap(slice7)) // 4

	/**
	Copy Slice
	=========================
	*/
	slice8 := make([]string, len(slice7), cap(slice7))

	copy(slice8, slice7)

	fmt.Println(slice8) // [toyota-supra nissan-gtr]

}
