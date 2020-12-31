package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// func main() {
func mainABC038D() {
	sc.Split(bufio.ScanWords)
	N, w, h := nextInt(), []int{}, []int{}
	for i := 0; i < N; i++ {
		w = append(w, nextInt())
		h = append(h, nextInt())
	}

	fmt.Printf("%d\n", ABC038D(w, h))
}

// ABC038D ...
func ABC038D(w, h []int) (res int) {
	rects := []Rect{}
	for i := 0; i < len(w); i++ {
		rects = append(rects, Rect{w[i], h[i]})
	}
	sort.Slice(rects, func(i, j int) bool {
		if rects[i].w < rects[j].w {
			return true
		} else if rects[i].w == rects[j].w {
			if rects[i].h > rects[j].h {
				return true
			}
		}
		return false
	})

	dp := make([]int, len(w))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1000000
	}

	dp[0], res = rects[0].h, 1
	for i := 0; i < len(dp); i++ {
		if dp[res-1] < rects[i].h {
			dp[res] = rects[i].h
			res++
		} else {
			idx := sort.Search(len(dp), func(j int) bool {
				return dp[j] >= rects[i].h
			})
			dp[idx] = rects[i].h
		}
	}

	return
}

// Rect ...
type Rect struct {
	w, h int
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
