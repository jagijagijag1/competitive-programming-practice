package main

import (
	"fmt"
)

func main() {
	// sc.Split(bufio.ScanWords)
	// N := nextInt()
	var S string
	fmt.Scanf("%s", &S)

	cnt0, cnt1 := 0, 0
	for _, s := range S {
		if s == '0' {
			cnt0++
		} else {
			cnt1++
		}
	}

	fmt.Println(2 * min(cnt0, cnt1))
}

func min(a, b int) int {
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
