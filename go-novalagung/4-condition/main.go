package main

import "fmt"

func main() {
	var point float32 = 10000.0

	/* if else temporary */
	if percent := point / 100; percent >= 80 {
		fmt.Printf("%.1f%s perfect!\n\n", percent, "%")
	} else if percent < 80 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}

	/* swith case with { } */
	switch point {
	case 10000.0:
		{
			fmt.Println("aweosome")
			fmt.Println("your point is", point)
			fmt.Println("")
		}
	case 10001, 9999:
		fmt.Println("good")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you can be better!")
		}
	}

	/* swith case with IF ELSE */
	switch {
	case point == 10001.0:
		{
			fmt.Println("aweosome")
			fmt.Println("your point is", point)
		}
	case (point < 10001) && (point > 9999):
		fmt.Println("awesome :p")
		fmt.Println("")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

	/* swith case with FALLTHROUGH */
	switch {
	case point == 10000.0:
		fmt.Println("CASE 1")
		fallthrough
	case (point < 10001) && (point > 9999):
		fmt.Println("CASE 2")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}
}
