package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	sc := bufio.NewScanner(os.Stdin)
	stringAList := ""
	if sc.Scan() {
		stringAList = sc.Text()
	}

	arrayStringA := strings.Split(stringAList, " ")
	arrayIntA := stringToInt(arrayStringA)

	sort.Sort(sort.Reverse(sort.IntSlice(arrayIntA)))

	Alice, Bob := 0, 0
	for i, a := range arrayIntA {
		if i%2 == 0 {
			Alice += a
		} else {
			Bob += a
		}
	}
	fmt.Println(Alice - Bob)
}

func stringToInt(arrayStringA []string) []int {
	sliceA := make([]int, len(arrayStringA))
	for i, v := range arrayStringA {
		sliceA[i], _ = strconv.Atoi(v)
	}
	return sliceA
}
