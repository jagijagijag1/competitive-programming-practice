package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	N, H, a, b := nextInt(), nextInt(), []int{}, []int{}
	for i := 0; i < N; i++ {
		a = append(a, nextInt())
		b = append(b, nextInt())
	}
	sort.Ints(a)
	sort.Ints(b)

	res, dmg := 0, 0
	for i := N - 1; 0 <= i; i-- {
		if b[i] < a[N-1] {
			break
		}
		res++
		dmg += b[i]
		if H <= dmg {
			fmt.Println(res)
			return
		}
	}

	res += int(math.Ceil(float64(H-dmg) / float64(a[N-1])))
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
