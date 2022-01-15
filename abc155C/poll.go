package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var n int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)

	var s string
	smap := make(map[string]int)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &s)

		if _, ok := smap[s]; ok {
			smap[s]++
		} else {
			smap[s] = 1
		}
	}

	max := 0
	for _, v := range smap {
		if max < v {
			max = v
		}
	}

	answers := make([]string, 0)
	for s, v := range smap {
		if v == max {
			answers = append(answers, s)
		}
	}

	sort.Strings(answers)
	for _, ans := range answers {
		fmt.Println(ans)
	}
}
