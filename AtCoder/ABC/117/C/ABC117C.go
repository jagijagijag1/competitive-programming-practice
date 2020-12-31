package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

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

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

//func main() {
func mainABC117C() {
	sc.Split(bufio.ScanWords)
	N, M, X := nextInt(), nextInt(), []int{}
	for i := 0; i < M; i++ {
		X = append(X, nextInt())
	}

	fmt.Printf("%d\n", ABC117C(N, X))
}

// ABC117C ...
func ABC117C(N int, X []int) int {
	sort.Ints(X)

	if N == 1 {
		return X[len(X)-1] - X[0]
	}

	dists := []int{}
	for i := 0; i < len(X)-1; i++ {
		dists = append(dists, X[i+1]-X[i])
	}
	sort.Ints(dists)

	res := 0
	for i := 0; i < len(dists)-N+1; i++ {
		res += dists[i]
	}

	return res
}

// ABC117CWA ...
func ABC117CWA(N int, X []int) int {
	sort.Ints(X)

	if N == 1 {
		return X[len(X)-1] - X[0] + 1
	}

	div := [][]int{X}
	for d := 0; d < N-1; d++ {
		// select longest range
		var t []int
		max, maxIndex := -1, 0
		for i := 0; i < len(div); i++ {
			if max < len(div[i]) {
				max = len(div[i])
				maxIndex = i
			}
		}

		if max == 1 {
			break
		}

		t = div[maxIndex]

		l, r := t[0], t[len(t)-1]
		llist, rlist := []int{l}, []int{r}
		for i := 1; i < len(t)-1; i++ {
			if abs(t[i]-l) < abs(t[i]-r) {
				llist = append(llist, t[i])
				l = t[i]
			} else {
				rlist = append(rlist, t[i])
				sort.Ints(rlist)
				r = rlist[0]
			}
		}

		newdiv := [][]int{llist, rlist}
		for i := 0; i < len(div)-1; i++ {
			if i == maxIndex {
				continue
			}
			newdiv = append(newdiv, div[i])
		}
		div = newdiv
	}

	// count res
	//res := len(div)
	res := 0
	for _, d := range div {
		res += d[len(d)-1] - d[0]
	}

	return res
}
