package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n string
	var k int
	fmt.Scan(&n, &k)

	for i := 0; i < k; i++ {
		eight, _ := strconv.ParseInt(n, 8, 64)
		nine := []byte(strconv.FormatInt(eight, 9))

		for j := range nine {
			if nine[j] == '8' {
				nine[j] = '5'
			}
		}
		n = string(nine)
	}
	fmt.Println(n)
}
