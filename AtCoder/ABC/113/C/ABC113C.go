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

//func main() {
func mainABC113C() {
	sc.Split(bufio.ScanWords)
	N, M := nextInt(), nextInt()
	P, Y := []int{}, []int{}
	for i := 0; i < M; i++ {
		P = append(P, nextInt())
		Y = append(Y, nextInt())
	}

	//fmt.Printf("%s\n", ABC113C(N, M, P, Y))
	ABC113C(N, M, P, Y)
}

// City ...
type City struct {
	i   int
	p   int
	y   int
	cid int
}

// ABC113C ...
func ABC113C(N, M int, P, Y []int) string {
	cties := []City{}
	for i := 0; i < M; i++ {
		cties = append(cties, City{i, P[i], Y[i], -1})
	}

	sort.Slice(cties, func(i, j int) bool {
		if cties[i].p < cties[j].p {
			return true
		} else if cties[i].p == cties[j].p {
			if cties[i].y < cties[j].y {
				return true
			}
		}
		return false
	})

	ccid, lastPid := 0, 0
	for i := 0; i < len(cties); i++ {
		if cties[i].p != lastPid {
			lastPid = cties[i].p
			ccid = 1
			cties[i].cid = ccid
		} else {
			ccid++
			cties[i].cid = ccid
		}
	}

	sort.Slice(cties, func(i, j int) bool {
		if cties[i].i < cties[j].i {
			return true
		}
		return false
	})

	for _, c := range cties {
		fmt.Printf("%06d%06d\n", c.p, c.cid)
	}

	return ""
}
