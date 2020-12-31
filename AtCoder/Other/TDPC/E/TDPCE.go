package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// func main() {
func mainTDPCE() {
	sc.Split(bufio.ScanWords)
	D := nextInt()
	tmpN, N := nextLine(), []int{}
	for i := range tmpN {
		num, _ := strconv.Atoi(string(tmpN[i]))
		N = append(N, num)
	}

	fmt.Printf("%d\n", TDPCE(D, N))
}

// TDPCE ...
func TDPCE(D int, N []int) int {
	mod := 1000000007

	// dp[i][smaller][j]: i桁目までNに一致し，i+1桁目がNより小さい(smaller=1)/Nより大きい(smaller=0)ときの，Dの剰余がjとなる数
	dp := make([][][]int, len(N)+1)
	for i := range dp {
		dp[i] = make([][]int, 2)
		for j := range dp[i] {
			dp[i][j] = make([]int, D)
		}
	}

	dp[0][0][0] = 1
	for i := 0; i < len(N); i++ {
		for j := 0; j < D; j++ {
			for k := 0; k < 10; k++ {
				// 走査中のi+1桁目がNのi+1桁目(N[i])より大きい場合，はN以下の数値にならないので考慮不要
				if k == N[i] {
					// 走査中のi+1桁目がNのi+1桁目と一致する場合，走査中のi桁目とNのi桁目が一致していたケース(dp[i][0][*])の値を使って更新
					// (dp[i+1][0][*]はそれ以外からは更新されない)
					dp[i+1][0][(j+k)%D] += dp[i][0][j]
					dp[i+1][0][(j+k)%D] %= mod
				} else if k < N[i] {
					// 走査中のi+1桁目がNのi+1桁目(N[i])より小さい場合，走査中のi桁目とNのi桁目が一致していたケース(dp[i][0][*])の値を使って更新
					// dp[i+1][1][*]はdp[i][1][*]からも更新されるがその処理は後述
					dp[i+1][1][(j+k)%D] += dp[i][0][j]
					dp[i+1][1][(j+k)%D] %= mod
				}
				// 走査中のi桁目がNのi桁目より小さいケースは，k=0〜9のすべてで更新
				dp[i+1][1][(j+k)%D] += dp[i][1][j]
				dp[i+1][1][(j+k)%D] %= mod
			}
		}
	}

	return dp[len(N)][0][0] + dp[len(N)][1][0] - 1
}

var sc = bufio.NewScanner((os.Stdin))

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	l := nextLine()
	i, e := strconv.Atoi(l)
	if e != nil {
		panic(e)
	}
	return i
}
