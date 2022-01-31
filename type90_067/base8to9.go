package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n string
	var k int
	fmt.Scan(&n, &k)

	for i := 0; i < k; i++ {
		eight, _ := strconv.ParseInt(n, 8, 64)
		nine := strconv.FormatInt(eight, 9)
		n = strings.Replace(nine, "8", "5", -1)
	}
	fmt.Println(n)
}
