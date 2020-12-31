package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// func main() {
func mainABC149C() {
	sc.Split(bufio.ScanWords)
	X := nextInt()

	fmt.Printf("%d\n", ABC149C(X))
}

// ABC149C ...
func ABC149C(X int) int {
	for i := X; i < 1000000; i++ {
		if isPrime(i) {
			return i
		}
	}
	return 0
}

func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
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
