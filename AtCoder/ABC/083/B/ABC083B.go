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
func mainABC083B() {
	sc.Split(bufio.ScanWords)
	N, A, B := nextInt(), nextInt(), nextInt()

	fmt.Printf("%d\n", ABC083B(N, A, B))
}

// ABC083B ...
func ABC083B(N, A, B int) int {
	res := 0

	for n := 1; n <= N; n++ {
		s := 0
		for tmp := n; tmp != 0; tmp /= 10 {
			s += tmp % 10
			if s > B {
				break
			}
		}

		if A <= s && s <= B {
			res += n
		}
	}

	return res
}
