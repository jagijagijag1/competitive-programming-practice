package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var board [][]bool

func main() {
	sc.Split(bufio.ScanWords)
	sx, sy, tx, ty := nextInt(), nextInt(), nextInt(), nextInt()
	dx, dy := tx-sx, ty-sy

	fmt.Print(strings.Repeat("U", dy))
	fmt.Print(strings.Repeat("R", dx))

	fmt.Print(strings.Repeat("D", dy))
	fmt.Print(strings.Repeat("L", dx))

	fmt.Print("L")
	fmt.Print(strings.Repeat("U", dy+1))
	fmt.Print(strings.Repeat("R", dx+1))
	fmt.Print("D")

	fmt.Print("R")
	fmt.Print(strings.Repeat("D", dy+1))
	fmt.Print(strings.Repeat("L", dx+1))
	fmt.Println("U")
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
