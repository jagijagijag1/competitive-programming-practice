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
	n, arr, x, y := nextInt(), []int{}, []int{}, []int{}
	for i := 0; i < n; i++ {
		arr = append(arr, i)
		x = append(x, nextInt())
		y = append(y, nextInt())
	}

	perm := permutations(arr)
	sum := 0.0
	for _, p := range perm {
		lx, ly, d := x[p[0]], y[p[0]], 0.0
		for i := 1; i < n; i++ {
			nx, ny := x[p[i]], y[p[i]]
			dx, dy := nx-lx, ny-ly
			td := math.Sqrt(float64(dx*dx) + float64(dy*dy))
			d += td
			lx, ly = nx, ny
		}
		sum += d
	}
	fmt.Println(sum / float64(len(perm)))
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

// const (
// 	initialBufSize = 10000
// 	maxBufSize     = 1000000
// )

// var (
// 	sc *bufio.Scanner = func() *bufio.Scanner {
// 		sc := bufio.NewScanner(os.Stdin)
// 		buf := make([]byte, initialBufSize)
// 		sc.Buffer(buf, maxBufSize)
// 		return sc
// 	}()
// )
