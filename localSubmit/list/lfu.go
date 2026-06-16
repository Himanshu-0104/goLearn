package main

// methods are GetLfu,PutFlu, updateFrequency{{ as we are updating means for already existing not creating/new/inserting}}
type LFU struct {
	//1. map (keyNode)
	keyNode map[int]*DllNode
	//2. map (freqDll)
	freqDll map[int]*Dll //also a node just points head of that freq Dll <= wrong assumption need a layer top on DllNode which has head,tail,size(for helps in mantaining minFreq on deletion will need to set minFreq to 0, so helps here)
	// freqDll: map[int]*DllNode // <=wrong
	//3. capacity
	capacity int
	//4. minimumFrequency
	minFreq int
	//5. currentSize //this will runtime calculated by len(keyNode)

}

type DllNode struct {
	key int
	val int
	freq int
	next *DllNode
	prev *DllNode
}

// methods are addNodeDll,deleteNodeDll
type Dll struct {
	head *DllNode
	tail *DllNode
	size int
}
func newDll() *Dll{
	dll:=&Dll{
		head:&DllNode{}, //dummyNode
		tail:&DllNode{}, //dummyNode
		size:0
	}
	dll.head.next=dll.tail
	dll.tail.prev=dll.head
	return dll
}

func NewLFU (capacity int) *LFU {
	lfu:=&LFU{
		capacity:capacity,
		keyNode: make(map[int]*DllNode),
		freqDll: make(map[int]*Dll),
		// freqDll: make(map[int]*DllNode), // <=wrong
		minFreq: 0,
	}
	return lfu
}

func (this *LFU) GetLfu(key int) int {
	if _,ok := this.keyNode[key]; !ok {
		return -1
	}
	node:=this.keyNode[key]
	this.updateFrequency(node)
	return node.val
}

func (this *LFU) PutLfu(key int, value int) {
	if this.capacity == 0 {
		return
	}
	if oldNode,ok := this.keyNode[key]; ok {
		oldNode.val=value
		this.updateFrequency(oldNode)
		return
	}
	if sz:=len(this.keyNode); sz==this.capacity {
		lfuDll:=this.freqDll[this.minFreq]
		lfuNode:=lfuDll.tail.prev
		lfuDll.deleteNodeDll(lfuNode) //eviction
		if lfuDll.size == 0 {
			delete(this.freqDll,lfuNode.freq) //optional
		}
		delete(this.keyNode,lfuNode.key)
	}
	newNode:=&DllNode{key:key,val:value,freq:1}
	if _,ok:=this.freqDll[newNode.freq]; !ok{
		this.freqDll[newNode.freq] = newDll()
	}
	newNodeDll:=this.freqDll[newNode.freq]
	newNodeDll.addNodeDll(newNode)
	this.keyNode[key]=newNode
	this.minFreq=1
}

// this have (1.delete, 2.update freq, 3.add) logic
func (this *LFU) updateFrequency(node *DllNode) {
	nodeDll:=this.freqDll[node.freq]
	nodeDll.deleteNodeDll(node)
	if node.freq == this.minFreq && nodeDll.size == 0 {
		this.minFreq=node.freq+1
		delete(this.freqDll,node.freq) //optional
	}
	// node.freq=node.freq+1
	node.freq++
	if _,ok:= this.freqDll[node.freq]; !ok {
		this.freqDll[node.freq]= newDll()
	}
	nodeDll=this.freqDll[node.freq]
	nodeDll.addNodeDll(node)
}


func (this *Dll) addNodeDll(newNode *DllNode) {
	newNode.next=this.head.next
	newNode.prev=this.head
	this.head.next.prev=newNode
	this.head.next=newNode
	// this.size=this.size+1
	this.size++
}

func (this *Dll) deleteNodeDll(node *DllNode) {
	node.next.prev=node.prev
	node.prev.next=node.next
	// this.size=this.size-1
	this.size--
}

// func (this *Dll) lfuNodeDll() { //eviction //no need of this as we can directly access tail.prev for eviction

// }


/** // Leetcode submited code [verdict: accepted]

// methods are GetLfu,PutFlu, updateFrequency{{ as we are updating means for already existing not creating/new/inserting}}
type LFUCache struct {
    //1. keyNode
    keyNode map[int]*DllNode
    //2. freqDll
    freqDll map[int]*Dll
    //3. capacity
    capacity int
    //4. minFreq
    minFreq int
}

type DllNode struct {
    key int
    val int
    freq int
    next *DllNode
    prev *DllNode
}

// methods are addNodeDll,deleteNodeDll
type Dll struct {
    head *DllNode
    tail *DllNode
    size int
}

func newDll() *Dll{
    dll:= &Dll{
        head: &DllNode{},
        tail: &DllNode{},
        size:0,
    }
    dll.head.next=dll.tail
    dll.tail.prev=dll.head
    return dll
}




func Constructor(capacity int) LFUCache {
    lfu:=LFUCache{
        capacity: capacity,
        keyNode: make(map[int]*DllNode),
        freqDll: make(map[int]*Dll),
        minFreq:0,
    }
    return lfu
}


func (this *LFUCache) Get(key int) int {
    if _,ok := this.keyNode[key]; !ok {
        return -1
    }
    node:=this.keyNode[key]
    this.updateFreq(node)
    return node.val
}


func (this *LFUCache) Put(key int, value int)  {
    if this.capacity ==0 {
        return
    }
    if oldNode,ok := this.keyNode[key]; ok{
        oldNode.val=value
        this.updateFreq(oldNode)
        return
    }
    if sz:=len(this.keyNode); sz==this.capacity {
        lfuDll:= this.freqDll[this.minFreq]
        lfuNode:=lfuDll.tail.prev
        lfuDll.deleteNodeDll(lfuNode)
        if lfuDll.size == 0 {
            delete(this.freqDll,lfuNode.freq)
        }
        delete(this.keyNode,lfuNode.key)
    }
    newNode:=&DllNode{key:key,val:value,freq:1}
    if _,ok := this.freqDll[newNode.freq]; !ok {
        this.freqDll[newNode.freq] = newDll()
    }
    newNodeDll:= this.freqDll[newNode.freq]
    newNodeDll.addNodeDll(newNode)
    this.keyNode[newNode.key] = newNode
    this.minFreq=1
}

func (this *LFUCache) updateFreq(node *DllNode) {
    nodeDll:=this.freqDll[node.freq]
    nodeDll.deleteNodeDll(node)
    if node.freq == this.minFreq && nodeDll.size == 0 {
        this.minFreq=node.freq+1
        delete(this.freqDll,node.freq)
    }
    node.freq++
    if _,ok:=this.freqDll[node.freq]; !ok {
        this.freqDll[node.freq] = newDll()
    }
    nodeDll=this.freqDll[node.freq]
    nodeDll.addNodeDll(node)
}



//  * Your LFUCache object will be instantiated and called as such:
//  * obj := Constructor(capacity);
//  * param_1 := obj.Get(key);
//  * obj.Put(key,value);



func (this *Dll) addNodeDll(newNode *DllNode) {
    newNode.next= this.head.next
    newNode.prev= this.head
    this.head.next.prev=newNode
    this.head.next=newNode
    this.size++
}
func (this *Dll) deleteNodeDll(node *DllNode) {
    node.prev.next=node.next
    node.next.prev=node.prev
    this.size--
}



*/