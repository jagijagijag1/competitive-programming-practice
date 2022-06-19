package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var startTime time.Time

func main() {
	startTime = time.Now()

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

	// solveGreedy(d, c, s)
	solveHillClimbing(d, c, s)
}

func solveHillClimbing(d int, c []int, s [][]int) {
	out := make([]int, d)

	// random solution
	rand.Seed(startTime.UnixNano())
	for i := range out {
		out[i] = rand.Intn(26)
	}
	tmpscore := score(d, c, out, s)

	for time.Now().Sub(startTime).Seconds()+0.03 < 2.0 {
		d, q := rand.Intn(d), rand.Intn(26)
		old := out[d]
		out[d] = q

		newscore := score(d, c, out, s)
		if newscore < tmpscore {
			out[d] = old
		} else {
			tmpscore = newscore
		}
	}

	for i := range out {
		fmt.Println(out[i] + 1)
	}
}

func solveGreedy(d int, c []int, s [][]int) {
	out := []int{}
	for i := 0; i < d; i++ {
		tmpscore, tmpindex := math.MinInt64, 0
		for j := 0; j < 26; j++ {
			// tt := score(i, c, append(out, j), s)
			tt := eval(d, c, append(out, j), s, 13)
			if tmpscore < tt {
				tmpindex = j
				tmpscore = tt
			}
		}
		out = append(out, tmpindex)
	}
	for i := range out {
		fmt.Println(out[i] + 1)
	}
}

func eval(D int, c, out []int, s [][]int, k int) int {
	score := 0
	last := make([]int, 26)

	for d := range out {
		last[out[d]] = d + 1
		for i := range last {
			score -= (d + 1 - last[i]) * c[i]
		}
		score += s[d][out[d]]
	}

	for d := len(out); d < min(len(out)+k, D); d++ {
		for i := 0; i < 26; i++ {
			score -= (d + 1 - last[i]) * c[i]
		}
	}

	return score
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

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
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
