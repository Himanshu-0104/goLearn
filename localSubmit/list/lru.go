package main

type LRU struct {
	//1.map (keyNode)
	keyNode map[int]*DllNode
	//2. capacity
	capacity int
	//3. head
	head *DllNode
	//4. tail
	tail *DllNode
}
type DllNode struct {
	key int
	val int
	next *DllNode
	prev *DllNode
}

//constructor
func NewLRU(capacity int) *LRU {
	lru := &LRU{
		capacity: capacity,
		keyNode: make(map[int]*DllNode),
		head: &DllNode{}, //dummyNode
		tail: &DllNode{}, //dummyNode
	}
	lru.head.next = lru.tail
	lru.tail.prev = lru.head
	return lru
}

// lru method (func which attacked to type using receiver)
func (lru *LRU) GetLru(key int) int {
	if _, ok :=lru.keyNode[key]; !ok {
		return -1
	}
	node := lru.keyNode[key]
	lru.deleteNodeLRU(node) //how handle delete last node (edge case?)
	lru.addNodeLRU(node)
	return node.val
}

func (lru *LRU) PutLru(key int, val int) {
	if lru.capacity == 0 {
		return
	}
	if oldNode, ok := lru.keyNode[key]; ok {
		// lru.deleteNodeLRU(oldNode)
		// delete(lru.keyNode, key) //delete map[key]
		//below are optimised line to reuse current node instead of creating new node to use optimised comment above code
		lru.deleteNodeLRU(oldNode)
		oldNode.val = val
		lru.addNodeLRU(oldNode)
		return

	}
	if sz := len(lru.keyNode); sz == lru.capacity {
		lruNode := lru.tail.prev
		lru.deleteNodeLRU(lruNode)
		delete(lru.keyNode, lruNode.key) //delete map[key]
	}
	newNode := &DllNode{key: key, val: val} //create new Node (DllNode)
	lru.addNodeLRU(newNode)
	lru.keyNode[key] = newNode
}

func (lru *LRU) deleteNodeLRU(node *DllNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (lru *LRU) addNodeLRU(newNode *DllNode) {
	newNode.next = lru.head.next //newNode next to curr head
	newNode.prev = lru.head //newNode prev to curr head
	lru.head.next.prev = newNode //curr head next prev to newNode
	lru.head.next = newNode //curr head next to newNode
}


/* // LeetCode submited code [verdict: accepted]
type LRUCache struct {
    //1. keynode
    keyNode map[int]*DllNode
    capacity int
    head *DllNode
    tail *DllNode
}

type DllNode struct {
    key int
    val int
    next *DllNode
    prev *DllNode
}


func Constructor(capacity int) LRUCache { //change here with *LRUCache
    lru:=LRUCache{ //change here with &LRUCache
        capacity: capacity,
        keyNode: make(map[int]*DllNode),
        head:&DllNode{},
        tail:&DllNode{},//dummyNode
    }
    //dummyNode initialise
    lru.head.next=lru.tail
    lru.tail.prev=lru.head
    return lru
}


func (this *LRUCache) Get(key int) int {
    if _,ok := this.keyNode[key]; !ok {
       return -1
    }
    node:=this.keyNode[key]
    this.deleteLru(node)
    this.addNodeLru(node)
    return node.val
}


func (this *LRUCache) Put(key int, value int)  {
    if this.capacity==0 {
        return
    }
    if oldNode,ok := this.keyNode[key]; ok {
        this.deleteLru(oldNode)
        oldNode.val=value
        this.addNodeLru(oldNode)
        return
    }
    if sz:=len(this.keyNode); sz==this.capacity {
        lruNode:=this.tail.prev
        this.deleteLru(lruNode)
        delete(this.keyNode,lruNode.key)
    }
    newNode := &DllNode{key:key, val:value}
    this.addNodeLru(newNode)
    this.keyNode[key]=newNode
}



//  * Your LRUCache object will be instantiated and called as such:
//  * obj := Constructor(capacity);
//  * param_1 := obj.Get(key);
//  * obj.Put(key,value);



func (this *LRUCache) deleteLru (node *DllNode) {
    node.prev.next=node.next
    node.next.prev=node.prev
}
func (this *LRUCache) addNodeLru (newNode *DllNode) {
    newNode.next=this.head.next
    newNode.prev=this.head
    this.head.next.prev=newNode
    this.head.next=newNode
}


/*




/*  // OLD CODE
package main

type lru struct {
	//1.map (keyNode)
	keyNode map[int] *DllNode
	//2. capacity
	capacity int
	//3. head
	head *DllNode
	//4. tail
	tail *DllNode

}

type DllNode struct {
	key int
	val int
	next *DllNode
	prev *DllNode
}
// lru method (func which attacked to type using receiver)
func (lru *lru) createLRU(capacity) {
	lru.capacity = capacity
	lru.head.next = lru.tail
	lru.tail.prev = lru.head
}

func getLru(lru, key) int{
	if _,ok := lru.keyNode[key]; !ok {
		return -1
	}
	node = lru.keyNode[key]
	deleteNode(node) //how handle delete last node (edge cases)
	addNode(node)
	return node.val
}

func putLru(lru, key) {
	if lru.capacity == 0 {
		return
	}
	if _,ok := lru.keyNode[key]; ok {
		oldNode=lru.keyNode[key]
		deleteNode(oldNode)
		lru.keyNode[key] = nil //delete map[key]
	}
	if sz:= len(lru.keyNode); sz==lru.capacity {
		lruNode=lru.tail.prev
		deleteNode(lruNode)
		lru.keyNode[lruNode.key] = nil
	}
	newNode := &DllNode{key: key, val: val}  //createDllNode{key,val}
	addNode(newNode)
	lru.keyNode[key] = newNode
}

func deleteNode(node *DllNode) {
	node.prev.next=node.next
	node.next.prev=node.prev
}

func addNode(lru *lru, newNode *DllNode) {
	newNode.next=lru.head.next //newNode next to curr head
	newNode.prev=lru.head // newNode prev points to head
	lru.head.next.prev=newNode //currhead prev points to newNode
	lru.head.next=newNode //head pointing to newnode
}

*/