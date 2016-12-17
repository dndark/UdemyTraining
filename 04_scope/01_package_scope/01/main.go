<<<<<<< Updated upstream
package _1

import "fmt"

func PrintVar(){
    fmt.Println(Myname)
    fmt.Println(yourName)
=======
package main

import "fmt"

//package level scope
var x int = 42

func main(){
	fmt.Println(x)
	foo()
	//local level scope
	y :=17
	fmt.Println(y)
}


func foo(){
	fmt.Println(x)
>>>>>>> Stashed changes
}
