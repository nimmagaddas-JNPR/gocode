package main
import "fmt"
func notify(numbers []int,callback func(int)){
	for _,n:=range numbers{
	callback(n)
	}
}
func main(){
	notify([]int{1,2,3,4,5,6},func(n int){
		fmt.Println(n)
	})
}