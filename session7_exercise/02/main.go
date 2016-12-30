package main
import "fmt"
func main(){
    r,b := func(num int) (int, bool){
        return num/2,num%2 ==0
    }(1)

    fmt.Println(r,b)
}
