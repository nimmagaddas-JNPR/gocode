package main
import "fmt"
func greet(fname string,age int) string{
	return fmt.Sprint(fname,age)
	}
func main(){
	greet("john",27)
}