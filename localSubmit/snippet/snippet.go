package main 

import (
    "bufio"
    "fmt"
    "os"
    "time"
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



func solve() {
    // fmt.Fprintln(out,"this is in solve")
    // var a,b int
    // fmt.Fscan(in,&a,&b)
    // fmt.Fprintln(out,a,b)
    
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