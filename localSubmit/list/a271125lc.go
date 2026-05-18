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
func reverseList(head *ListNode) *ListNode{
    var prev *ListNode
    cur := head

    for cur != nil {
        next := cur.Next
        cur.Next = prev
        prev = cur
        cur = next
    }

    return prev
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
