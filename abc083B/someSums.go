package main

import "fmt"

func main() {
	var n, a, b, total int
	fmt.Scanf("%d %d %d", &n, &a, &b)

	for i := 1; i <= n; i++ {
		sum := sum(i)
		if sum >= a && sum <= b {
			total += i
		}
	}
	fmt.Println(total)
}

func sum(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n = n / 10
	}
	return sum
}
