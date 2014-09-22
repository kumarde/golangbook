package main

import (
    "fmt"
    "math"
)

func circleArea(c *Circle) float64 {
    return math.Pi * c.r * c.r
}

type Circle struct{
   x, y, r float64 
}

func main(){
    c := Circle{4, 3, 3}
    fmt.Println(circleArea(&c))
}
