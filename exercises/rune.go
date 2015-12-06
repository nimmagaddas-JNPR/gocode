package main
import ("fmt"
	"reflect"
)
func main(){
	rune:='d'
	note:="helloWorld"
	fmt.Println(reflect.TypeOf(rune))
	fmt.Println(note)
	fmt.Println(note[0])
	fmt.Println(note[4])
	fmt.Println("--------")
	fmt.Println("what the .....")
	fmt.Println(note[:4])
}
	
	 