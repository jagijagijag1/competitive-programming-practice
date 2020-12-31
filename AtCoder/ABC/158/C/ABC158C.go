package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	A, B := nextInt(), nextInt()

	for i := 0; i < 2000; i++ {
		aa, bb := i*8/100, i*10/100
		if A == aa && B == bb {
			fmt.Println(i)
			return
		}
	}

	fmt.Println(-1)
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
