// CodeSourced from:
// https://golang.org/doc/articles/wiki/
// https://ianmcloughlin.github.io
// https://gist.github.com/ianmcloughlin/c4c2b8dc586d06943f54b75d9e2250fe
//https://www.youtube.com/watch?v=X8aLGge3GIM&t=990s
// https://gist.github.com/ianmcloughlin/c4c2b8dc586d06943f54b75d9e2250fe
//https://golang.org/pkg/reflect/

//Student Name: Kevin Moran
//Student Number: G00306474

package main

import (
	"fmt"
	"math/rand" // imports
	"regexp"    //used
	"strings"
	"time"
)

//Array of user input
var userInput = []string{
	"People say I look like both my mother and father.",
	"Father was a teacher.",
	"I was my father’s favourite.",
	"I am looking forward to the weekend.",
	"My grandfather was French!",
	"I am happy.",
	"I am not happy with your responses.",
	"I am not sure that you understand the effect that your questions are having on me.",
	"I am supposed to just take what you’re saying at face value?",
	"Im so tired",
	"Should I quit college?",
	"Should I quit college?",
	"Goodnight Computer",
} //userInput End

func ElizaResponse(input string) string {
	if matched, _ := regexp.MatchString(`(?i).*\bfather\b.*`, input); matched {
		return "Why don’t you tell me more about your father?" //checks for the word "father" and gives this response
	}
	if matched, _ := regexp.MatchString(`(?i).*\bgoodnight\b.*`, input); matched {
		return "Goodbye Kevin" //checks for the word "father" and gives this response
	}
	if matched, _ := regexp.MatchString(`(?i).*\bcollege\b.*`, input); matched {
		return "Your three years in might as well finish it?" //checks for the word "father" and gives this response
	}
	if matched, _ := regexp.MatchString(`(?i).*\btired\b.*`, input); matched {
		return "You should get some rest." //checks for the word "finished" and replys if found
	}

	//https: //regex101.com/r/xE2vT0/1 (rerex tester)
	re := regexp.MustCompile("(?i)" + `(?i)i\'?(?:\s?am|m)([^.?!]*)[.?!]?`)
	if matched := re.MatchString(input); matched {
		return re.ReplaceAllString(input, "How do you know you are $1?")
		// if "I am, I'm & Im" and gives this response
	}

	answers := []string{ // randon answers given if there are no matches
		"I’m not sure what you’re trying to say. Could you explain it to me?",
		"How does that make you feel?",
		"Why do you say that?",
	}
	rand.Seed(time.Now().UTC().UnixNano())

	return answers[rand.Intn(len(answers))]

} // elizaResponse end
func Reflect(input string) string {

	// Split the input on word boundaries.
	boundaries := regexp.MustCompile(`\b`)
	tokens := boundaries.Split(input, -1)

	// List the reflections.
	reflections := [][]string{
		{`I`, `you`},
		{`me`, `you`},
		{`you`, `me`},
		{`my`, `your`},
		{`your`, `my`},
	}

	// Loop through each token, reflecting it if there's a match.
	for i, token := range tokens {
		for _, reflection := range reflections {
			if matched, _ := regexp.MatchString(reflection[0], token); matched {
				tokens[i] = reflection[1]
				break
			}
		}
	}

	// Put the tokens back together.
	return strings.Join(tokens, " ")
} // reflect end

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	question := userInput[rand.Intn(len(userInput))]

	fmt.Println("Question:" + question)
	fmt.Println("Answer:" + ElizaResponse(question))

} // main  end
