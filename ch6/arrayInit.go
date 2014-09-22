package main

import "fmt"

func main(){
    x := [5]float64{
        90,
        87,
        32,
        245,
        23,
    }
    var total float64 = 0
    
    for i := 0; i < len(x); i++{
        total += x[i]
    }
    
    var total2 float64 = 0
    for _, value := range x {
        total2 += value
    }
    fmt.Println(total/float64(len(x)))
    fmt.Println(total2)
}
