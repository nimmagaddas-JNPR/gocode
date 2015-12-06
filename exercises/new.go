package main
import "fmt"
func main(){
	var i *int = new(int)
	fmt.Println(i)
	fmt.Println(*i)
	var s *string  = new(string)
	fmt.Println(s)
	fmt.Println(*s)
	var b*bool =new(bool)
	fmt.Println(b)
	fmt.Println(*b)
	}