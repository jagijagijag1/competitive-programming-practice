package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	X, Y := nextInt(), nextInt()
	mod := 1000000007

	if (X+Y)%3 != 0 {
		fmt.Println(0)
		return
	}

	n := (X + Y) / 3
	X, Y = X-n, Y-n
	if X < 0 || Y < 0 {
		fmt.Println(0)
		return
	}

	res := modchoose(X+Y, X, mod)
	for res < 0 {
		res += mod
	}
	fmt.Println(res)
}

func modchoose(n, a, p int) int {
	x, y := 1, 1

	for i := 0; i < a; i++ {
		x *= n - i
		x %= p
		y *= i + 1
		y %= p
	}

	inv := modpow(y, p-2, p)
	res := x * inv
	return res % p
}

func modpow(a, n, p int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return a % p
	}
	if n%2 == 1 {
		return (a * modpow(a, n-1, p)) % p
	}
	t := modpow(a, n/2, p)
	return (t * t) % p
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

type tuple struct {
	x, y int
}

func mainTLE() {
	sc.Split(bufio.ScanWords)
	X, Y := nextInt(), nextInt()

	board := make([][]int, Y+1)
	for i := range board {
		board[i] = make([]int, X+1)
	}

	q := []tuple{{0, 0}}
	for 0 < len(q) {
		h := q[0]
		q = q[1:]

		next := []tuple{{h.x + 1, h.y + 2}, {h.x + 2, h.y + 1}}
		for _, n := range next {
			if n.x <= X && n.y <= Y {
				q = append(q, n)
				board[n.y][n.x]++
			}
		}
	}
	fmt.Println(board[Y][X])
}
