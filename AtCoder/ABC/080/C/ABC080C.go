package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	N := nextInt()
	F := make([][]int, N)
	for i := range F {
		F[i] = make([]int, 10)
		for j := range F[i] {
			F[i][j] = nextInt()
		}
	}

	P := make([][]int, N)
	for i := range P {
		P[i] = make([]int, 11)
		for j := range P[i] {
			P[i][j] = nextInt()
		}
	}

	res := math.MinInt64
	for bit := 1; bit < (1 << uint(10)); bit++ {
		bs := make([]int, 10)
		for i := 0; i < 10; i++ {
			bs[i] = bit >> uint(i) & 1
		}

		tmp := 0
		for i := 0; i < N; i++ {
			cnt := 0
			for j := 0; j < 10; j++ {
				if bs[j] == 1 && F[i][j] == 1 {
					cnt++
				}
			}
			tmp += P[i][cnt]
		}
		res = max(res, tmp)
	}
	fmt.Println(res)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
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
