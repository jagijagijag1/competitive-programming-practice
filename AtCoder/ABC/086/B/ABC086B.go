package main

import (
	"fmt"
	"math"
	"strconv"
)

// func main() {
func mainARC086B() {
	// sc.Split(bufio.ScanWords)
	var a, b string
	fmt.Scanf("%s %s", &a, &b)
	ab, _ := strconv.Atoi(a + b)

	fmt.Printf("%s\n", ABC086B(ab))
}

// ABC086B ...
func ABC086B(ab int) string {
	s := math.Sqrt(float64(ab))
	if s == float64(int(s)) {
		return "Yes"
	}

	return "No"
}

// var sc = bufio.NewScanner((os.Stdin))
//
// func nextLine() string {
// 	sc.Scan()
// 	return sc.Text()
// }
//
// func nextInt() int {
// 	l := nextLine()
// 	i, e := strconv.Atoi(l)
// 	if e != nil {
// 		panic(e)
// 	}
// 	return i
// }
