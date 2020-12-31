package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	c := make([]int, 9)
	for i := 0; i < 9; i++ {
		c[i] = nextInt()
	}

	a1b1, a2b2, a3b3 := split(c[0]), split(c[4]), split(c[8])
	for _, t1 := range a1b1 {
		for _, t2 := range a2b2 {
			for _, t3 := range a3b3 {
				if t1.a+t2.b == c[1] && t1.a+t3.b == c[2] &&
					t2.a+t1.b == c[3] && t2.a+t3.b == c[5] &&
					t3.a+t1.b == c[6] && t3.a+t2.b == c[7] {
					fmt.Println("Yes")
					return
				}
			}
		}
	}

	fmt.Println("No")
}

type tuple struct {
	a, b int
}

func split(c int) (res []tuple) {
	if c == 0 {
		return []tuple{{0, 0}}
	}

	for i := 0; i <= c; i++ {
		res = append(res, tuple{i, c - i})
	}
	return
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
