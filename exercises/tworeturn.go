package main
import "fmt"
func greet(fmsg , lmsg string) (string, string) {
	return fmt.Sprintln(fmsg,lmsg), fmt.Sprint(lmsg,fmsg)
		}
func main(){
	fmt.Println(greet("good","luck"))
}