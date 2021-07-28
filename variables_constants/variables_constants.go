package main

import "fmt"

func main() {
	/*
		VARIABLES
		- can be changed after declaration
	*/

	// declare a variable using keyword "var"
	// Go is a stictly typed language, so need to define data type when declaring a variable
	var name string
	name = "Drew"
	fmt.Printf("Your name is %s\n", name)

	// shorthand - use colon/equals to declare a variable and infer its type
	first_name := "Drew"
	last_name := "Cook"
	fmt.Printf("Your name is %s %s\n", first_name, last_name)

	// variables should be camel case, not underscores
	firstName := "Drew"
	lastName := "Cook"
	fmt.Printf("Your name is %s %s\n", firstName, lastName)

	// you can declare variables on multiline
	var (
		a string
		b int
		c bool = false
	)
	a = "Rosky"
	b = 9
	fmt.Println(a, b, c)

	// declare and set multiple with type inferences on same line
	var d, e, f = "I am a string", 3.0, true
	// %v - default vslue
	// %._f - format as float with _ decimal places
	fmt.Printf("d: %v, e: %.2f, f: %v\n", d, e, f)

	/*
		CONSTANTS
		- cannot be changed after declaration
	*/

	// define using all caps, to differentiate it from variables throughout code
	const API_KEY = "aldkfnal.32902n.skns"
	fmt.Println(API_KEY)
	// API_KEY = "hello.world.123" - compile error because cannot reassign a constant
}
