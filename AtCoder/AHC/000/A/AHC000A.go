package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	d := nextInt()
	c, s := make([]int, 26), make([][]int, d)
	for i := range s {
		s[i] = make([]int, 26)
	}

	for i := range c {
		c[i] = nextInt()
	}
	for i := range s {
		for j := range s[i] {
			s[i][j] = nextInt()
		}
	}

	out := []int{}
	for i := 0; i < d; i++ {
		// fmt.Println("current seq:", out)
		tmpscore, tmpindex := math.MinInt64, 0
		for j := 0; j < 26; j++ {
			tt := score(i, c, append(out, j), s)
			// fmt.Println("  ", i, j, "next:", tt, ", tmp:", tmpscore)
			if tmpscore < tt {
				tmpindex = j
				tmpscore = tt
				// fmt.Println("    now select", j, "with score", tt)
			}
		}
		out = append(out, tmpindex)
	}
	for i := range out {
		fmt.Println(out[i] + 1)
	}
}

func score(D int, c, out []int, s [][]int) int {
	score := 0
	last := make([]int, 26)

	for d := range out {
		last[out[d]] = d + 1
		for i := range last {
			score -= (d + 1 - last[i]) * c[i]
		}
		score += s[d][out[d]]
	}

	return score
}

// var sc = bufio.NewScanner((os.Stdin))

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
