package greetings

import "fmt"

func Hello(name string) string {
    // -- way 1
	// var message string;
	// message = fmt.Sprintf("Hi, %v. Welcome!", name)

	// -- way 2
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
    
	return message
}