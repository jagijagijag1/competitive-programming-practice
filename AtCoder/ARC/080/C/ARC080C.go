package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	N, mod4, odd := nextInt(), []int{}, []int{}
	for i := 0; i < N; i++ {
		aa := nextInt()
		if aa%4 == 0 {
			mod4 = append(mod4, aa)
		} else if aa%2 != 0 {
			odd = append(odd, aa)
		}
	}

	if len(odd) <= len(mod4) || (len(odd) == len(mod4)+1 && len(odd)+len(mod4) == N) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
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
