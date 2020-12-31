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
	Q, H, S, D, N := nextInt(), nextInt(), nextInt(), nextInt(), nextInt()
	qhsd := []tuple{{Q, 0.25, float64(Q) / 0.25}, {H, 0.5, float64(H) / 0.5}, {S, 1.0, float64(S) / 1.0}, {D, 2.0, float64(D) / 2.0}}

	sort.Slice(qhsd, func(i, j int) bool {
		if qhsd[i].unit < qhsd[j].unit {
			return true
		}
		return false
	})

	rest, total := float64(N), 0
	for _, t := range qhsd {
		tmp := int(rest / t.l)
		rest, total = rest-float64(tmp)*t.l, total+tmp*t.val
	}

	fmt.Println(total)
}

type tuple struct {
	val     int
	l, unit float64
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
