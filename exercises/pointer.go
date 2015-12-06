package main
import "fmt"
func main(){
a:=30
fmt.Println(a)
fmt.Println("Memory address is",&a)
var b *int=&a
fmt.Println("Pointer value is:",*b)
}