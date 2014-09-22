package main

import "fmt"

func largest(args ...int) int{
    smallest := args[0]
    for _, value := range args{
        if value > smallest{
            smallest = value
        }
    }
    return smallest
}

func main(){
    fmt.Println(largest(34, 534, 3, 133, 90))
}
