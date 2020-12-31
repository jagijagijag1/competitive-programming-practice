package main

import (
	"fmt"
)

func main() {
	// sc.Split(bufio.ScanWords)
	// N := nextInt()
	var S string
	fmt.Scanf("%s", &S)

	res, curr, last := 0, "", ""
	for _, s := range S {
		ss := string(s)
		curr += ss
		if curr != last {
			res++
			curr, last = "", curr
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
