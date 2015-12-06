package main
import "fmt"
func main(){
	Note("hello: ","go","language")
}
func Note(prefix string, who....string) {
	fmt.Println(prefix)
	for _,value:= range who {
		fmt.Println(value)
	}
}