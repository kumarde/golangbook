package main

import "fmt"

func main(){
    var x [5]float64
    x[0] = 98
    x[1] = 93
    x[2] = 77
    x[3] = 88
    x[4] = 90
   
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
