// A simple web application example.
// Taken from: https://golang.org/doc/articles/wiki/
// https://ianmcloughlin.github.io
// https://gist.github.com/ianmcloughlin/c4c2b8dc586d06943f54b75d9e2250fe

package main

import (
	"fmt"
	"math/rand" // imports

	"time"
)

func main() {

	rand.Seed(time.Now().UTC().UnixNano())
	answers := []string{ // randon answers
		"I’m not sure what you’re trying to say. Could you explain it to me?",
		"How does that make you feel?",
		"Why do you say that?",
	}

	output := "Eliza: " + answers[rand.Intn(len(answers))]
	input := "User: Hi "
	fmt.Println(input)
	fmt.Println(output)
} // main
