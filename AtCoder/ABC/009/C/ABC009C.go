package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// func main() {
func mainABC009C() {
	sc.Split(bufio.ScanWords)
	_, K, S := nextInt(), nextInt(), nextLine()

	fmt.Printf("%s\n", ABC009C(K, S))
}

// ABC009C ...
func ABC009C(K int, S string) string {
	s, charCount := []rune{}, make([]int, 26)
	for _, c := range S {
		s = append(s, c)
		charCount[c-'a']++
	}

	T := []rune{}
	for i := 0; i < len(S); i++ {
		for j, cnt := range charCount {
			if cnt == 0 {
				continue
			}

			tmpT := append(T, rune(j+'a'))
			headDiff := distOfRunes(s, tmpT)

			tmpCharCount := make([]int, 26)
			copy(tmpCharCount, charCount)
			tmpCharCount[j]--
			tailDiff := distByCharCount(s[i+1:], tmpCharCount)

			if headDiff+tailDiff <= K {
				T = append(T, rune(j+'a'))
				charCount[j]--
				break
			}
		}
	}

	return string(T)
}

// ABC009CWA2 ...
func ABC009CWA2(K int, S string) (T string) {
	var s, rest []rune
	for _, c := range S {
		s = append(s, c)
		rest = append(rest, c)
	}
	sort.Slice(rest, func(i, j int) bool {
		return rest[i] < rest[j]
	})

	for i := 0; i < len(s); i++ {
		for _, c := range rest {
			rest = rest[1:]
			T = T + string(c)

			tmpT := swap(s, i, c)
			mvCntT := distOfRunes([]rune(S[:i]), []rune(tmpT[:i]))
			mvCntRest := distOfRunes([]rune(S[i:]), tmpT[i:])

			if mvCntT+mvCntRest == K {
				return string(s)
			} else if mvCntT+mvCntRest > K {
				continue
			} else {
				break
			}
		}
	}

	return string(s)
}

// ABC009CWA ...
func ABC009CWA(K int, S string) (T string) {
	var s, r []rune
	for _, c := range S {
		s = append(s, c)
		r = append(r, c)
	}
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})

	head := ""
	for i := 0; i < len(s); i++ {
		for _, c := range r {
			ns := swap(s, i, c)

			if distOfRunes([]rune(S), ns) <= K {
				head = head + string(c)
				T = head + string(s[i+1:])
				r = r[1:]
				break
			} else {
				return
			}
		}
	}

	return
}

// swap s[i] and charcter c appeared first in s
func swap(s []rune, i int, c rune) (res []rune) {
	var idx int
	for i, r := range s {
		if r == c {
			idx = i
			break
		}
	}

	res = make([]rune, len(s))
	copy(res, s)
	res[i], res[idx] = res[idx], res[i]

	return
}

func distOfRunes(s, t []rune) (res int) {
	for i := 0; i < len(t); i++ {
		if s[i] != t[i] {
			res++
		}
	}

	return
}

func distByCharCount(s []rune, charCount []int) (res int) {
	for i := 0; i < len(s); i++ {
		if charCount[s[i]-'a'] > 0 {
			charCount[s[i]-'a']--
		} else {
			res++
		}
	}

	return
}

var sc = bufio.NewScanner((os.Stdin))

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	l := nextLine()
	i, e := strconv.Atoi(l)
	if e != nil {
		panic(e)
	}
	return i
}
