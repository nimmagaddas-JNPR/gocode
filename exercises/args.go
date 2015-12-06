package main
import "fmt"
func main(){
	s:=[]string{"golang","world"}//Given the slice s and call
	Greeting("hello:", s...)
}
func Greeting(prefix string,who...string){
	fmt.Println(prefix)
	for _,value := range who {
		fmt.Println(value)
	}
}