package main

import (
	"fmt"
)

func main() {
	// sc.Split(bufio.ScanWords)
	// s := nextLine()
	var s string
	fmt.Scanf("%s", &s)

	res := len(s) * (len(s) - 1)
	for i := range s {
		if s[i] == 'U' {
			res += i
		} else {
			res += len(s) - 1 - i
		}
	}

	fmt.Println(res)
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
