package main

import (
	"fmt"
)

func main() {
	// sc.Split(bufio.ScanWords)
	// A, B, C, K := nextInt(), nextInt(), nextInt(), nextInt()
	var A, B, C, K int64
	fmt.Scanf("%d %d %d %d", &A, &B, &C, &K)

	var res int64
	if K%2 == 1 {
		res = B - A
	} else {
		res = A - B
	}
	if absInt(res) > 1000000000000000000 {
		fmt.Printf("Unfair\n")
	}
	fmt.Printf("%d\n", res)
}

func absInt(a int64) int64 {
	if a < 0 {
		return -a
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
