package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	n, a, sum := nextInt(), []int{}, []int{}
	for i := 0; i < n; i++ {
		a = append(a, nextInt())
		if i == 0 {
			sum = append(sum, a[0])
		} else {
			sum = append(sum, sum[i-1]+a[i])
		}
	}

	x := nextInt()
	res := (x / sum[n-1]) * n
	x %= sum[n-1]
	for i := 0; i < n; i++ {
		res++
		if x < sum[i] {
			break
		}
	}

	// res, x := 0, nextInt()
	// res, x = x/sum[n-1]*n, x%sum[n-1]
	// for i := range sum {
	// 	if x <= sum[i] {
	// 		res += i + 1
	// 		break
	// 	}
	// }
	fmt.Println(res)
}

const (
	initialBufSize = 10000
	maxBufSize     = 1000000
)

var (
	sc *bufio.Scanner = func() *bufio.Scanner {
		sc := bufio.NewScanner(os.Stdin)
		buf := make([]byte, initialBufSize)
		sc.Buffer(buf, maxBufSize)
		return sc
	}()
)

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

// var sc = bufio.NewScanner((os.Stdin))
