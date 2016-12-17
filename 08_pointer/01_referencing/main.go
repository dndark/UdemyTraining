package main

import "fmt"

func main() {

	a := 43

	fmt.Println(a)
	fmt.Println(&a)

	//type -- pointer to a int
	var b *int = &a

	fmt.Println(b)

	// the above code makes b a pinter to the memory address where an int is store
	// b is a type of "

}
