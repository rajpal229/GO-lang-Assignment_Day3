package main

import "fmt"

func factorial_nonr(num int) int {
	//Ques1
	var output int
	if num == 0 {
		return 1
	}
	output = 1
	for i := 2; i < num+1; i++ {
		output = output * i
	}
	return output
}

func main() {
	var n int
	fmt.Print("Enter the Number: ")
	fmt.Scan(&n)
	fmt.Printf("Factorial of %d is %d", n, factorial_nonr(n))
}
