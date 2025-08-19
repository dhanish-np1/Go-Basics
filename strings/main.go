package main 

import "fmt"

var s = "hello world"

func main() {
	fmt.Println(s[:])     // hello world
	fmt.Println(s[2:5])	 //  llo
	fmt.Println(s[:7])  //   hello w
	fmt.Println(s[4:]) //    o world
}