package main

import (
	"fmt"
)

// func main() {
func mainABC161C() {
	// sc.Split(bufio.ScanWords)
	// N, K := nextInt(), nextInt()
	var N, K int64
	fmt.Scanf("%d %d\n", &N, &K)

	fmt.Printf("%d\n", ABC161C(N, K))
}

// ABC161C ...
func ABC161C(N, K int64) (res int64) {
	last := N - (K * (N / K))
	for {
		tmp := absUint64(last - K)
		if tmp >= last {
			return last
		}
		last = tmp
	}
}

func absUint64(a int64) int64 {
	if a < 0 {
		return -a
	}
	return a
}

// var sc = bufio.NewScanner((os.Stdin))
//
// func nextLine() string {
// 	sc.Scan()
// 	return sc.Text()
// }
//
// func nextInt() int {
// 	l := nextLine()
// 	i, e := strconv.Atoi(l)
// 	if e != nil {
// 		panic(e)
// 	}
// 	return i
// }
