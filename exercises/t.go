package main
import(
	"fmt"
	"reflect"
)
func main(){
	length:=30
	fmt.Printf("%T\n",length)
	fmt.Println(reflect.TypeOf(length))
}