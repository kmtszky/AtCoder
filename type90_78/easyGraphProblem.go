package main

import (
	"fmt"
)

func main() {
	var n, m, a, b int
	fmt.Scan(&n, &m)

	counts := make([]int, n+1)
	for i := 0; i < m; i++ {
		fmt.Scan(&a, &b)
		if a > b {
			counts[a]++
		} else if b > a {
			counts[b]++
		}
	}

	ans := 0
	for _, v := range counts {
		if v == 1 {
			ans++
		}
	}
	fmt.Println(ans)
}
