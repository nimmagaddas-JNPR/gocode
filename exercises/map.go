package main
import "fmt"
func main(){
	map1 := map[int]string{
		0: "programming lang",
		1: "Go lang",
		2: "java",
		3: "c",
		4: "python",	
	}
	fmt.Println(len(map1))
	for _,val :=range map1{
	fmt.Println(val)
	}
	delete(map1,1)
	fmt.Println("after deleting")
	fmt.Println(map1)
	fmt.Println("\nchecking for 1rd element in map")
	if _,exists:= map1[1];exists{
	fmt.Println("value exists")
	fmt.Println(map1)
	}else{
	fmt.Println("value doesnot exists:")
	fmt.Println(map1)
	}
}