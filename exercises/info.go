package main
import "fmt"
type info struct{
	id string
	name string
	sal float64
	age int
}	
func main(){
	c1 := new(info)
	c1.name = "william kennedy"
	c1.id = "p1234"
	c1.sal = 55.900
	c1.age = 30
	fmt.Println(c1.name)
	fmt.Println(*c1)
	fmt.Println("before update")
	c1.name="Martin"
	fmt.Println("after update")
	fmt.Println(c1.name)
	fmt.Println(*c1)
}

	