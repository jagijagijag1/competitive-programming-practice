package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	A, B, K := nextInt(), nextInt(), nextInt()
	nums := map[int]struct{}{}

	for i := A; i < minInt(A+K, B); i++ {
		nums[i] = struct{}{}
	}
	for i := maxInt(B-K+1, A); i <= B; i++ {
		nums[i] = struct{}{}
	}

	keys := make([]int, 0, len(nums))
	for k := range nums {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		fmt.Println(k)
	}
}

func minInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func maxInt(a, b int) int {
	if a < b {
		return b
	}
	return a
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
