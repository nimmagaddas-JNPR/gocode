package main
import "fmt"
func main(){
	s:=[]string{"hi","go","language"}
	argu(s...) //arguments
}
func argu(s...string){//parameters
	for _,value:=range s {
	fmt.Println(value)
	
		
	}
}