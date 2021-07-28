// ask first name
// ask last name
// Your name is: Drew Cook
// Hint - read printF in fmt package
// use fmt.Scan() or fmt.Scanln for reading user input

package main

import "fmt"

func main() {
	var firstName string = ""
	var lastName string = ""
	// first name
	fmt.Println("What is your first name?")
	fmt.Scanln(&firstName)
	// last name
	fmt.Println("What is your last name?")
	fmt.Scanln(&lastName)
	// read out
	fmt.Printf("Your name is: %s %s", firstName, lastName)
}
