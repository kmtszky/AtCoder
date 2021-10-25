package main

import (
	"fmt"
)

func main() {
	var h, w int
	fmt.Scan(&h, &w)

	grid := make([][]int, h)
	for i := 0; i < h; i++ {
		grid[i] = make([]int, w)

		for j := 0; j < w; j++ {
			fmt.Scan(&grid[i][j])
		}
	}

	for i := 0; i < h-1; i++ {
		for j := 0; j < w-1; j++ {
			if grid[i][j]+grid[i+1][j+1] > grid[i+1][j]+grid[i][j+1] {
				fmt.Println("No")
				return
			}
		}
	}
	fmt.Println("Yes")
}
