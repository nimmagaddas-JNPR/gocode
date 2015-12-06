package main
import "fmt"
func main(){
	var name interface{}="sydney"
	str,ok:= name.(string)
	if ok{
		fmt.Printf("%T\n",string)
	} else {
		fmt.Printf("value is not string\n")
	}
}