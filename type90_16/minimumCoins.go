package main

import (
	"fmt"
	"math"
)

func main() {
	var n, a, b, c int
	fmt.Scan(&n, &a, &b, &c)

	ans := int(1e4)
	for i := 0; i < 9999; i++ {
		for j := 0; i+j < 9999; j++ {
			res := n - a*i - b*j

			if res%c != 0 || res < 0 {
				continue
			}
			k := res / c

			ans = int(math.Min(float64(ans), float64(i+j+k)))
		}
	}
	fmt.Println(ans)
}
