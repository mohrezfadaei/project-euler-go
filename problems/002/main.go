package main

import "fmt"

func main() {
	const limit = 4_000_000
	sum, a, b := 0, 1, 2 // a = F(1), b = F(2)

	for b <= limit {
		if b%2 == 0 {
			sum += b
		}
		a, b = b, a+b
	}

	fmt.Println(sum)
}
