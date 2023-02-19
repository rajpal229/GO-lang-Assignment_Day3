package main

import "fmt"

func fibonacci_r(n int) int {
	if n == 1 {
		return 0
	}
	if n <= 3 {
		return 1
	} else {
		return fibonacci_r(n-1) + fibonacci_r(n-2)
	}
}

func main() {
	var n int
	fmt.Print("Enter the size of the Fibonacci series: ")
	fmt.Scan(&n)
	fmt.Printf("Fibonacci Series of size %d: ", n)
	for i := 0; i < n; i++ {
		fmt.Print(fibonacci_r(i + 1))
		fmt.Print(" ")
	}
}
