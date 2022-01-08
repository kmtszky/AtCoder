package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)

	f := calc(calc(calc(t)+t) + calc(calc(t)))
	fmt.Println(f)
}

func calc(t int) int {
	f := t*t + 2*t + 3
	return f
}
