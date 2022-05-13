package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(in, &n)

	a := make([]int, n+2)
	for i := 1; i < n+1; i++ {
		fmt.Fscan(in, &a[i])
	}
	inf := int(1e18)
	a[0] = -inf
	a[n+1] = inf
	sort.Ints(a)

	var q int
	fmt.Fscan(in, &q)

	var b int
	for i := 0; i < q; i++ {
		fmt.Fscan(in, &b)
		v := sort.SearchInts(a, b)
		ans := min(abs(a[v]-b), abs(a[v-1]-b))
		fmt.Println(ans)
	}
}

func abs(v int) int {
	if v > 0 {
		return v
	}
	return -v
}

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}
