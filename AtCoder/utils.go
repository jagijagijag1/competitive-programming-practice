package main

import (
	"bufio"
	"container/heap"
	"math"
	"os"
	"sort"
	"strconv"
)

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

func nextUint64() uint64 {
	l := nextLine()
	i, e := strconv.ParseUint(l, 10, 64)
	if e != nil {
		panic(e)
	}
	return i
}

func nextFloat64() float64 {
	l := nextLine()
	i, e := strconv.ParseFloat(l, 64)
	if e != nil {
		panic(e)
	}
	return i
}

type stack []string

func (s stack) Push(v string) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, string) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func permutationsRune(arr []rune) []string {
	var helper func([]rune, int)
	res := []string{}

	helper = func(arr []rune, n int) {
		if n == 1 {
			tmp := make([]rune, len(arr))
			copy(tmp, arr)
			res = append(res, string(tmp))
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func permutationsStr(arr []string) [][]string {
	var helper func([]string, int)
	res := [][]string{}

	helper = func(arr []string, n int) {
		if n == 1 {
			tmp := make([]string, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

// https://go.dev/play/p/ljft9xhOEn
// NextPermutation generates the next permutation of the
// sortable collection x in lexical order.  It returns false
// if the permutations are exhausted.
//
// Knuth, Donald (2011), "Section 7.2.1.2: Generating All Permutations",
// The Art of Computer Programming, volume 4A.
func NextPermutation(x sort.Interface) bool {
	n := x.Len() - 1
	if n < 1 {
		return false
	}
	j := n - 1
	for ; !x.Less(j, j+1); j-- {
		if j == 0 {
			return false
		}
	}
	l := n
	for !x.Less(j, l) {
		l--
	}
	x.Swap(j, l)
	for k, l := j+1, n; k < l; {
		x.Swap(k, l)
		k++
		l--
	}
	return true
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func minIntInSlice(v []int) (res int) {
	res = math.MaxInt32
	for i, e := range v {
		if i == 0 || res > e {
			res = e
		}
	}
	return
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func maxIntInSlice(v []int) (res int) {
	res = math.MinInt32
	for i, e := range v {
		if i == 0 || res < e {
			res = e
		}
	}
	return
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

// a^nをpで割った余り
func modpow(a, n, p int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return a % p
	}
	if n%2 == 1 {
		return (a * modpow(a, n-1, p)) % p
	}
	t := modpow(a, n/2, p)
	return (t * t) % p
}

// 大きい数字に対するa/b % p
func div(a, b, p int) int {
	return a * modpow(b, p-2, p) % p
}

// nCaをpで割った余り (aば小さい場合に有効)
func modchoose(n, a, p int) int {
	x, y := 1, 1

	for i := 0; i < a; i++ {
		x *= n - i
		x %= p
		y *= i + 1
		y %= p
	}

	inv := modpow(y, p-2, p)
	res := x * inv
	return res % p
}

func intSliceEq(a, b []int) bool {
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func stringSliceEq(a, b []string) bool {
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func gcds(a, b int, integers ...int) int {
	res := gcd(a, b)
	for i := 0; i < len(integers); i++ {
		res = gcd(res, integers[i])
	}
	return res
}

// コメント版の方がa*bで数値が大きすぎになる可能性あるっぽい
// func lcm(a, b int) int {
// 	return a * b / gcd(a, b)
// }
func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

func lcms(a, b int, integers ...int) int {
	result := a * b / gcd(a, b)

	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}

	return result
}

type intHeap []int

func (h intHeap) Len() int           { return len(h) }
func (h intHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h intHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type float64Heap []float64

func (h float64Heap) Len() int           { return len(h) }
func (h float64Heap) Less(i, j int) bool { return h[i] < h[j] }
func (h float64Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *float64Heap) Push(x interface{}) {
	*h = append(*h, x.(float64))
}
func (h *float64Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

/*** deque ***/ // dq := Deque()
type deque_type struct {
	Dat []int
	L   int
	R   int
}

func Deque() *deque_type {
	return &deque_type{Dat: make([]int, 200001), L: 100000, R: 100001}
}
func (dq *deque_type) PushBack(val int) {
	dq.Dat[dq.R] = val
	dq.R++
}
func (dq *deque_type) PushFront(val int) {
	dq.Dat[dq.L] = val
	dq.L--
}
func (dq *deque_type) RemoveBack() {
	dq.R--
}
func (dq *deque_type) RemoveFront() {
	dq.L++
}
func (dq *deque_type) GetBack() int {
	return dq.Dat[dq.R-1]
}
func (dq *deque_type) GetFront() int {
	return dq.Dat[dq.L+1]
}
func (dq *deque_type) GetIth(idx int) int {
	return dq.Dat[dq.L+idx+1]
}
func (dq *deque_type) Len() int {
	return dq.R - dq.L - 1
}

func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func primeFactorization(n int) (pfs map[int]int) {
	pfs = make(map[int]int)

	for n%2 == 0 {
		if _, ok := pfs[2]; ok {
			pfs[2]++
		} else {
			pfs[2] = 1
		}
		n = n / 2
	}

	for i := 3; i*i <= n; i = i + 2 {
		for n%i == 0 {
			if _, ok := pfs[i]; ok {
				pfs[i]++
			} else {
				pfs[i] = 1
			}
			n = n / i
		}
	}

	if n > 2 {
		pfs[n] = 1
	}

	return
}

func divisor(n int) (res []int) {
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			res = append(res, i)
			if i*i != n {
				res = append(res, n/i)
			}
		}
	}
	sort.Ints(res)
	return
}

type tuple struct {
	t1, t2 int
}

func reverse(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}

// nCa
func choose(n, a int) int {
	x, y := 1, 1

	for i := 0; i <= a; i++ {
		x *= n - i
		y *= i + 1
	}

	return x / y
}

// from https://qiita.com/hiroykam/items/d182ae2a41dedb663380#golang
func permutation(n int, k int) int {
	v := 1
	if 0 < k && k <= n {
		for i := 0; i < k; i++ {
			v *= (n - i)
		}
	} else if k > n {
		v = 0
	}
	return v
}

func factorial(n int) int {
	return permutation(n, n-1)
}

func combination(n int, k int) int {
	return permutation(n, k) / factorial(k)
}

func homogeneous(n int, k int) int {
	return combination(n+k-1, k)
}

func contains(s []int, e int) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}

// 10進数のNをk進数の文字列に変換
func convertNtoBaseK(n, k int) (res string) {
	tmp := n
	for digitNum := 1; ; {
		remainder := tmp % k
		if remainder >= 10 {
			res = string('A'+(remainder-10)) + res
		} else {
			res = string('0'+remainder) + res
		}
		tmp = tmp / k
		if tmp == 0 {
			break
		}
		digitNum *= 10
	}
	return
}

// 桁数: digits
func d(N int) int {
	return len(strconv.Itoa(N))
}

// 数字和,　各桁の和, digit sum
func dsum(n int) (res int) {
	for 0 < n {
		res += n % 10
		n /= 10
	}
	return
}

// グラフ構造 + 探索アルゴリズム
type Graph struct {
	edges [][]Edge
}

type Edge struct {
	to, weight int
}

func NewGraph(n int) *Graph {
	return &(Graph{make([][]Edge, n)})
}

func (g *Graph) AddEdge(s, t int) {
	g.AddWeightedEdge(s, t, 1)
}

func (g *Graph) AddWeightedEdge(s, t, w int) {
	g.edges[s] = append(g.edges[s], Edge{t, w})
}

type DijkstraNode struct {
	node, cost int
}
type DijkstraPriorityQueue []*DijkstraNode

func (pq DijkstraPriorityQueue) Len() int            { return len(pq) }
func (pq DijkstraPriorityQueue) Less(i, j int) bool  { return pq[i].cost < pq[j].cost }
func (pq DijkstraPriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *DijkstraPriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(*DijkstraNode)) }
func (pq *DijkstraPriorityQueue) Pop() interface{} {
	o := *pq
	n := len(o) - 1
	item := o[n]
	*pq = o[0:n]
	return item
}

// calc min cost from s to t
func (g *Graph) dijkstra(s, t int) int {
	n := len(g.edges)
	pq := make(DijkstraPriorityQueue, 0)
	cost := make([]int, n)
	for i := 0; i < n; i++ {
		var c int
		if i == s {
			c = 0
		} else {
			c = int(1e18)
		}
		cost[i] = c
		heap.Push(&pq, &DijkstraNode{i, c})
	}

	for pq.Len() > 0 {
		u := heap.Pop(&pq).(*DijkstraNode)
		if u.node == t {
			break
		}
		for i := 0; i < len(g.edges[u.node]); i++ {
			v := g.edges[u.node][i]
			c := cost[u.node] + v.weight
			if cost[v.to] > c {
				cost[v.to] = c
				heap.Push(&pq, &DijkstraNode{v.to, c})
			}
		}
	}

	return cost[t]
}

// calc min cost from s to all nodes
func (g *Graph) dijkstraAll(s int) []int {
	n := len(g.edges)
	pq := make(DijkstraPriorityQueue, 0)
	cost := make([]int, n)
	for i := 0; i < n; i++ {
		var c int
		if i == s {
			c = 0
		} else {
			c = int(1e18)
		}
		cost[i] = c
		heap.Push(&pq, &DijkstraNode{i, c})
	}

	for pq.Len() > 0 {
		u := heap.Pop(&pq).(*DijkstraNode)
		for i := 0; i < len(g.edges[u.node]); i++ {
			v := g.edges[u.node][i]
			c := cost[u.node] + v.weight
			if cost[v.to] > c {
				cost[v.to] = c
				heap.Push(&pq, &DijkstraNode{v.to, c})
			}
		}
	}

	return cost
}

// calc min cost from each node to all nodes
func (g *Graph) WarshallFloyd() [][]int {
	n := len(g.edges)
	d := make([][]int, n)
	for i := 0; i < n; i++ {
		d[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == j {
				d[i][j] = 0
			} else {
				d[i][j] = int(1e18)
			}
		}
		for j := 0; j < len(g.edges[i]); j++ {
			k := g.edges[i][j]
			d[i][k.to] = k.weight
		}
	}

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				d[i][j] = min(d[i][j], d[i][k]+d[k][j])
			}
		}
	}

	return d
}
