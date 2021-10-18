package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	fmt.Scanf("%s", &s)

	var s1, s2, s3 int
	s1, _ = strconv.Atoi(s[0:1])
	s2, _ = strconv.Atoi(s[1:2])
	s3, _ = strconv.Atoi(s[2:3])

	numbers := []int{s1, s2, s3}
	sum := 0
	for _, number := range numbers {
		if number == 1 {
			sum += 1
		}
		continue
	}
	fmt.Printf("%d\n", sum)
}
