package main

import (
    "fmt"
)

type Person struct{
    Name string
}

type Android struct{
    Person Person
    Model string
}

func(p *Person) Talk(){
    fmt.Println("Hi, my name is ", p.Name)
}
