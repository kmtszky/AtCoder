package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	stringAList := ""
	var sc = bufio.NewScanner(os.Stdin)
	if sc.Scan() {
		stringAList += sc.Text()
	}

	arrayStringA := strings.Split(stringAList, " ")
	arrayIntA := stringToInt(arrayStringA)

	maxTryNum := 0
	for {
		if !evenCheck(arrayIntA) {
			break
		}
		for i, v := range arrayIntA {
			arrayIntA[i] = v / 2
		}
		maxTryNum += 1
	}
	fmt.Println(maxTryNum)
}

func stringToInt(arrayStringA []string) []int {
	sliceA := make([]int, len(arrayStringA))
	for i, v := range arrayStringA {
		sliceA[i], _ = strconv.Atoi(v)
	}
	return sliceA
}

func evenCheck(arrayIntA []int) bool {
	for _, v := range arrayIntA {
		if v%2 == 1 {
			return false
		}
	}
	return true
}
