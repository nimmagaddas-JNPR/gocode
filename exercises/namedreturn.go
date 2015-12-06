package main
import "fmt"
func greet(fmsg , lmsg string) (s string) {
	s= fmt.Sprintln(fmsg,lmsg)
	return	
}
func main(){
	fmt.Println(greet("good","luck"))
}