package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var userInput int
	var secretNumber int
	var guessed bool = false

	rand.Seed(time.Now().Unix())
	secretNumber = rand.Intn(10)

	for guessed != true {
		fmt.Println("Please enter a number:")
		fmt.Scan(&userInput)

		if userInput == secretNumber {
			guessed = true
			fmt.Println("You guessed the number. You won!")
		} else if userInput < secretNumber {
			fmt.Println("To low... Guess again.")
		} else if userInput > secretNumber {
			fmt.Println("To high... Guess again.")
		}
	}
}
