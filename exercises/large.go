package main
import "fmt"
func max(numbers...int) int{
	var largest int
	for _, v:= range numbers{
		if v > largest{
			largest=v
		}
	}
	return largest
}
func main(){
	greatest := max(2,56,35,48,32,1,89,678,345,213)
	fmt.Println(greatest)
}