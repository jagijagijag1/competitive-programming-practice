package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	N := nextInt()
	to, cost := make([][]int, N), make([][]int, N)

	for i := 0; i < N-1; i++ {
		a, b, w := nextInt()-1, nextInt()-1, nextInt()
		to[a], cost[a] = append(to[a], b), append(cost[a], w)
		to[b], cost[b] = append(to[b], a), append(cost[b], w)
	}

	res := make([]int, N)
	for i := range res {
		res[i] = -1
	}
	q := []int{0}
	for 0 < len(q) {
		h := q[0]
		q = q[1:]

		for i := 0; i < len(to[h]); i++ {
			t, w := to[h][i], cost[h][i]
			if res[t] != -1 {
				continue
			}
			res[t] = (res[h] + w) % 2
			q = append(q, t)
		}
	}

	for _, r := range res {
		fmt.Println(r)
	}
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
