package main

import "fmt"

func main() {
	switch "Medhi" {
	case "Daniel":
		fmt.Println("Wassup Daniel")
	case "Medhi":
		fmt.Println("Wassup Medhi")
        fallthrough
	case "Jenny":
		fmt.Println("Wassup Jenny")
    case "Liam":
        fmt.Println("Liam Jenny")
    }
}
