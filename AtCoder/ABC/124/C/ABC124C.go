package main

import (
	"fmt"
)

// func main() {
func mainABC124C() {
	// sc.Split(bufio.ScanWords)
	// S := nextLine()
	var S string
	fmt.Scanf("%s", &S)
	fmt.Printf("%d\n", ABC124C(S))
}

// ABC124C ...
func ABC124C(S string) (res int) {
	last := S[0]
	for i := 1; i < len(S); i++ {
		if last == S[i] {
			res++
			if S[i] == '0' {
				last = '1'
			} else {
				last = '0'
			}
		} else {
			last = S[i]
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
