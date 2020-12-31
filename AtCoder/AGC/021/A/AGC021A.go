package main

import (
	"fmt"
)

// func main() {
func mainAGC021A() {
	// sc.Split(bufio.ScanWords)
	var N string
	fmt.Scanf("%s", &N)

	fmt.Printf("%d\n", AGC021A(N))
}

// AGC021A ...
func AGC021A(N string) (res int) {
	all9 := true
	for i := 1; i < len(N); i++ {
		if N[i] != '9' {
			all9 = false
			break
		}
	}

	if all9 {
		res = int(N[0]-'0') + 9*(len(N)-1)
	} else {
		res = int(N[0]-'0') - 1 + 9*(len(N)-1)
	}

	return
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
