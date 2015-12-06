package main
import "fmt"
func fibonacci(n uint) uint {
	if n==0 {
		return 0
	}
	if n==1 {
		return 1
	}
	return fib(n-1)*fib(n-2)
}
func main(){
	fmt.Println(fibonacci(5))
}