package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	n, m := nextInt(), nextInt()
	ss := make([][]int, m)
	for i := 0; i < m; i++ {
		k := nextInt()
		for j := 0; j < k; j++ {
			ss[i] = append(ss[i], nextInt()-1)
		}
	}
	p := make([]int, m)
	for i := 0; i < m; i++ {
		p[i] = nextInt()
	}

	res := 0
	for bit := 0; bit < (1 << uint(n)); bit++ {
		on := 0
		for i := 0; i < m; i++ {
			s_on := 0
			for j := range ss[i] {
				if bit>>uint(ss[i][j])&1 == 1 {
					s_on++
				}
			}
			if s_on%2 == p[i] {
				on++
			}
		}
		if on == m {
			res++
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
