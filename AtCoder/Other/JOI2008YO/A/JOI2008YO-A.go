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
func mainJOI2008YOA() {
	sc.Split(bufio.ScanWords)
	cost := nextInt()

	fmt.Printf("%d\n", JOI2008YOA(cost))
}

// JOI2008YOA ...
func JOI2008YOA(cost int) (res int) {
	change := 1000 - cost
	coins := []int{500, 100, 50, 10, 5, 1}

	for _, coin := range coins {
		res += change / coin
		change = change % coin
	}

	return
}
