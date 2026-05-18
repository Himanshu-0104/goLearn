package main

import "fmt"

type ListNode struct {
    Val  int
    Next *ListNode
}

func buildList(nums []int) *ListNode {

    var temp *ListNode
    temp = &ListNode{} //empty list 
    curr := temp

    for i:=0;i<len(nums);i++ {
        curr.Next = &ListNode{Val: nums[i]}
        curr=curr.Next
    }
    return temp.Next //bcz first node created at curr.Next
}

func listToSlice (head *ListNode) []int{
    nums:=make([]int,0)
    for head != nil {
        nums = append(nums,head.Val)
        head= head.Next
    }
    return nums
}

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


// // idomatic way
// func buildList(nums []int) *ListNode {
//     dummy := &ListNode{}
//     cur := dummy

//     for _, v := range nums {
//         cur.Next = &ListNode{Val: v}
//         cur = cur.Next
//     }

//     return dummy.Next
// }
