package main

import "fmt"

func main() {

	a := 43

	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a
	fmt.Println(b)
	fmt.Println(*b)

	// b is a int pointer;
	//b pointer to the memory address where an int is store
	//to see the value in that memory address, add a * in front of the value
	// this is known as dereferencing
	// the * is an operator in this case
}
