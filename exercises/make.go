package main
import "fmt"
func main(){
	capslice:=make([]string,5,10)
	capslice[0]="web"
	capslice[1]="programming"
	capslice[2]="go"
	capslice[3]="language"
	capslice[4]="world"
	for _,value:=range capslice{
	fmt.Println(value)
	}
}