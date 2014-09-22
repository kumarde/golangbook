package main

import "fmt"

func main(){
    elements := make(map[string]string)
    elements["H"] = "Hydrogen"
    elements["He"] = "Helium"
    elements["N"] = "Nitrogen"
    elements["As"] = "Arsenic"
    elements["Ag"] = "Silver"

    fmt.Println(elements["He"])
    
    name, ok := elements["un"]
    fmt.Println(name, ok)
}
