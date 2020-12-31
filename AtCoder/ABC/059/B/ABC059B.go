package main

import (
	"fmt"
)

func main() {
	// sc.Split(bufio.ScanWords)
	// N := nextInt()
	var A, B string
	fmt.Scanf("%s\n%s", &A, &B)

	if len(A) > len(B) {
		fmt.Println("GREATER")
		return
	} else if len(A) < len(B) {
		fmt.Println("LESS")
		return
	}

	for i := range A {
		if A[i] == B[i] {
			continue
		} else if A[i] > B[i] {
			fmt.Println("GREATER")
			return
		} else {
			fmt.Println("LESS")
			return
		}
	}

	fmt.Println("EQUAL")
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
