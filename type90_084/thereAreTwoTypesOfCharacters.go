package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n int
	var s string
	fmt.Fscan(in, &n, &s)

	// o = 111, x=120
	A := make([]int, n+1)
	B := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		if s[i-1] == 111 {
			A[i] = i
			B[i] = B[i-1]
		} else if s[i-1] == 120 {
			A[i] = A[i-1]
			B[i] = i
		}
	}

	ans := 0
	for i := 1; i < n+1; i++ {
		ans += int(math.Min(float64(A[i]), float64(B[i])))
	}
	fmt.Println(ans)
}
