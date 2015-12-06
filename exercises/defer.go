package main
import "fmt"
func greet(){
	fmt.Println("greet")
}
func well(){
	fmt.Println("well")
}
func main(){
	defer well()
	greet()
}
	