package main

import (
	"github.com/UdemyTraining/04_scope/01_package_scope/01"
	"fmt"
)


func max(x int) int {
	max := 42 + x
	return max
}

func main() {
	max := max(7)
   	 _1.PrintVar()
	fmt.Println(max)
}
