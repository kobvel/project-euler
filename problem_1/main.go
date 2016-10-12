package main

import "fmt"

func main() {

	fmt.Println(calculate(23))
	fmt.Println(calculate(1000))
}

// could be improved for unlimited arguments, 3,5 ...
func calculate(max int) int {
	var sum int
	for i := 0; i < max; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}
