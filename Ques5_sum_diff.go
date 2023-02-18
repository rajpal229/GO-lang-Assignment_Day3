package main

import "fmt"

func sum(x float64, y float64) float64 {
	return x + y
}
func diff(x float64, y float64) float64 {
	if x > y {
		return x - y
	}
	return y - x
}
func getSum_n_Diff(x float64, y float64) (float64, float64) {
	return sum(x, y), diff(x, y)
}

func main() {
	var x, y, op float64
	fmt.Print("Enter the Number\nx: ")
	fmt.Scan(&x)
	fmt.Print("y: ")
	fmt.Scan(&y)
	fmt.Print("Enter the function to perform\n1. Sum\n2. Difference\n3. Both\n")
	fmt.Scan(&op)
	switch op {
	case 1:
		fmt.Printf("The sum of %v and %v is %v", x, y, sum(x, y))
	case 2:
		fmt.Printf("The difference of %v and %v is %v", x, y, diff(x, y))
	case 3:
		var (
			s, d float64 = getSum_n_Diff(x, y)
		)
		fmt.Printf("The sum of %v and %v is %v\nThe difference of %v and %v is %v", x, y, s, x, y, d)
	}
}
