package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var sum int = 0
	prod := n * (n + 1) / 2

	var list = make([]int, n-1)

	for i := 0; i < n-1; i++ {
		fmt.Scan(&list[i])
		sum += list[i]
	}

	fmt.Print(prod - sum)
}
