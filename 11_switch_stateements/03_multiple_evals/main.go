package main

import "fmt"

func main() {
	switch "Medhi" {
	case "Daniel","Jenny":
		fmt.Println("Wassup Daniel")
	case "Medhi":
		fmt.Println("Wassup Medhi")
    case "Liam":
        fmt.Println("Liam Jenny")
    }
}
