package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// func main() {
func mainABC150C() {
	sc.Split(bufio.ScanWords)
	N, P, Q := nextInt(), []int{}, []int{}
	//P, Q = strings.Replace(P, " ", "", -1), strings.Replace(Q, " ", "", -1)
	for i := 0; i < N; i++ {
		P = append(P, nextInt())
	}
	for i := 0; i < N; i++ {
		Q = append(Q, nextInt())
	}

	fmt.Printf("%d\n", ABC150C(N, P, Q))
}

// ABC150C ...
func ABC150C(N int, P, Q []int) int {
	pstr, qstr := "", ""
	for i := range P {
		pstr += string(P[i])
		qstr += string(Q[i])
	}

	prem := permutations(P)
	ppstr := []string{}
	for i := range prem {
		tmpp := ""
		for j := range prem[i] {
			tmpp += string(prem[i][j])
		}
		ppstr = append(ppstr, tmpp)
	}

	a, b := -1, -1
	sort.Strings(ppstr)
	for i := range ppstr {
		if ppstr[i] == pstr {
			a = i
		}
		if ppstr[i] == qstr {
			b = i
		}
	}

	return absInt(a - b)
}

func absInt(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
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
