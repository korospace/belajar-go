package main

import "fmt"

func main() {
	/*
	* Integer
	*==================
	 */
	var data8 int8 = 127                   // -128
	var data16 int16 = 32767               // -32768
	var data32 int32 = 2147483647          // -2147483648
	var data64 int64 = 9223372036854775807 // -9223372036854775808

	fmt.Println(data8)
	fmt.Println(data16)
	fmt.Println(data32)
	fmt.Println(data64)

	/*
	* Unsign integer
	*=========================
	 */
	var unsigndata8 uint8 = 255                    // 0
	var unsigndata16 uint16 = 65535                // 0
	var unsigndata32 uint32 = 4294967295           // 0
	var unsigndata64 uint64 = 18446744073709551615 // 0

	fmt.Println(unsigndata8)
	fmt.Println(unsigndata16)
	fmt.Println(unsigndata32)
	fmt.Println(unsigndata64)

	/*
	* Floating
	*==================
	 */
	var dataFloat32 float32 = 2.2 // min: 1.8*10 -38 max: 3.4*10 38
	var dataFloat64 float64 = 2.2 // min: 2.23*10 -308 max: 1.80*10 308

	fmt.Println(dataFloat32)
	fmt.Println(dataFloat64)

	/*
	* Alias
	*==================
	 */
	var aliasUint8 byte = 255
	var aliasint32 rune = 2147483647
	var aliasMinInt32 int = 1   // minimal int32
	var aliasMinUint32 uint = 1 // minimal uint32

	fmt.Println(aliasUint8)
	fmt.Println(aliasint32)
	fmt.Println(aliasMinInt32)
	fmt.Println(aliasMinUint32)

}
