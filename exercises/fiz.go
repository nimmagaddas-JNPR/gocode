package main
import "fmt"
func main(){
	var i  int
	for i=1;i<=100;i++{
		if i%3==0{
			fmt.Print("fizz")
			}
		if i%5==0{
			fmt.Print("buzz")
			}
		if i%5==0{
			fmt.Print("buzz")
			}
		if i%3==0 && i%5==0{
			fmt.Println("fizzbuzz")
			}
		}
	}
	