package main
import "fmt"
func greet() func() string{
	return func() string{
		return "hello notice"
	}
}
func main(){
	greet := greet()
	fmt.Println(greet())
}