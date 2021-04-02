package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	N, K := nextInt(), nextInt()
	cnt := make([]int, K)

	for i := 1; i <= N; i++ {
		cnt[i%K]++
	}

	res := 0
	for a := 0; a < K; a++ {
		b, c := (K-a)%K, (K-a)%K
		if (b+c)%K == 0 {
			res += cnt[a] * cnt[b] * cnt[c]
		}
	}
	fmt.Println(res)
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

// const (
// 	initialBufSize = 10000
// 	maxBufSize     = 1000000
// )

// var (
// 	sc *bufio.Scanner = func() *bufio.Scanner {
// 		sc := bufio.NewScanner(os.Stdin)
// 		buf := make([]byte, initialBufSize)
// 		sc.Buffer(buf, maxBufSize)
// 		return sc
// 	}()
// )
