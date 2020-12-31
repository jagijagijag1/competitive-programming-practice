package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	N := nextInt()

	L := make([]int, N+1)
	L[0], L[1] = 2, 1
	for i := 2; i <= N; i++ {
		L[i] = L[i-1] + L[i-2]
	}

	fmt.Println(L[N])
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
