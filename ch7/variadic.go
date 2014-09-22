package main

import "fmt"

func add(args ...int) int{
    total := 0
    for _, value := range args{
       total += value 
    }
    return total
}

func main(){
    fmt.Println(add(1,2,3))
    fmt.Println(add(1,2,3,4,5,))
}
