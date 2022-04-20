package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)

	a := make([]int, m)
	b := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &a[i], &b[i])
	}

	var k int
	fmt.Fscan(in, &k)

	c := make([]int, k)
	d := make([]int, k)
	for i := 0; i < k; i++ {
		fmt.Fscan(in, &c[i], &d[i])
	}

	ans := 0
	for i := 0; i < (1 << k); i++ { // K人が2枚のどちらの皿にボールを置くかは2^k通り
		dish := map[int]int{}    // map初期化
		for j := 0; j < k; j++ { // お皿はK組あるので、全網羅は K*2^k
			if (i>>j)&1 == 0 { // iのj桁目の数字が1か0か判定
				dish[c[j]] = 1 // 0のときCにボールが入る
			} else {
				dish[d[j]] = 1 // 1のときDにボールが入る
			}
		}

		match := 0
		for j := 0; j < m; j++ {
			if dish[a[j]] == 0 || dish[b[j]] == 0 { // 皿a, b両方にボールが入っていない場合
				continue
			}
			match++ // 条件を満たしている個数をインクリメント
		}
		ans = int(math.Max(float64(match), float64(ans))) // 満たされる条件の最大数
	}
	fmt.Println(ans)
}
