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
	N, P := nextInt(), nextInt()

	if N == 1 {
		fmt.Println(P)
		return
	}

	if P == 1 {
		fmt.Println(1)
		return
	}

	pfs, res := primeFactorization(P), 1
	for k, v := range pfs {
		c := v / N
		if c != 0 {
			res *= pow(k, c)
		}
	}

	fmt.Println(res)
}

func pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
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
