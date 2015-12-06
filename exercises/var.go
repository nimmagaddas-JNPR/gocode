package main
import "fmt"
func main(){
	var num *[]string = make([]string)
	fmt.Println(num)
	fmt.Println(*num)
	}