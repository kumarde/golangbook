package main

import "fmt"

func fact(x uint) uint{
    if x == 0 {
        return 1
    }
    return x * fact(x - 1) 
}
