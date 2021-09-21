package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	x := nextInt()
	fmt.Println(lcm(x, 360) / x)
}

func gcd(x, y int) int {
	for y > 0 {
		x, y = y, x%y
	}
	return x
}

func lcm(x, y int) int {
	return x / gcd(x, y) * y
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
