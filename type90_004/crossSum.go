package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var h, w int
	fmt.Scan(&h, &w)

	values := [2000][2000]int{}
	rowSum := [2000]int{}
	colSum := [2000]int{}
	sc := bufio.NewScanner(os.Stdin)
	for i := 0; i < h; i++ {
		sc.Scan()
		row := sc.Text()
		numbers := strings.Split(row, " ")
		for j, v := range numbers {
			values[i][j], _ = strconv.Atoi(v)
			rowSum[i] += values[i][j]
			colSum[j] += values[i][j]
		}
	}

	for i := 0; i < h; i++ {
		ans := make([]int, w)
		for j := 0; j < w; j++ {
			ans[j] = rowSum[i] + colSum[j] - values[i][j]
		}
		fmt.Println(strings.Trim(fmt.Sprint(ans), "[]"))
	}
}
