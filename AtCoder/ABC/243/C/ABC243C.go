package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	p := make([]tuple, n)
	for i := 0; i < n; i++ {
		p[i] = tuple{i, nextInt(), nextInt()}
	}
	sort.Slice(p, func(i, j int) bool {
		if p[i].y < p[j].y {
			return true
		} else if p[i].y == p[j].y {
			if p[i].x < p[j].x {
				return true
			}
		}
		return false
	})

	s := nextLine()
	for i := 0; i < n; i++ {
		pre := p[i]
		if s[pre.id] != 'R' {
			continue
		}

		for j := i + 1; j < n; j++ {
			cur := p[j]
			if pre.y != cur.y {
				i = j
				break
			}
			if s[cur.id] == 'L' {
				fmt.Println("Yes")
				return
			}
		}
	}

	fmt.Println("No")
}

type tuple struct {
	id, x, y int
}

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

// var sc = bufio.NewScanner((os.Stdin))
