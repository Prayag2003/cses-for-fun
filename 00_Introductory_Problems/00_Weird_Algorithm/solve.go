package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for n != 1 {
		fmt.Print(fmt.Sprintf("%d ", n))
		if n%2 == 1 {
			n = 3*n + 1
		} else {
			n = n / 2
		}
	}
	fmt.Print(1)
}
