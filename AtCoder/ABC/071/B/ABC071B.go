package main

import (
	"fmt"
)

// func main() {
func mainABC071B() {
	// sc.Split(bufio.ScanWords)
	// S := nextLine()
	var S string
	fmt.Scanf("%s", &S)

	fmt.Printf("%s\n", ABC071B(S))
}

// ABC071B ...
func ABC071B(S string) string {
	used := make([]bool, 26)
	for _, s := range S {
		used[s-'a'] = true
	}

	for i := range used {
		if !used[i] {
			return string(rune('a' + i))
		}
	}

	return "None"
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
