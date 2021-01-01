package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	// sc.Split(bufio.ScanWords)
	s, t := nextLine(), nextLine()
	ss, tt := []byte(s), []byte(t)

	sort.Slice(ss, func(i, j int) bool {
		if ss[i] < ss[j] {
			return true
		}
		return false
	})
	sort.Slice(tt, func(i, j int) bool {
		if tt[i] > tt[j] {
			return true
		}
		return false
	})

	if string(ss) < string(tt) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
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
