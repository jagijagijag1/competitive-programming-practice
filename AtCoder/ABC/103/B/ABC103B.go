package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// sc.Split(bufio.ScanWords)
	// N := nextInt()
	S, T := nextLine(), nextLine()
	if S == T {
		fmt.Println("Yes")
		return
	}

	SS := S
	for i := 0; i < len(S); i++ {
		SS = string(SS[len(SS)-1]) + SS[0:len(SS)-1]
		if T == SS {
			fmt.Println("Yes")
			return
		}
	}

	fmt.Println("No")
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
