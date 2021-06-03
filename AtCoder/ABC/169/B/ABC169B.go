package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	// n, a, res := nextInt(), []int{}, 1
	// for i := 0; i < n; i++ {
	// 	a = append(a, nextInt())
	// 	if a[i] == 0 {
	// 		fmt.Println(0)
	// 		return
	// 	}
	// }

	// for i := 0; i < n; i++ {
	// 	if a[i] <= 1000000000000000000/res {
	// 		res *= a[i]
	// 	} else {
	// 		fmt.Println(-1)
	// 		return
	// 	}
	// }

	n, a, res, upper := nextInt(), []int{}, big.NewInt(1), big.NewInt(1000000000000000000)
	for i := 0; i < n; i++ {
		a = append(a, nextInt())
		if a[i] == 0 {
			fmt.Println(0)
			return
		}
	}

	for i := 0; i < n; i++ {
		ba := big.NewInt(int64(a[i]))
		res.Mul(res, ba)
		if res.Cmp(upper) == 1 {
			fmt.Println(-1)
			return
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
