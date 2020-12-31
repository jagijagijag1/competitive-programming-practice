package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// func main() {
func mainABC116B() {
	sc.Split(bufio.ScanWords)
	s := nextInt()

	fmt.Printf("%d\n", ABC116B(s))
}

// ABC116B ...
func ABC116B(s int) int {
	fs, fsmap := make([]int, 1000000), map[int]struct{}{}
	fs[0], fs[1], fsmap[s] = 0, s, struct{}{}
	for i := 2; i < 1000000; i++ {
		if fs[i-1]%2 == 0 {
			fs[i] = fs[i-1] / 2
		} else {
			fs[i] = 3*fs[i-1] + 1
		}

		if _, ok := fsmap[fs[i]]; ok {
			return i
		}
		fsmap[fs[i]] = struct{}{}
	}

	return 0
}

func minInt(a, b int) int {
	if a > b {
		return b
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
