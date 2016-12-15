package main

import "fmt"

func main() {
	fmt.Println("hello world")
	fmt.Printf("%d -%b -%x \n", 42, 42, 42)
	fmt.Printf("%d -%b -%#X \n", 42, 42, 42)

	for i := 0; i < 200; i++ {
		fmt.Printf("%d \t %b \t %x  \n", i, i, i)
	}
}
