package main

import (
	"fmt"
)

// func main() {
func mainABC158B() {
	// sc.Split(bufio.ScanWords)
	// S := nextLine()
	var N, A, B uint64
	fmt.Scanf("%d %d %d", &N, &A, &B)

	fmt.Printf("%d\n", ABC158B(N, A, B))
}

// ABC158B ...
func ABC158B(N, A, B uint64) uint64 {
	rep, mod := N/(A+B), N%(A+B)
	res := rep*A + minUint64(A, mod)

	return res
}

func minUint64(a, b uint64) uint64 {
	if a > b {
		return b
	}
	return a
}

// var sc = bufio.NewScanner((os.Stdin))

// func nextLine() string {
// 	sc.Scan()
// 	return sc.Text()
// }

// func nextInt() int {
// 	l := nextLine()
// 	i, e := strconv.Atoi(l)
// 	if e != nil {
// 		panic(e)
// 	}
// 	return i
// }
