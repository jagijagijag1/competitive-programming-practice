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

	pfs, slicePfs := primeFactorization(N), []int{}
	for k, v := range pfs {
		for i := 0; i < v; i++ {
			slicePfs = append(slicePfs, k)
		}
	}

	n, res := len(slicePfs), 1000000000
	for bit := 0; bit < (1 << uint(n)); bit++ {
		a, b := 1, 1
		for i := 0; i < n; i++ {
			if bit>>uint(i)&1 == 1 {
				a *= slicePfs[i]
			} else {
				b *= slicePfs[i]
			}
		}

		res = min(res, max(len(fmt.Sprint(a)), len(fmt.Sprint(b))))
	}

	fmt.Println(res)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func primeFactorization(n int) (pfs map[int]int) {
	pfs = make(map[int]int)

	for n%2 == 0 {
		if _, ok := pfs[2]; ok {
			pfs[2]++
		} else {
			pfs[2] = 1
		}
		n = n / 2
	}

	for i := 3; i*i <= n; i = i + 2 {
		for n%i == 0 {
			if _, ok := pfs[i]; ok {
				pfs[i]++
			} else {
				pfs[i] = 1
			}
			n = n / i
		}
	}

	if n > 2 {
		pfs[n] = 1
	}

	return
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
