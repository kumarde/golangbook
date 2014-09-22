package main

import(
    "encoding/gob"
    "fmt"
    "net"
)

func server(){
    ln, err := net.Listen("tcp", ":9999")
    if err != nil {
        fmt.Println(err)
        return
    }
    for{
        //Accept a connection 
        c, err := ln.Accept()
        if err != nil {
            fmt.Println(err)
            continue
        }
        //Handle the connection
        go handleServerConnection(c)
    }
}

func handleServerConnection(c net.Conn){
    var msg string
    err := gob.NewDecoder(c).Decode(&msg)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("Received : ", msg)
    }
    c.Close()
}

func client(){
    c, err := net.Dial("tcp", ":9999")
    if err != nil {
        fmt.Println(err)
        return
    }
    //Send the messsage
    msg := "Hello, world!"
    fmt.Println("Sending", msg)
    err = gob.NewEncoder(c).Encode(msg)
    if err != nil {
        fmt.Println(err)
    }
    c.Close()
}

func main(){
    go server()
    go client()

    var input string
    fmt.Scanln(&input)
}
