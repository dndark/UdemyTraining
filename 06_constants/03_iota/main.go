package main

import "fmt"

//iota is represent a very small number,
//if you keep declare it, it will increment
const (
	A = iota
	B = iota
	C = iota
)

const (
	_  = iota;  KB = 1 << (iota * 10) //1 << (1*10) 1 shift 10 place to the left
	MB = 1 << (iota * 10) //1 << (2*10)
)

const (	D = iota
	E = iota
	F = iota
)

func main() {
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
	fmt.Println(E)
	fmt.Println(F)
	fmt.Printf("%b\t", KB)
	fmt.Printf("%d\n", KB)

	fmt.Printf("%b\t", MB)
	fmt.Printf("%d\n", MB)
}
