package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, n)

		for j := 0; j < n; j++ {
			fmt.Scan(&a[i][j])
		}
	}

	var m, x, y int
	fmt.Scan(&m)

	bad := [11][11]bool{}

	for i := 0; i < m; i++ {
		fmt.Scan(&x, &y)
		x--
		y--
		bad[x][y] = true
		bad[y][x] = true
	}

	runners := make([]int, n)
	for i := 0; i < n; i++ {
		runners[i] = i
	}

	// 一区間を走るのに時間が最大1000必要であり、かつコースはn区間ある
	limit := 1000*n + 1
	ans := limit

	for {
		t := 0
		for i := 0; i < n; i++ {
			// i区画を走る選手j
			j := runners[i]

			if i+1 < n && bad[j][runners[i+1]] {
				t = -1
				break
			}
			t += a[j][i]
		}
		if t != -1 {
			ans = min(ans, t)
		}

		// 全探索が完了したら終了
		if !nextPermutation((sort.IntSlice(runners))) {
			break
		}
	}

	if ans == limit {
		ans = -1
	}
	fmt.Println(ans)
}

func min(timeA, timeB int) int {
	a := float64(timeA)
	b := float64(timeB)
	return int(math.Min(a, b))
}

// https://www.kumadon.biz/go_permutation/
// N!通りを全探索する
func nextPermutation(x sort.Interface) bool {
	n := x.Len() - 1
	if n < 1 {
		return false
	}
	j := n - 1
	for ; !x.Less(j, j+1); j-- {
		if j == 0 {
			return false
		}
	}
	l := n
	for !x.Less(j, l) {
		l--
	}
	x.Swap(j, l)
	for k, l := j+1, n; k < l; {
		x.Swap(k, l)
		k++
		l--
	}
	return true
}
