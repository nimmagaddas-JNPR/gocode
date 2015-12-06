package main
import "fmt"
func main(){
	var radius float64
	var area float64
	fmt.Println("Enter Radius:")
	fmt.Scanln(&radius)
	area=3.14*radius*radius
	fmt.Println("area of circle:",area)
}