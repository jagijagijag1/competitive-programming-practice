package main

import (
	"fmt"
	"strings"
)

func main() {
	// sc.Split(bufio.ScanWords)
	// N := nextInt()
	var s string
	fmt.Scanf("%s", &s)
	fmt.Println(strings.LastIndex(s, "Z") - strings.Index(s, "A") + 1)
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
