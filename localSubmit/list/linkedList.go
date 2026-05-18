package main

import "fmt"

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