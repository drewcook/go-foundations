package main

import "fmt"

func main() {
	// if - else
	num := 10
	if num < 10 {
		fmt.Println("Less than 10")
	} else if num > 10 {
		fmt.Println("Greater than 10")
	} else if num == 10 {
		fmt.Println("Equal to 10")
	}

	// instead of multiple conditions on multiple lines, you can put them in one line
	num = 11
	if num <= 10 {
		fmt.Println("Less than or equal to 10")
	} else if num > 10{
		fmt.Println("Greater than 10")
	}

	// simplified with an OR statement, no elseif
	username := "Drew" // if
	username = "drew" // if
	username = "maira" // else
	if username == "Drew" || username == "drew" {
		fmt.Println("Username is correct")
	} else {
		fmt.Println("Wrong username")
	}

	// using the AND condtion
	username = "drew"
	age := 21
	if username == "drew" && age == 21 {
		fmt.Println("Username is correct")
	} else {
		fmt.Println("Wrong username")
	}
}
