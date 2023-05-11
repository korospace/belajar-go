package greetings

import (
	"fmt"

	"errors"

	"math/rand"

	"time"

	"testing"

	"regexp"
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

func randomHello() string {
	hellos := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return hellos[rand.Intn(len(hellos))]
}

func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b`+name+`\b`)

	msg,err := Hello("")

	if !want.MatchString(msg) || err != nil {
        t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
    }
}

func TestHelloEmpty(t *testing.T) {
    msg, err := Hello("")
    if msg != "" || err == nil {
        t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
    }
}