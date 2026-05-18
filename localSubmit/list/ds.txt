package main

type Stack struct {
	data []int
}

func NewStack() *Stack {
	return &Stack{data: make([]int, 0)}
}

func (s *Stack) Push(x int) {
	s.data = append(s.data, x)
}

func (s *Stack) Pop() int {
	n := len(s.data)
	x := s.data[n-1]
	s.data = s.data[:n-1]
	return x
}

func (s *Stack) Top() int {
	return s.data[len(s.data)-1]
}

func (s *Stack) Empty() bool {
	return len(s.data) == 0
}

func (s *Stack) Len() int {
	return len(s.data)
}

func (s *Stack) Clear() {
	s.data = s.data[:0]
}

type Queue struct {
	data []int
	head int
}

func NewQueue(cap int) *Queue {
	return &Queue{data: make([]int, 0, cap)}
}

func (q *Queue) Push(x int) {
	q.data = append(q.data, x)
}

func (q *Queue) Pop() int {
	x := q.data[q.head]
	q.head++
	if q.head*2 >= len(q.data) {
		q.data = append([]int(nil), q.data[q.head:]...)
		q.head = 0
	}
	return x
}

func (q *Queue) Front() int {
	return q.data[q.head]
}

func (q *Queue) Empty() bool {
	return q.head == len(q.data)
}

func (q *Queue) Len() int {
	return len(q.data) - q.head
}

func (q *Queue) Clear() {
	q.data = q.data[:0]
	q.head = 0
}

type LinkedList struct {
	val        []int
	prev, nxt  []int
	head, tail int
	size       int
}

func NewLinkedList(cap int) *LinkedList {
	return &LinkedList{
		val:  make([]int, 0, cap),
		prev: make([]int, 0, cap),
		nxt:  make([]int, 0, cap),
		head: -1,
		tail: -1,
	}
}

func (l *LinkedList) newNode(x int) int {
	id := len(l.val)
	l.val = append(l.val, x)
	l.prev = append(l.prev, -1)
	l.nxt = append(l.nxt, -1)
	l.size++
	return id
}

func (l *LinkedList) PushFront(x int) int {
	id := l.newNode(x)
	if l.head == -1 {
		l.head = id
		l.tail = id
		return id
	}
	l.nxt[id] = l.head
	l.prev[l.head] = id
	l.head = id
	return id
}

func (l *LinkedList) PushBack(x int) int {
	id := l.newNode(x)
	if l.tail == -1 {
		l.head = id
		l.tail = id
		return id
	}
	l.prev[id] = l.tail
	l.nxt[l.tail] = id
	l.tail = id
	return id
}

func (l *LinkedList) InsertAfter(pos int, x int) int {
	if pos == -1 {
		return l.PushFront(x)
	}
	if pos == l.tail {
		return l.PushBack(x)
	}
	id := l.newNode(x)
	after := l.nxt[pos]
	l.prev[id] = pos
	l.nxt[id] = after
	l.nxt[pos] = id
	l.prev[after] = id
	return id
}

func (l *LinkedList) InsertBefore(pos int, x int) int {
	if pos == -1 {
		return l.PushBack(x)
	}
	if pos == l.head {
		return l.PushFront(x)
	}
	id := l.newNode(x)
	before := l.prev[pos]
	l.prev[id] = before
	l.nxt[id] = pos
	l.nxt[before] = id
	l.prev[pos] = id
	return id
}

func (l *LinkedList) Delete(id int) {
	p := l.prev[id]
	n := l.nxt[id]
	if p != -1 {
		l.nxt[p] = n
	} else {
		l.head = n
	}
	if n != -1 {
		l.prev[n] = p
	} else {
		l.tail = p
	}
	l.prev[id] = -1
	l.nxt[id] = -1
	l.size--
}

func (l *LinkedList) Head() int {
	return l.head
}

func (l *LinkedList) Tail() int {
	return l.tail
}

func (l *LinkedList) Value(id int) int {
	return l.val[id]
}

