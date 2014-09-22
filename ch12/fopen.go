package main

import(
    "fmt"
    "os"
)

func main(){
    file, err := os.Open("file.txt")
    if err != nil{
        //Handle error via error
        return
    }
    defer file.Close()

    stat, err := file.Stat()
    if err != nil{
        //Handle the error via Error
        return
    }

    bs := make([]byte, stat.Size())

    _, err = file.Read(bs)
    if err != nil {
        return
    }

    str := string(bs)
    fmt.Println(str)
}
