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
	N, A := nextInt(), []int{}
	for i := 0; i < N; i++ {
		A = append(A, nextInt())
	}
	sort.Ints(A)

	c := make([]int, N)
	for i, a := range A {
		if i == 0 {
			c[i] = a
		} else {
			c[i] = c[i-1] + a
		}
	}

	res := 1
	for i := N - 2; 0 <= i; i-- {
		if A[i+1] <= 2*c[i] {
			res++
		} else {
			break
		}
	}
	fmt.Println(res)

	// N, c := nextInt(), []creature{}
	// for i := 0; i < N; i++ {
	// 	A := nextInt()
	// 	c = append(c, creature{i, A})
	// }

	// res := dfs(c, map[int]struct{}{})
	// fmt.Println(len(res))
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

type creature struct {
	color, size int
}

func dfs(cc []creature, cr map[int]struct{}) map[int]struct{} {
	if len(cc) == 1 {
		cr[cc[0].color] = struct{}{}
		return cr
	}

	for i := 0; i < len(cc)-1; i++ {
		for j := i + 1; j < len(cc); j++ {
			if cc[j].size <= 2*cc[i].size {
				tmp := []creature{}
				for k := range cc {
					if k != i && k != j {
						tmp = append(tmp, cc[k])
					}
				}
				tmp = append(tmp, creature{cc[i].color, cc[i].size + cc[j].size})
				cr = dfs(tmp, cr)
			}
			if cc[i].size <= 2*cc[j].size {
				tmp := []creature{}
				for k := range cc {
					if k != i && k != j {
						tmp = append(tmp, cc[k])
					}
				}
				tmp = append(tmp, creature{cc[j].color, cc[i].size + cc[j].size})
				cr = dfs(tmp, cr)
			}
		}
	}

	return cr
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
