package main

import "fmt"

func main() {
	var n, sum int
	fmt.Scanf("%d %d", &n, &sum)

	x, y, z := valueCheck(n, sum)
	fmt.Println(x, y, z)
}

func valueCheck(n, sum int) (int, int, int) {
	for x := 0; x <= n; x++ {
		for y := 0; y <= n; y++ {
			if sum == 10000*x+5000*y+1000*(n-x-y) && n-x-y >= 0 {
				return x, y, n - x - y
			}
		}
	}
	return -1, -1, -1
}
