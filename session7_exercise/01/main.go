package main
import "fmt"
func main(){
    result,b:=half(2)
    fmt.Println(result,b)
}

func half(num int)(int, bool){
    return num/2,num%2 ==0
}
