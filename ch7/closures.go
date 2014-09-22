package main

import "fmt"

func main(){
    x := 0
    increment := func(){
        x++
    }
    increment()
    fmt.Println(x)
}
