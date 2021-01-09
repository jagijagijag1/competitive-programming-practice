package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	K, A, B, n := nextInt(), nextInt(), nextInt(), 1

	if B <= A+1 {
		fmt.Println(K + 1)
		return
	}

	for i := 0; i < K; i++ {
		if A <= n && i < K-1 {
			i++
			n += B - A
		} else {
			n++
		}
	}

	fmt.Println(n)
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
