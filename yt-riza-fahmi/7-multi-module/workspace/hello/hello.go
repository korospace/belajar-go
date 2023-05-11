package main

import (
    "fmt"

	"log"

    // "golang.org/x/example/stringutil"

	"example.com/greetings"
)

func main() {
    // fmt.Println(stringutil.Reverse("Hello")) // olleH
    // fmt.Println(stringutil.ToUpper("Hello")) // HELLO

	names := []string{"johan","ridho","jamal"}

	message,err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message) 
}