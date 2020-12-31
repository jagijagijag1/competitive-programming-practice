package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// sc.Split(bufio.ScanWords)
	S := nextLine()

	if S[0] != 'A' || !strings.Contains(S[2:maxInt(3, len(S)-1)], "C") {
		fmt.Printf("WA\n")
		return
	}

	ccnt := 0
	for i := range S {
		if S[i] == 'A' {
			continue
		} else if S[i] == 'C' {
			ccnt++
		} else if S[i] < 'a' || 'z' < S[i] {
			fmt.Printf("WA\n")
			return
		}
	}

	if ccnt != 1 {
		fmt.Printf("WA\n")
		return
	}

	fmt.Printf("AC\n")
}

func maxInt(a, b int) int {
	if a < b {
		return b
	}
	return a
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
