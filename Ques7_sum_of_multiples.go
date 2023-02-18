package main

import "fmt"

func sum_multiples(num int) int {
	var sum int
	for i := 1; i < num; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum = sum + i
		}
	}
	return sum
}

func main() {
	var n int
	fmt.Print("Enter the number: ")
	fmt.Scan(&n)
	fmt.Printf("the sum of all the multiples of 3 or 5 below number %d is %d", n, sum_multiples(n))
}
