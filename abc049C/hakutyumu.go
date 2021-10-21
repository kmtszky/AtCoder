package main

import (
	"fmt"
	"strings"
)

// 英小文字からなる文字列 S が与えられます。
// Tが空文字列である状態から始め、以下の操作を好きな回数繰り返すことで S=T とすることができるか判定してください。
// T の末尾に dream dreamer erase eraser のいずれかを追加する。
func main() {
	var s string
	fmt.Scan(&s)

	words := [2]string{"dream", "erase"}
	replacedWords := [2]string{"D", "E"}
	for i := 0; i < len(words); i++ {
		s = strings.Replace(s, words[i], replacedWords[i], -1)
	}

	if deleteWords(s) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func deleteWords(s string) bool {
	shortWords := [4]string{"Der", "Er", "D", "E"}
	for i := 0; i < len(shortWords); i++ {
		s = strings.Replace(s, shortWords[i], "", -1)
		if s == "" {
			break
		}
	}
	return s == ""
}
