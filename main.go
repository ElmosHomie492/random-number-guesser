package main

import (
	"fmt"
	"math/rand"
	"time"
)

const GrantsMagicalNumber int = 123
const MIN int = 1
const MAX int = 1000

func generateRandomNumber() int {
	seedVal := rand.NewSource(time.Now().UnixNano())
	randGen := rand.New(seedVal)
	return randGen.Intn(MAX-MIN) + MIN
}

func checkGuessAgainstAnswer(guess int) bool {
	return guess == GrantsMagicalNumber
}

func main() {
	var guess int

	for exit := false; !exit; exit = checkGuessAgainstAnswer(guess) {
		guess = generateRandomNumber()
	}

	fmt.Println("Kaleb's first guess is definitely: ", guess)
}
