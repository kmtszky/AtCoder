package main

import (
	"fmt"
)

func main() {
	var n, q, c, p, l, r int

	fmt.Scan(&n)
	c1sum := make([]int, n+1)
	c2sum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&c, &p)
		if c == 1 {
			c1sum[i] = c1sum[i-1] + p
			c2sum[i] = c2sum[i-1]
		} else {
			c1sum[i] = c1sum[i-1]
			c2sum[i] = c2sum[i-1] + p
		}
	}

	fmt.Scan(&q)
	for i := 0; i < q; i++ {
		fmt.Scan(&l, &r)
		c1 := c1sum[r] - c1sum[l-1]
		c2 := c2sum[r] - c2sum[l-1]
		fmt.Println(c1, c2)
	}
}
