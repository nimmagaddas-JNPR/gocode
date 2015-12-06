package main
import "fmt"
func main(){
var a=10
var b=20
var c=a|b
var d=a&b
var e=^a
fmt.Printf("%b\n",a)
fmt.Printf("%b\n",b)
fmt.Printf("OR operation:%b\n",c)
fmt.Printf("AND operation:%b\n",d)
fmt.Printf("NOT operation:%b",e)
}