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
	N, exams, res := nextInt(), []exam{}, 0
	for i := 0; i < N; i++ {
		exams = append(exams, exam{nextInt(), 0, 0})
	}

	p, m := 0, 0
	for i := 0; i < N; i++ {
		b := nextInt()
		d := exams[i].a - b
		exams[i] = exam{exams[i].a, b, d}

		if d > 0 {
			p += d
		} else {
			m += abs(d)
			res++
		}
	}

	if m == 0 {
		fmt.Println(0)
		return
	} else if p < m {
		fmt.Println(-1)
		return
	}

	sort.Slice(exams, func(i, j int) bool {
		if exams[i].diff > exams[j].diff {
			return true
		}
		return false
	})

	for _, e := range exams {
		if e.diff <= 0 {
			continue
		}

		m -= e.diff
		res++
		if m <= 0 {
			fmt.Println(res)
			return
		}
	}
}

type exam struct {
	a, b, diff int
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
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
