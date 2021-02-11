package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	n, m := nextInt(), map[rune]int{}
	for i := 'a'; i <= 'z'; i++ {
		m[i] = 1000000000
	}

	for i := 0; i < n; i++ {
		s, sm := nextLine(), map[rune]int{}
		for _, ss := range s {
			sm[ss]++
		}
		for j := range m {
			m[j] = min(sm[j], m[j])
		}

	}
	res := []rune{}
	for i := 'a'; i <= 'z'; i++ {
		for j := 0; j < m[i]; j++ {
			res = append(res, i)
		}
	}
	fmt.Println(string(res))
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main2() {
	sc.Split(bufio.ScanWords)
	n, m := nextInt(), map[rune]tuple{}
	for i := 'a'; i <= 'z'; i++ {
		m[i] = tuple{-1, -1}
	}

	for i := 0; i < n; i++ {
		s, ms := nextLine(), map[rune]int{}
		for _, ss := range s {
			ms[ss]++
		}
		for j := 'a'; j <= 'z'; j++ {
			if i == 0 {
				// init
				m[j] = tuple{ms[j], 1}
				continue
			}

			if i != 0 && ms[j] == 0 {
				m[j] = tuple{0, 0}
				continue
			}

			if ms[j] < m[j].cnt1 {
				m[j] = tuple{ms[j], m[j].cnt2 + 1}
			} else {
				m[j] = tuple{m[j].cnt1, m[j].cnt2 + 1}
			}
		}
	}

	res := []rune{}
	for i := 'a'; i <= 'z'; i++ {
		if 0 < m[i].cnt2 {
			upper := m[i].cnt1
			for j := 0; j < upper; j++ {
				res = append(res, i)
			}
		}
	}
	fmt.Println(string(res))
}

type tuple struct {
	cnt1, cnt2 int
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
