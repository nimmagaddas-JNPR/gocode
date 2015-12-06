package main
import(
	"fmt"
	"reflect"
)
func main(){
	name:="golang"
	year:=2015
	
	fmt.Println(reflect.TypeOf(name))
	
	fmt.Println(reflect.TypeOf(year))
}