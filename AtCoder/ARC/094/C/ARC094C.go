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
	A, B, C := nextInt(), nextInt(), nextInt()
	abc := []int{A, B, C}
	sort.Ints(abc)

	res := 0

	d1, d2 := abc[2]-abc[0], abc[2]-abc[1]
	if d1%2 == 1 && d2%2 == 0 {
		abc[2]++
		abc[1]++
		res++
	} else if d1%2 == 0 && d2%2 == 1 {
		abc[2]++
		abc[0]++
		res++
	} else if d1%2 == 1 && d2%2 == 1 {
		abc[1]++
		abc[0]++
		res++
	}

	res += (abc[2] - abc[0]) / 2
	res += (abc[2] - abc[1]) / 2

	fmt.Println(res)
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
