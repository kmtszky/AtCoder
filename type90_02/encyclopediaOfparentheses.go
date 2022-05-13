package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	if n%2 == 1 {
		return
	}

	for i := 0; i < (1 << n); i++ { // 2^n通りの文字列を生成
		ans := ""
		abs := 0
		for j := n - 1; j >= 0; j-- {
			if (i & (1 << j)) == 0 { // iのjビット目（2^jの位）が0の時
				abs++
				ans += "("
			} else { // iのjビット目（2^jの位）が1の時
				abs--
				ans += ")"
			}
			if abs < 0 { // )が(よりも多く、正しくないかっこが含まれる場合
				break
			}
		}
		if abs == 0 { // (の文字数 = )の文字数
			fmt.Println(ans)
		}
	}
}
