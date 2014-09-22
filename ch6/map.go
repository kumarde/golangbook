package main

import "fmt"

func main(){
    x := make(map[string]int)
    x["key"] = 4
    fmt.Println(x)
    fmt.Println(x["key"])
}
