package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &b[i])
	}

	sort.Ints(a)
	sort.Ints(b)

	ans := 0
	for i, a_asc := range a {
		ans += distance(a_asc, b[i])
	}
	fmt.Println(ans)
}

func distance(u, v int) int {
	if u-v < 0 {
		return -(u - v)
	}
	return u - v
}
