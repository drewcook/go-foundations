package main

import (
	"fmt"
	"math"
)

func main() {
	// PART 1:
	// Print out the data types along with the values
	var1 := math.Pi
	var2 := "Hey Gopher!"
	var3 := true
	var4 := 49
	// Print out should be in the following format - e.g. var5 := "I am learning Go"
	// Output ->
	// Value: I am learning Go    Data Type: string
	fmt.Printf("Value: %.5f\tData Type: %T\n", var1, var1) // float - %.5f
	fmt.Printf("Value: %s\tData Type: %T\n", var2, var2) // string - %s
	fmt.Printf("Value: %v\tData Type: %T\n", var3, var3) // general - %v
	fmt.Printf("Value: %v\tData Type: %T\n", var4, var4) // general - %v

	// PART 2:
	var5 := false
	// Guess the output and print out the data types
	if var5 == false {
		// Print out after each variable declaration
		var1 = 35 + 4673 / 127
		fmt.Printf("Value: %v\tData Type: %T\n", var1, var1)
		var2 = "Yo Gopher!"
		fmt.Printf("Value: %v\tData Type: %T\n", var2, var2)
		var3 = false
		fmt.Printf("Value: %v\tData Type: %T\n", var3, var3)
		var4 = 3675 - 982 + 1293
		fmt.Printf("Value: %v\tData Type: %T\n", var4, var4)
		var5 = true
	}
}
