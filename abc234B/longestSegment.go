package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var n int

	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)

	x := make([]float64, n)
	y := make([]float64, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x[i], &y[i])
	}

	ans := 0.0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			d := calc(x[i], y[i], x[j], y[j])
			ans = math.Max(ans, d)
		}
	}
	fmt.Println(ans)
}

func calc(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}
