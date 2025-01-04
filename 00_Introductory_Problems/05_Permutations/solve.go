package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	if n == 1 {
		fmt.Println(1)
	} else if n == 2 || n == 3 {
		fmt.Println("NO SOLUTION")
	} else {
		for i := 2; i <= n; i += 2 {
			fmt.Printf("%d ", i)
		}
		for i := 1; i <= n; i += 2 {
			fmt.Printf("%d ", i)
		}
		fmt.Println()
	}
}
