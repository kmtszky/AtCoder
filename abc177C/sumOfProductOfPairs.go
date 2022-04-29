package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(in, &n)

	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	mod := int(1e9 + 7) // 10^9+7

	s := make([]int, n+1)
	for i := 0; i < n; i++ {
		s[i+1] += s[i] + a[i] // 累積和
	}

	ans := 0
	for i := 0; i < n-1; i++ {
		sum := s[n] - s[i+1]      // a[i]のとき、a[i+1]〜を足した累積和が必要なので、合計値からs[i+1]を除く（s[0]は0なので、iに+1する必要あり）
		ans += a[i] * (sum % mod) // オーバーフローしないように先に%modしておく
		ans %= mod
	}

	fmt.Println(ans)
}
