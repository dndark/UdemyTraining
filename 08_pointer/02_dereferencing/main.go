package main

import "fmt"

//EVERYTHING IN GO IS COPY BY VALUE
func main() {

	a := 43

	fmt.Println(a)  // 43
	fmt.Println(&a) // 0x20818a220

	var b *int = &a
	fmt.Println(b)  // 0x20818a220
	fmt.Println(*b) // 43

	*b = 42        // b says, the value at this address chane it to 42
	fmt.Println(a) //42

	// b is a int pointer;
	//b pointer to the memory address where an int is store
	//to see the value in that memory address, add a * in front of the value
	// this is known as dereferencing
	// the * is an operator in this case
}
