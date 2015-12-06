package main
import(
	"fmt"
	"strconv"
)
func main(){
	var x int =7
	str:= "hi golang" + strconv.Itoa(x)//converts Integer to string
	fmt.Println(str)
	i, _:=strconv.Atoi("32")//converts String to integer
	fmt.Println(i+7)
}