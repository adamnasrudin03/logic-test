package other

import (
	"fmt"
)

func PyramidPattern(n int) {
	str := "-----------------"
	pyramidPatternA(n)
	fmt.Println(str)
	pyramidPatternB(n)
	fmt.Println(str)
	pyramidPatternC(n)
}

// *
// * *
// * * *
func pyramidPatternA(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

// * * *
// * *
// *
func pyramidPatternB(n int) {

	for i := n; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

// *
// * *
// * * *
func pyramidPatternC(n int) {

	// Outer loop to iterate through rows
	for i := 0; i < n; i++ {

		// Inner loop to print spaces
		// The number of spaces decreases by 1 in each row
		for j := 0; j < n-i-1; j++ {
			fmt.Print(" ") // Print a single space
		}

		// Inner loop to print stars
		// The number of stars increases by 1 in each row
		for j := 0; j <= i; j++ {
			fmt.Print("* ") // Print a star followed by a space
		}

		// Print a newline character to move to the next line
		fmt.Println()
	}

}
