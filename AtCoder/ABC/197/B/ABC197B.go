package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	h, _, x, y, s := nextInt(), nextInt(), nextInt()-1, nextInt()-1, []string{}
	for i := 0; i < h; i++ {
		s = append(s, nextLine())
	}

	if s[x][y] == '#' {
		fmt.Println(0)
		return
	}

	res := 1
	for i := y + 1; i < len(s[0]); i++ {
		if s[x][i] != '#' {
			res++
		} else {
			break
		}
	}
	for i := y - 1; 0 <= i; i-- {
		if s[x][i] != '#' {
			res++
		} else {
			break
		}
	}
	for i := x + 1; i < h; i++ {
		if s[i][y] != '#' {
			res++
		} else {
			break
		}
	}
	for i := x - 1; 0 <= i; i-- {
		if s[i][y] != '#' {
			res++
		} else {
			break
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
