package main

import (
	"fmt"

	"log"

	"example.com/greetings"
)

func main() {
	// without log
	// fmt.Println(greetings.Hello("bagas"))

	// with log
	log.SetPrefix("greetings: ")
    log.SetFlags(0)

	message,err := greetings.Hello("bagas");

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
	
}