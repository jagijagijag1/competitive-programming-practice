package main

import (
	"fmt"
)

func main() {
	// sc.Split(bufio.ScanWords)
	// N := nextInt()
	var S string
	fmt.Scanf("%s", &S)
	var K int64
	fmt.Scanf("%d", &K)

	for i := int64(0); i < K; i++ {
		if S[i] != '1' {
			fmt.Println(S[i] - '0')
			return
		}
	}

	fmt.Println("1")
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
