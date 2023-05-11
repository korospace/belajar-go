package main

import (
	"fmt"
	"example.com/greetings" // internal
	"github.com/dustinkirkland/golang-petname" // external 
)

func main() {
	message := greetings.Hello("korospace")

	fmt.Println(message)

	fmt.Println(petname.Name())
}