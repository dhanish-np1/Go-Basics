package main 

import "fmt"

var x = "Hello world"


func main() {
	// We can declare a local variable with the same name as a package-level variable

	fmt.Println(x)// result: "Hello world"
	x:= "hello"
	fmt.Println(x) // result: "hello"

}