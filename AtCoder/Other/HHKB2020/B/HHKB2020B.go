package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	h, w, res := nextInt(), nextInt(), 0

	last_row := ""
	for i := 0; i < h; i++ {
		s := nextLine()
		for j := 0; j < w; j++ {
			if j != 0 && s[j-1] == '.' && s[j] == '.' {
				res++
			}
			if i != 0 && s[j] == '.' && last_row[j] == '.' {
				res++
			}
		}
		last_row = s
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
