package main

import "fmt"

func average(xs []float64) float64 {
    total := 0.0
    for _, value := range xs{
        total += value
    }
    return (total/float64(len(xs)))
}

func main(){
    xs := []float64{98,93,94,95,67}
    fmt.Println(average(xs))
}
