package main

import "fmt"

func count_non_prime(num int) int {
	var count int
	if num > 0 {
		for i := 1; i < num; i++ {
			var flag bool
			if i == 1 || i == 2 {
				flag = true
			}
			for j := 2; j <= i/2; j++ {
				if i%j == 0 || i == 1 || i == 2 {
					flag = true
					break
				}
			}
			if flag == true {
				if i%2 != 0 {
					count = count + 1
				}
			}
		}
	}
	return count
}

func main() {
	var n int
	fmt.Print("Enter the number: ")
	fmt.Scan(&n)
	fmt.Printf("count of non-prime odd numbers below %d is %d", n, count_non_prime(n))
}
