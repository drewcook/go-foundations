/*
	DATA TYPES
	- string
	- int
	- bool
	- float64 - usually default for most floats
*/
package main

import "fmt"

func main() {
	const PI float64 = 3.145
	name := "John"
	age := 29
	weight := 174.5
	adult := true

	fmt.Println(PI)

	// using %T displays the type of the variable
	// using \n is newline; using \t is a tab
	fmt.Printf("name %T\t", name)
	fmt.Printf("age %T\n", age)
	fmt.Printf("weight %T\n", weight)
	fmt.Printf("adult %T\n", adult)

	/*
		OTHER DATA TYPES - less commonly used
		- int8, int16, int32, int64 - width of the integer is X bits
		- uint8, uint16, uint32, uint64, uintptr - unsigned integers at X bits, can only hold positive integers
		- byte - alias for uint8
		- rune - alias for int32 - represents a Unicode code print
		- float32 - working with 32 bit floating numbers
		- complex64, complex32 - working with more complex numbers
	*/
}
