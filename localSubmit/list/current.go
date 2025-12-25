package main 

import (
    // "fmt"
)





















func Merge(intervals [][]int) [][]int {
    // //bruteforce n^2(each with other lefts)
    n:=len(intervals)
    
    ans:=make([][]int,0)
    visit:=make([]bool,n) //total no. of pairs
    for i:=0;i<n;i++ {
        if visit[i] {continue} // to skip already merged by innerloop
        visit[i]=true
        start:=intervals[i][0]
        end:=intervals[i][1]

        for j:=0; j<n; j++ {
            if visit[j] {continue}
            cstart:=intervals[j][0]
            cend:=intervals[j][1]
            //condition of overlap
            if max(cstart,start) <= min(cend,end) {
                start = min(cstart,start)
                end = max(cend,end)
                visit[j]=true //merged to skip in outerloop
            }
            
        }
        ans=append(ans,[]int{start,end})
    }
    return ans

    // //optimal nlogn(sort acc to starting point)
}








































