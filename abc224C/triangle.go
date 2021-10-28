package main

import "fmt"

type grid struct {
	x int
	y int
}

func main() {
	var n int
	fmt.Scan(&n)

	grid := make([]grid, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&grid[i].x, &grid[i].y)
	}

	ans := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if ((grid[j].x-grid[i].x)*(grid[k].y-grid[i].y) - (grid[k].x-grid[i].x)*(grid[j].y-grid[i].y)) != 0 {
					ans++
				}
			}

		}
	}
	fmt.Println(ans)
}
