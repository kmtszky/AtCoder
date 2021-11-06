package main

import (
	"fmt"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&b[i])
	}

	abs := 0
	for i := 0; i < n; i++ {
		abs += calc(b[i] - a[i])
	}

	if (abs+k)%2 == 0 && abs <= k {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func calc(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
