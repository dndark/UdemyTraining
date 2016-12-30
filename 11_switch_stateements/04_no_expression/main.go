package main

import "fmt"


func main() {
    myFriend := "Mar"
	switch  {
    case len(myFriend) == 2:
        fmt.Println("my friend name have len 2")
	case myFriend == "Daniel":
		fmt.Println("Wassup Daniel")
	case myFriend == "Medhi":
		fmt.Println("Wassup Medhi")
    case myFriend == "Liam":
        fmt.Println("Liam Jenny")
    }
}
