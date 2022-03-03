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

	count := make([]int, n+1)
	var a int
	for i := 1; i < n; i++ {
		fmt.Fscan(in, &a)
		count[a]++
	}

	for i, v := range count {
		if i == 0 {
			continue
		}
		fmt.Println(v)
	}
}
