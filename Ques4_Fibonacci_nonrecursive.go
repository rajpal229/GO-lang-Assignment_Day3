package main

import "fmt"

func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		a := 0
		b := 1
		for i := 2; i <= n; i++ {
			c := a + b
			a = b
			b = c
		}
		return b
	}
}
func main() {
	var n, count int
	fmt.Print("Enter the size of the Fibonacci series: ")
	fmt.Scan(&n)
	fmt.Println("The Fibonacci series of size", n, ":")
	for i := 0; i < n; i++ {
		fmt.Print(fibonacci(i))
		fmt.Print(" ")
		if i != 0 {
			if fibonacci(i)%2 == 0 {
				count = count + 1
			}
		}
	}
	fmt.Printf("\n\nthe count of even values of a Fibonacci series of size %d is %d", n, count)
}
