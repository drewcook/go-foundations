package main

import "fmt"

func main() {
	var age int
	fmt.Printf("Please enter your age: ")
	fmt.Scan(&age) // use the '&' character to assign the input to a variable, followed by the variable name

	switch {
		case age < 10:
			fmt.Println("You are a child")
		case age < 19:
			fmt.Println("You are a teen")
		case age > 19:
			fmt.Println("You are an adult")
	}

	var username string
	fmt.Printf("Please enter your username: ")
	fmt.Scan(&username)

	// Just like other languages, there is a default case
	// You can also have the same result for multiple cases
	switch username {
		case "Drew":
		case "drew":
			fmt.Println("Your username is correct")
		case "Maira":
			fmt.Println("Your username is Maira")
		case "Rosky":
			fmt.Println("You're a dog")
		default:
			fmt.Println("Nice to meet you")
	}
}
