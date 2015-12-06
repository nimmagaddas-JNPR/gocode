package main
import "fmt"
func main(){
	note := func(){
		fmt.Println("Submit soon")
		}
	note()
	fmt.Printf("%T\n",note)
}