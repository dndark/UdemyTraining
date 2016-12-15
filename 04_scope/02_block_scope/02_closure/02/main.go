package main

import "fmt"

var x int = 42

func main() {
	x := 0
	// ananymous function a function without a name
	// func expression
	// assigning a function to a varaible
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
}
