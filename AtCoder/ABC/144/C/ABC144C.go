package main

import (
	"fmt"
	"math"
)

func main() {
	// sc.Split(bufio.ScanWords)
	// N := nextInt()
	var N int64
	fmt.Scanf("%d", &N)

	min := int64(math.MaxInt64)
	for i := int64(2); i*i <= N; i++ {
		if N%i != 0 {
			continue
		}
		mv := int64(i + N/i - 2)
		if mv < min {
			min = mv
		}
	}

	if min == math.MaxInt64 {
		min = N - 1
	}

	fmt.Println(min)
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