func (l *LinkedList) Next(id int) int {
	return l.nxt[id]
}

func (l *LinkedList) Prev(id int) int {
	return l.prev[id]
}

func (l *LinkedList) Len() int {
	return l.size
}

func (l *LinkedList) ToSlice() []int {
	ans := make([]int, 0, l.size)
	for cur := l.head; cur != -1; cur = l.nxt[cur] {
		ans = append(ans, l.val[cur])
	}
	return ans
}

type Edge struct {
	To int
	W  int
}

type Graph struct {
	N        int
	Directed bool
	Adj      [][]Edge
}

func NewGraph(n int, directed bool) *Graph {
	return &Graph{
		N:        n,
		Directed: directed,
		Adj:      make([][]Edge, n),
	}
}

func (g *Graph) AddEdge(u int, v int, w int) {
	g.Adj[u] = append(g.Adj[u], Edge{To: v, W: w})
	if !g.Directed {
		g.Adj[v] = append(g.Adj[v], Edge{To: u, W: w})
	}
}

func (g *Graph) BFS(src int) []int {
	dist := make([]int, g.N)
	for i := range dist {
		dist[i] = -1
	}
	q := NewQueue(g.N)
	dist[src] = 0
	q.Push(src)
	for !q.Empty() {
		u := q.Pop()
		for _, e := range g.Adj[u] {
			if dist[e.To] == -1 {
				dist[e.To] = dist[u] + 1
				q.Push(e.To)
			}
		}
	}
	return dist
}

func (g *Graph) DFSOrder(src int) []int {
	vis := make([]bool, g.N)
	order := make([]int, 0, g.N)
	var dfs func(int)
	dfs = func(u int) {
		vis[u] = true
		order = append(order, u)
		for _, e := range g.Adj[u] {
			if !vis[e.To] {
				dfs(e.To)
			}
		}
	}
	dfs(src)
	return order
}

type Tree struct {
	N      int
	Adj    [][]int
	Parent []int
	Depth  []int
}

func NewTree(n int) *Tree {
	return &Tree{
		N:      n,
		Adj:    make([][]int, n),
		Parent: make([]int, n),
		Depth:  make([]int, n),
	}
}

func (t *Tree) AddEdge(u int, v int) {
	t.Adj[u] = append(t.Adj[u], v)
	t.Adj[v] = append(t.Adj[v], u)
}

func (t *Tree) Build(root int) {
	for i := 0; i < t.N; i++ {
		t.Parent[i] = -1
		t.Depth[i] = -1
	}
	q := NewQueue(t.N)
	t.Parent[root] = -1
	t.Depth[root] = 0
	q.Push(root)
	for !q.Empty() {
		u := q.Pop()
		for _, v := range t.Adj[u] {
			if v == t.Parent[u] {
				continue
			}
			t.Parent[v] = u
			t.Depth[v] = t.Depth[u] + 1
			q.Push(v)
		}
	}
}

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func NewTreeNode(x int) *TreeNode {
	return &TreeNode{Val: x}
}

type DSU struct {
	parent []int
	size   []int
	comps  int
}

func NewDSU(n int) *DSU {
	d := &DSU{
		parent: make([]int, n),
		size:   make([]int, n),
		comps:  n,
	}
	for i := 0; i < n; i++ {
		d.parent[i] = i
		d.size[i] = 1
	}
	return d
}

func (d *DSU) Find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.Find(d.parent[x])
	}
	return d.parent[x]
}

func (d *DSU) Union(a int, b int) bool {
	ra := d.Find(a)
	rb := d.Find(b)
	if ra == rb {
		return false
	}
	if d.size[ra] < d.size[rb] {
		ra, rb = rb, ra
	}
	d.parent[rb] = ra
	d.size[ra] += d.size[rb]
	d.comps--
	return true
}

func (d *DSU) Same(a int, b int) bool {
	return d.Find(a) == d.Find(b)
}

func (d *DSU) Size(x int) int {
	return d.size[d.Find(x)]
}

func (d *DSU) Components() int {
	return d.comps
}
