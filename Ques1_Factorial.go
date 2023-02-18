package main

import "fmt"

func factorial_r(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial_r(num-1)
}

func main() {
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)
	fmt.Printf("Factorial of %d is %d\n", n, factorial_r(n))
}
