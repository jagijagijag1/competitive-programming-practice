package main

import (
	"bufio"
	"fmt"
	"os"
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
func mainABC116C() {
	sc.Split(bufio.ScanWords)
	N, H := nextInt(), []int{}
	for i := 0; i < N; i++ {
		H = append(H, nextInt())
	}

	fmt.Printf("%d\n", ABC116C(N, H))
}

// ABC116C ...
func ABC116C(N int, H []int) int {
	res, active := 0, 0

	for i := 0; i < N; i++ {
		if active >= H[i] {
			active = H[i]
		} else {
			res += H[i] - active
			active = H[i]
		}
	}

	return res
}

// ABC116CNaive ...
func ABC116CNaive(N int, H []int) int {
	res := 0

	c := []int{}
	for i := 0; i < len(H); i++ {
		c = append(c, 0)
	}

	l := 0
	for l != N-1 || c[N-1] < H[N-1] {
		did := false
		for r := l; r < len(H); r++ {
			if H[r] > c[r] {
				c[r]++
				did = true
			} else {
				break
			}
		}
		if did {
			res++
		}

		for i := l; i < len(H); i++ {
			if H[i] > c[i] || i == len(H)-1 {
				l = i
				break
			}
		}
	}

	return res
}
