package main

import (
	"fmt"
)

func main() {
	// sc.Split(bufio.ScanWords)
	// N := nextInt()
	var N, A, B int64
	fmt.Scanf("%d %d %d", &N, &A, &B)

	dist := B - A
	if dist%2 == 0 {
		fmt.Println(dist / 2)
	} else {
		fmt.Println(minInt(A-1, N-B) + 1 + (dist-1)/2)
	}
}

func minInt(a, b int64) int64 {
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
