package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	// sc.Split(bufio.ScanWords)
	// N := nextInt()
	var N int
	fmt.Scanf("%d", &N)
	a := make([]int, 3*N)

	sc.Split(bufio.ScanWords)
	for i := range a {
		// fmt.Scanf("%d", &a[i])
		a[i] = nextInt()
	}

	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	var res int
	for i := 0; i < N; i++ {
		res += a[2*i+1]
	}
	fmt.Printf("%d\n", res)
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
