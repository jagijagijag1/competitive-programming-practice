package main

import (
	"fmt"
)

// func main() {
func mainAGC002A() {
	// sc.Split(bufio.ScanWords)
	// a, b := nextInt(), nextInt()
	var a, b int64
	fmt.Scanf("%d %d", &a, &b)

	fmt.Printf("%s\n", AGC002A(a, b))
}

// AGC002A ...
func AGC002A(a, b int64) string {
	if a*b <= 0 {
		return "Zero"
	}

	if 0 < a {
		return "Positive"
	}

	if (b-a+1)%2 == 0 {
		return "Positive"
	}

	return "Negative"
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
