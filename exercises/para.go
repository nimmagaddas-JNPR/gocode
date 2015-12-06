package main
import "fmt"
func average(sf...float64) float64{
	total := 0.0
	for _,v := range sf{
		total +=v
	}
	return total / float64(len(sf))
}
func main(){
	n:=average(45,78,23,0,67,91,62)
	fmt.Println(n)
}