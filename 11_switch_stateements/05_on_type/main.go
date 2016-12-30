package main

import "fmt"

type Contact struct {
	greeting string
	name     string
}

func SwitchOnType(x interface{}) {
	switch x.(type) { // this is an assert; asserting, "x is of this type"
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case Contact:
		fmt.Println("Contact")
	default:
		fmt.Println("unknow")
	}
}
func main() {
	SwitchOnType(7)
	SwitchOnType("Mclead")
	var t = Contact{greeting: "Good to see you", name: "Liam"}
    SwitchOnType(t)
}
