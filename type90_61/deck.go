package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var q int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &q)

	var t, x int
	top, btm := 1, 2
	list := make(map[int]int)
	for i := 0; i < q; i++ {
		fmt.Fscan(in, &t, &x)

		switch t {
		case 1:
			list[top] = x
			top--
		case 2:
			list[btm] = x
			btm++
		case 3:
			ans := list[top+x]
			fmt.Println(ans)
		}
	}
}
