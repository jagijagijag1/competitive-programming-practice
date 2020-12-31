package main

import (
	"fmt"
)

// func main() {
func mainABC100C() {
	// sc.Split(bufio.ScanWords)
	// N, v := nextInt(), []int{}
	var N int
	fmt.Scanf("%d", &N)

	a := make([]int64, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &a[i])
	}

	fmt.Printf("%d\n", ABC100C(a))
}

// ABC100C ...
func ABC100C(a []int64) (res int) {
	for _, aa := range a {
		for {
			if aa%2 == 0 {
				res++
				aa = aa / 2
			} else {
				break
			}
		}
	}

	return
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
