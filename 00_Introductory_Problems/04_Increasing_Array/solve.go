package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	ans := 0
	for i := 1; i < n; i++ {
		if arr[i] < arr[i-1] {
			ans += arr[i-1] - arr[i]
			arr[i] = arr[i-1]
		}
	}

	fmt.Println(ans)
}
