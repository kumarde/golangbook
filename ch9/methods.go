package main

import(
    "fmt"
    "math"
)

func distance(x1, y1, x2, y2 float64) float64 {
    a := x2 - x1
    b := y2 - y1
    return math.Sqrt(a*a + b*b)
}

func (c *Circle) area() float64{
    return math.Pi * c.r * c.r
}

func (r *Rectangle) area() float64{
    l := distance(r.x1, r.y1, r.x1, r.y2)
    w := distance(r.x1, r.y1, r.x2, r.y1)
    return l * w
}

type Circle struct{
    x, y, r float64
}

type Rectangle struct{
    x1, y1, x2, y2 float64
}

func main(){
    c := Circle{3,3,3}
    r := Rectangle{0, 0, 10, 10}
    fmt.Println(c.area())
    fmt.Println(r.area())
}
