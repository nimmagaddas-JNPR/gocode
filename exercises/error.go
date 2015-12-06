package main
import "fmt"
func main(){
	capslice := make([]string,4,9)
	capslice[0]="go"
	capslice[1]="java"
	capslice[2]="c"
	capslice[3]="c++"
	capslice[4]="Programing lang"
	for _,value:= range capslice{
		fmt.Println(value)
	}
}