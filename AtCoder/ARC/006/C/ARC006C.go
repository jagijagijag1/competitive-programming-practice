package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// func main() {
func mainARC006C() {
	sc.Split(bufio.ScanWords)
	N, w := nextInt(), []int{}
	for i := 0; i < N; i++ {
		w = append(w, nextInt())
	}

	fmt.Printf("%d\n", ARC006C(w))
}

// ARC006C ...
func ARC006C(w []int) int {
	stack := [][]int{{w[0]}}
	for i := 1; i < len(w); i++ {
		c, tmpi, tmpc := w[i], 0, 1000000
		for i := 0; i < len(stack); i++ {
			s := stack[i]

			diff := s[len(s)-1] - c
			if 0 <= diff && diff < tmpc {
				tmpc = diff
				tmpi = i
			}
		}
		if tmpc != 1000000 {
			stack[tmpi] = append(stack[tmpi], c)
		} else {
			stack = append(stack, []int{c})
		}
	}

	return len(stack)
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
