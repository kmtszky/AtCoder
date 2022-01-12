package main

import (
	"fmt"
)

func main() {
	var n, p, q int
	fmt.Scan(&n, &p, &q)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				for l := 0; l < n; l++ {
					for m := 0; m < n; m++ {
						res := equalQ(a[i], a[j], a[k], a[l], a[m], p, q)
						if res {
							ans++
						}
					}
				}
			}
		}
	}
	fmt.Println(ans)
}

func equalQ(a1, a2, a3, a4, a5, p, q int) bool {
	return a1%p*a2%p*a3%p*a4%p*a5%p == q
}

