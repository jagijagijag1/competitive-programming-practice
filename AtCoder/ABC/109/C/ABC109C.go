package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	N, X, dist := nextInt(), nextInt(), []int{}

	for i := 0; i < N; i++ {
		dist = append(dist, abs(X-nextInt()))
	}

	if N == 1 {
		fmt.Println(dist[0])
		return
	}

	res := gcd(dist[0], dist[1])
	for i := 2; i < N; i++ {
		res = gcd(res, dist[i])
	}
	fmt.Println(res)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
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
