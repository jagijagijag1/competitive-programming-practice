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
	e := map[int][]int{}
	for i := 0; i < m; i++ {
		xx, yy := nextInt()-1, nextInt()-1
		e[xx] = append(e[xx], yy)
		e[yy] = append(e[yy], xx)
	}

	res := 0
	for bit := 0; bit < (1 << uint(n)); bit++ {
		size, valid := 0, true
		for i := 0; i < n; i++ {
			if bit>>uint(i)&1 == 1 {
				size++
			}
			for j := i + 1; j < n; j++ {
				if bit>>uint(i)&1 == 1 && bit>>uint(j)&1 == 1 {
					if !contains(e[i], j) {
						valid = false
					}
				}
			}
		}

		if valid && res < size {
			res = size
		}
	}
	fmt.Println(res)
}

func contains(s []int, e int) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
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
