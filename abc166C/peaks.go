package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscan(in, &n, &m)

	h := make([]int, n)
	c := make([]bool, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &h[i])
		c[i] = true
	}

	var a, b int
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &a, &b)
		a--
		b--
		if h[a] <= h[b] {
			c[a] = false
		}
		if h[b] <= h[a] {
			c[b] = false
		}
	}

	ans := 0
	for i := 0; i < n; i++ {
		if c[i] {
			ans++
		}
	}
	fmt.Println(ans)
}
