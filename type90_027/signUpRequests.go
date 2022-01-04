package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(in, &n)

	ans := make([]int, 0)
	entered := make(map[string]bool)

	var s string
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &s)

		if entered[s] {
			continue
		}
		entered[s] = true
		ans = append(ans, i)
	}

	for _, v := range ans {
		fmt.Println(v)
	}
}
