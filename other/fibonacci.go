package other

// Fibonacci function generates a Fibonacci sequence up to the specified length n.
// It takes an integer n as input and returns a slice of integers.
// if n = 10, it will return [1 1 2 3 5 8 13 21 34 55]
func Fibonacci(n int) []int {
	// Declare a slice to store the Fibonacci sequence.
	fib := make([]int, n)

	// Initialize the first two numbers in the sequence.
	a := 1 // first number deret fibonacci
	b := 1

	// Loop through the sequence and calculate each number.
	for i := 0; i < n; i++ {
		// Assign the current number in the sequence to the slice.
		fib[i] = a

		// Calculate the next two numbers in the sequence.
		// cara 1:
		// a, b = b, a+b

		// cara 2:
		c := a
		a = b
		b = c + b
	}

	// Return the Fibonacci sequence.
	return fib
}
