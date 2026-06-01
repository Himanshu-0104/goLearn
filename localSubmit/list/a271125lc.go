package main 

import (
    "bufio"
    "fmt"
    "os"
    "time"
    // "sort"
)



var Himaydv0104 = ""  //exported
// go run -ldflags "-X main.Himaydv0104=local" .

var (
    in  *bufio.Reader
    out *bufio.Writer
)

type (
	ll  = int64
	ull = uint64
	ld  = float64
)

const (
	INF = ll(1e18)
	MOD = ll(1e9 + 7)
)


//Ques11
func ques11solve(nums []int) int{
    //bruteforce 
    n:=len(nums)
    ans:=0
    for i:=0;i<n;i++ {
        for j:=i;j<n;j++ {
            if nums[j]<nums[i] {
                ans+=1
            }
        }
    }
    return ans

    //optimal merge sort

}

//Ques25 
// func reverseList(head *ListNode) *ListNode{
//     var prev *ListNode
//     cur := head

//     for cur != nil {
//         next := cur.Next
//         cur.Next = prev
//         prev = cur
//         cur = next
//     }

//     return prev
// }

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    // using Stack
    var st Stack
    temp := head
    for temp!=nil {
        st.Push(temp.Val)
        temp=temp.Next
    }
    temp = head
    for temp != nil {
        temp.Val = st.Top()
        st.Pop()
        temp = temp.Next
    }
    return head

    // //Iterative
    // var prev *ListNode
    // temp:=head
    // for temp !=nil {
    //     front:=temp.Next
    //     //change pointer (prev, first assign already prev then update prev)
    //     // 1step
    //     temp.Next = prev 
    //     prev=temp

    //     //2nd step
    //     temp=front
    // }
    // return prev

    // // Recursive
    // if head == nil || head.Next == nil {
    //     return head
    // }
    // newHead:= reverseList(head.Next)
    // front:=head.Next
    // front.Next=head
    // head.Next = nil
    // return newHead
}

func solve() {
    
/*

    // Ques11
    // var n int 
    // fmt.Fscan(in,&n)
    // nums:=make([]int,n)
    // for i:=0;i<n;i++ {
    //     fmt.Fscan(in,&nums[i])
    // }
    // ans:=ques11solve(nums)
    // fmt.Fprintln(out,ans)


    // Ques12

    // Ques25
*/
    n:=0
    fmt.Fscan(in,&n)
    nums := make([] int,n)
    for i:=0;i<n;i++ {
        fmt.Fscan(in,&nums[i])
    }

    head := buildList(nums)
    ans:= reverseList(head)
    pl:= listToSlice(ans)
    fmt.Fprintln(out,pl)
    // printList(ans)

    



















    

    

    

    
}



func advCon_time(){
    start := time.Now()
    t:=1
    fmt.Fscan(in,&t)
    for t>0 {
        solve()
        t--
    }
    if Himaydv0104 != "" {
		_ = start
		// fmt.Fprintf(out,"\n\nTime: %d ms", time.Since(start).Milliseconds())
	}

}


func main() {
    if Himaydv0104 != ""{
        IOES()
    }
    in = bufio.NewReader(os.Stdin)
    out = bufio.NewWriter(os.Stdout)
    defer out.Flush()

    advCon_time()

    
}


func ceilDiv(a, b ll) ll {
	if a%b == 0 {
		return a / b
	}
	if a > 0 {
		return a/b + 1
	}
	return a / b
}

func gcd(a, b ll) ll {
	for b != 0 {
		a, b = b, a%b
	}
	if a < 0 {
		return -a
	}
	return a
}

func power(base, p ll) ll {
	ans := ll(1)
	for p > 0 {
		if p%2 == 0 {
			base *= base
			p /= 2
		} else {
			ans *= base
			p--
		}
	}
	return ans
}

func max(a int, b int ) int {
    if a>b {return a}
    return b
}

/* //Linked List DS
type ListNode struct {
    Val  int
    Next *ListNode
}

// create linked list from input
func buildList(nums []int) *ListNode {
    dummy := &ListNode{}
    cur := dummy

    for _, v := range nums {
        cur.Next = &ListNode{Val: v}
        cur = cur.Next
    }

    return dummy.Next
}

// print linked list
func printList(head *ListNode) {
    for head != nil {
        fmt.Fprint(out, head.Val)
        if head.Next != nil {
            fmt.Fprint(out, " ")
        }
        head = head.Next
    }
    fmt.Fprintln(out)
}
*/


/* //Max Heap DS
type maxHeap struct {
    data []int
}

//h.insert(val) 
// func (h *maxHeap) insert(value int) {
//insert(h,val)
func insert(h *maxHeap, value int) {
    h.data = append(h.data,value)
    heapifyUp(len(h, h.data)-1) //what this does => refactor heap
}

func heapifyUp(h *maxHeap, idx int) {
    for idx > 0 { //logn
        par := (idx-1) / 2
        if h.data[idx] > h.data[par] {
            h.data[par],h.data[idx] = h.data[idx],h.data[par]
            idx = par
        } else {
            break
        }
    }
}

func delete(h *maxHeap) (int,bool) {
    if len(h.data) == 0 {
        return 0, false
    }
    mxVal := h.data[0]
    h.data[0]=h.data[len(h.data)-1] //last ele gets first
    h.data = h.data[: len(h.data)-1] //free size
    heapifyDown(0) // why zero => bcz delte from top(highest) so refactor down from top

    return mxVal, true
}

func heapifyDown(h *maxHeap, idx int) {
    n = len(h.data)-1
    for {
        leftIdx := h.data[2*idx+1]
        rightIdx = h.data[2*idx+2]
        largestIdx = idx

        if leftIdx<=n && h.data[leftIdx] > h.data[largestIdx] {
            largestIdx = leftIdx
        } 
        if rightIdx<=n && h.data[rightIdx] > h.data[largestIdx] {
            largestIdx = rightIdx
        }
        if largestIdx == idx {
            break
        }

        h.data[idx], h.data[largestIdx] = h.data[largestIdx], h.data[idx]
        idx = largestIdx // due to this log(n) , bcz traverse only below this
    }
}

func peek(h *maxHeap) (int,bool) {
    if len(h.data == 0) {return 0,false}
    return h.data[0],true
}

func size(h *maxHeap) int {
    return len(h.data)
}

func isEmpty(h *maxHeap) bool {
    return len(h.data) == 0
}
*/
