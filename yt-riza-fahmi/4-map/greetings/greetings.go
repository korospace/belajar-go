package greetings

import (
	"fmt"

	"errors"

	"math/rand"

	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("name is empty")
	}

	message := fmt.Sprintf(randomHello(), name)

	return message, nil

}

// func Hellos(names []string) (map[int]string,error) {
func Hellos(names []string) (map[string]string, error) {
	// messages := make(map[int]string)
	messages := make(map[string]string)

	// for i,name := range names {
	for _, name := range names {
		message, err := Hello(name)

		if err != nil {
			return nil, err
		}

		// messages[i] = message
		messages[name] = message
	}

	return messages, nil
}

func randomHello() string {
	hellos := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return hellos[rand.Intn(len(hellos))]
}
