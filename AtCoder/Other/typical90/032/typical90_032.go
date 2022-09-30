package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	a := make([][]int, n)
	for i := range a {
		a[i] = make([]int, n)
		for j := range a[i] {
			a[i][j] = nextInt()
		}
	}

	m := nextInt()
	badrel := make([][]bool, n)
	for i := range badrel {
		badrel[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		x, y := nextInt()-1, nextInt()-1
		badrel[x][y] = true
		badrel[y][x] = true
	}

	order := make([]int, n)
	for i := 0; i < n; i++ {
		order[i] = i
	}

	res := math.MaxInt64
	for {
		tmp := a[order[0]][0]
		for i := 1; i < n; i++ {
			tmp += a[order[i]][i]
			if badrel[order[i-1]][order[i]] {
				tmp = math.MaxInt64
				break
			}
		}
		res = min(res, tmp)

		if !NextPermutation((sort.IntSlice(order))) {
			break
		}
	}

	if res == math.MaxInt64 {
		res = -1
	}
	fmt.Println(res)
}

func NextPermutation(x sort.Interface) bool {
	n := x.Len() - 1
	if n < 1 {
		return false
	}
	j := n - 1
	for ; !x.Less(j, j+1); j-- {
		if j == 0 {
			return false
		}
	}
	l := n
	for !x.Less(j, l) {
		l--
	}
	x.Swap(j, l)
	for k, l := j+1, n; k < l; {
		x.Swap(k, l)
		k++
		l--
	}
	return true
}

func main_naive() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	a := make([][]int, n)
	for i := range a {
		a[i] = make([]int, n)
		for j := range a[i] {
			a[i][j] = nextInt()
		}
	}

	m := nextInt()
	x, y := make([]int, m), make([]int, m)
	for i := range x {
		x[i] = nextInt() - 1
		y[i] = nextInt() - 1
	}

	order := make([]int, n)
	for i := range order {
		order[i] = i
	}
	orders := permutations(order)

	res := math.MaxInt64
	for i := range orders {
		tmp, isFinishable := 0, true
		for j := range orders[i] {
			if j != 0 {
				for k := range x {
					if x[k] == orders[i][j-1] && y[k] == orders[i][j] {
						isFinishable = false
						break
					}
					if x[k] == orders[i][j] && y[k] == orders[i][j-1] {
						isFinishable = false
						break
					}
				}
			}
			if isFinishable {
				tmp += a[orders[i][j]][j]
			}
		}
		if isFinishable {
			res = min(res, tmp)
		}
	}

	if res == math.MaxInt64 {
		res = -1
	}
	fmt.Println(res)
}

func min(a, b int) int {
	if a > b {
		return b
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

// var sc = bufio.NewScanner((os.Stdin))
