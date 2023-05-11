package main

import (
	"fmt"

	"log"

	"example.com/greetings"
)


func main() {
	// fmt.Println(greetings.Hello(""))

	names := []string{"johan","ridho","jamal"}

	message,err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message) 	      // => map[jamal:Hail, jamal! Well met! johan:Great to see you, johan! ridho:Hail, ridho! Well met!]
	// fmt.Println(message["jamal"]) // => Hail, jamal! Well met!
}