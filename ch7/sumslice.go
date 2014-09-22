package main

import "fmt"

func sum(xs []int) int{
    total := 0
    for _, value := range xs{
        total += value
    }
    return total
}

func main(){
    xs := []int{56, 76, 79, 34, 43}
    fmt.Println(sum(xs))
}
