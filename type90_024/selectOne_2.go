package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	a := [1000]int{}
	b := [1000]int{}
	sc := bufio.NewScanner(os.Stdin)
	for i := 0; i < 2; i++ {
		sc.Scan()
		txt := sc.Text()
		ns := strings.Split(txt, " ")
		for j, v := range ns {
			if i == 0 {
				a[j], _ = strconv.Atoi(v)
			} else {
				b[j], _ = strconv.Atoi(v)
			}
		}
	}

	abs := 0
	for i := 0; i < n; i++ {
		abs += calc(b[i] - a[i])
	}

	if (abs+k)%2 == 0 && abs <= k {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func calc(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
