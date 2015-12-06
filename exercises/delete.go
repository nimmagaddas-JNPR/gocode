package main
import "fmt"
func main(){
slice1 := []string{"h","e","l","o"}
slice2 := []string{"p","q","r"}
slice1 = append(slice1,slice2...)
fmt.Println(slice1)
slice2 = append(slice1[:3],slice2[:1]...)
fmt.Println(slice2)
}