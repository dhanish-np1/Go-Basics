package main

import "fmt"

func main() {
	if x:= f(); x==0{
		fmt.Println(x)
	}else if y:=g(x); x == y{
		fmt.Println(x,y)
	}else {
		fmt.Println("else",x,y)
	}

}

func g(x int) int {
	return x
}


func f() int {
	return 1
}