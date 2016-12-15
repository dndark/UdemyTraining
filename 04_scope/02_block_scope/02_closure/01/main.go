package main

import "fmt"

var x int = 42

func main(){
    x := 42
    fmt.Println(x)
    {
        fmt.Println(x)
        y := "Ths credit belong with the one "
        fmt.Println(y)
    }
}
