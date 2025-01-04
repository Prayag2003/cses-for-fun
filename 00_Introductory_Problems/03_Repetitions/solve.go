package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	maxi := 1
	count := 1

	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count++
		} else {
			count = 1
		}

		maxi = max(maxi, count)
	}
	fmt.Println(maxi)
}
