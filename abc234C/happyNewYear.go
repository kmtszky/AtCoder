package main

import (
	"fmt"
	"strings"
)

func main() {
	var k uint64
	fmt.Scan(&k)

	// 2進数に変換
	binaryNum := fmt.Sprintf("%b", k)

	// 1を2に変換
	ans := strings.Replace(binaryNum, "1", "2", -1)
	fmt.Println(ans)
}
