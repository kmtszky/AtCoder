package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	pt, px, py := 0, 0, 0

	for i := 0; i < n; i++ {
		var t, x, y int
		fmt.Scan(&t, &x, &y)

		d := abs(x-px) + abs(y-py)
		dt := t - pt

		if dt < d || (dt-d)%2 == 1 {
			fmt.Println("No")
			return
		}
		pt, px, py = t, x, y
	}
	fmt.Println("Yes")
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
