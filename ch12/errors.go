package main

import(
    "fmt"
    "errors"
)

func main(){
    err := errors.New("This is a real error.")
    fmt.Println(err)
}
