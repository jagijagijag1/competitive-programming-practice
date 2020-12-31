package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	N, A, B := nextInt(), nextInt(), nextInt()

	if A > B || (N == 1 && A != B) {
		fmt.Println(0)
		return
	}

	if N == 1 {
		fmt.Println(1)
		return
	}

	max := A + (N-1)*B
	min := (N-1)*A + B

	fmt.Println(max - min + 1)
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
