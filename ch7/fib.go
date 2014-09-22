package main

import "fmt"

func fib(x int) int{
    if x == 0 || x == 1 {
        return 1
    }
    return fib(x - 1) + fib(x - 2)
}

func fib_tail(x int) int{
    fib_help := func (a, b, n int) int {
        if n > 0{
            return fib_help(b, a+b, n-1)
        }else{
            return a
        }
    }
    return fib_help(0, 1, x)
}

func main(){
    fmt.Println(fib(5))
    fmt.Println(fib_tail(5))
}
