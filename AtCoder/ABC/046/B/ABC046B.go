package main

import (
	"bufio"
	"fmt"
	"math"
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
func mainABC046B() {
	sc.Split(bufio.ScanWords)
	N, K := nextInt(), nextInt()

	fmt.Printf("%d\n", ABC046B(N, K))
}

// ABC046B ...
func ABC046B(N, K int) int {
	return K * int(math.Pow(float64(K-1), float64(N-1)))
}
