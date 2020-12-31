package main

import (
	"fmt"
)

// func main() {
func mainPanasonic2020B() {
	// sc.Split(bufio.ScanWords)
	// H, W := nextInt(), nextInt()
	var H, W int64
	fmt.Scanf("%d %d", &H, &W)

	fmt.Printf("%d\n", Panasonic2020B(H, W))
}

// Panasonic2020B ...
func Panasonic2020B(H, W int64) int64 {
	if H == 1 || W == 1 {
		return 1
	}

	size := H * W
	if size%2 == 1 {
		return size/2 + 1
	}

	return size / 2
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
