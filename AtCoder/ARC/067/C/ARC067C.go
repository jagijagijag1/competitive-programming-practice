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

	pfs := map[int]int{}
	for cn := 2; cn <= N; cn++ {
		n := cn
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

		if _, ok := pfs[n]; n > 2 && ok {
			pfs[n]++
		} else {
			pfs[n] = 1
		}
	}

	res := 1
	for k, v := range pfs {
		if k == 1 {
			continue
		}
		res *= (v + 1)
		res %= 1000000007
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
