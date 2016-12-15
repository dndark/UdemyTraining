package main

import "fmt"

var x int = 42

//return a function type
func wrapper() func()  {
    x := 0
    return func()  {
        x++
        fmt.Println(x)
    }
}
func main() {
    a := wrapper()
	a()
    a()
}
