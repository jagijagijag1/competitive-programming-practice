package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// sc.Split(bufio.ScanWords)
	// N, T := nextInt(), []int{}
	S := nextLine()
	mv := map[byte]int{}
	for i := range S {
		mv[S[i]]++
	}

	if (mv['E'] != 0 && mv['W'] == 0) || (mv['E'] == 0 && mv['W'] != 0) || (mv['N'] != 0 && mv['S'] == 0) || (mv['N'] == 0 && mv['S'] != 0) {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}

var sc = bufio.NewScanner((os.Stdin))

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

// func nextInt() int {
// 	l := nextLine()
// 	i, e := strconv.Atoi(l)
// 	if e != nil {
// 		panic(e)
// 	}
// 	return i
// }
