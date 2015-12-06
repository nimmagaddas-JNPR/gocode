package main
import "fmt"
func main(){
	slice1 := []int{9,8,7,6,5}
	fun(slice1...)
	}
func fun(slice2...int){
	for _,value:= range slice2{
	fmt.Println(value)
	}
}