package main

<<<<<<< Updated upstream
//this is file level scope. each file need to import it
import "fmt"
import "github.com/04_scope/01_package_scope/01"
func max(x int) int {
	max := 42 + x
	return max
}

func main() {
	max := max(7)
    _1.PrintVar()
	fmt.Println(max)
}
=======
import "fmt"

func main(){
	x:=42
	fmt.Println(x)
	foo()
}

func foo(){
	fmt.Println(y)
}

//this is in the out scope so foo can call this varaible
var y = 42
>>>>>>> Stashed changes
