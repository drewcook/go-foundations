package main

import "fmt"

func main() {
	fmt.Println("1")
	fmt.Println("2")
	fmt.Println("3")

	// for loops - there is only one type of loop in Go, the for loop
	// no while or do while loops
	fmt.Println("Iterating up:")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// we can declare a variable prior to the loop and use it within
	// omit the first iterator declaration, but use the semicolon delimiter
	j := 11
	fmt.Println("Iterating down:")
	for ; j >= 0; j-- {
		fmt.Printf("%d ", j) // use %d for digits
	}
	fmt.Println()

	// Alternatively, we can iterate within the fn
	k := 10
	fmt.Println("Alternate:")
	for k <= 20 {
		fmt.Printf("%d ", k)
		k += 2
	}
	fmt.Println()

	// We could also use a for loop for any variable with break statments and iteratiion
	// This could simulate a while loop in JS
	a := 5
	fmt.Println("Fives:")
	for {
		if (a > 100) {
			break;
		}
		fmt.Printf("%d ", a)
		a += 5
	}
	fmt.Println()

	// using the 'continue' keyword to break out and skip to the next iteration
	fmt.Println("Odds:")
	for b := 1; b <= 20; b++ {
		// omit even numbers
		if b % 2 == 0 {
			continue
		}
		fmt.Printf("%d ", b)
	}
}
