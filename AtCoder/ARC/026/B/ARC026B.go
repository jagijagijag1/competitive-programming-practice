package main

import (
	"fmt"
)

// func main() {
func mainARC026B() {
	// sc.Split(bufio.ScanWords)
	// N := nextInt()
	var N int64
	fmt.Scanf("%d", &N)

	fmt.Printf("%s\n", ARC026B(N))
}

// ARC026B ...
func ARC026B(N int64) string {
	sum := int64(0)
	for i := int64(1); i*i <= N; i++ {
		if N%i == 0 {
			sum += i
			tmp := N / i
			if tmp != i {
				sum += tmp
			}
		}
	}

	sum -= N

	if sum == N {
		return "Perfect"
	} else if sum < N {
		return "Deficient"
	}

	return "Abundant"
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
