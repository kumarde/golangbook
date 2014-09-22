package main

import "fmt"

func swap(xPtr, yPtr *int){
    xValue := *xPtr
   *xPtr = *yPtr
   *yPtr = xValue
}

func main(){
    x := 1
    y := 2
    swap(&x, &y)
    fmt.Println(x, y)
}
