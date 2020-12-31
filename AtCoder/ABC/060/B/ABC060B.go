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

// func main() {
func mainABC060B() {
	sc.Split(bufio.ScanWords)
	A, B, C := nextInt(), nextInt(), nextInt()

	fmt.Printf("%s\n", ABC060B(A, B, C))
}

// ABC060B ...
func ABC060B(A, B, C int) string {
	for i := 0; i < B; i++ {
		if A*i%B == C {
			return "YES"
		}
	}

	return "NO"
}
