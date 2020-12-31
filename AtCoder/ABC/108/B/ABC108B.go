package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// func main() {
func mainABC108B() {
	sc.Split(bufio.ScanWords)
	x1, y1, x2, y2 := nextInt(), nextInt(), nextInt(), nextInt()

	x3, y3, x4, y4 := ABC108B(x1, y1, x2, y2)
	fmt.Printf("%d %d %d %d\n", x3, y3, x4, y4)
}

// ABC108B ...
func ABC108B(x1, y1, x2, y2 int) (x3, y3, x4, y4 int) {
	dx, dy := x2-x1, y2-y1
	x3, y3 = x2-dy, y2+dx
	x4, y4 = x3-dx, y3-dy

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
