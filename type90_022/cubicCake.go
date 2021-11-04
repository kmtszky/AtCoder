package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	gcf := 0
	gcf = calc(b, c)
	gcf = calc(a, gcf)

	// 切る回数
	fmt.Println((a+b+c)/gcf - 3)
}

// 最大公約数を計算
func calc(b, c int) int {
	for {
		if b == 0 || c == 0 {
			return int(math.Max(float64(b), float64(c)))
		}
		if b >= c {
			b = b % c
		} else {
			c = c % b
		}
	}
}
