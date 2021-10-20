package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	dList := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&dList[i])
	}
	sort.Ints(dList)

	kagamiMochi := 1
	for i := 0; i < n; i++ {
		if i == 0 {
			continue
		}
		if dList[i] > dList[i-1] {
			kagamiMochi++
		}
	}
	fmt.Println(kagamiMochi)
}
