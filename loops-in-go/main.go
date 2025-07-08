package main

import "fmt"

//The for loop is the only loop statement in Go
//But it has number of forms

func main() {
	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}

	massage := "Welcome to my house"

	for _, v := range massage {
		fmt.Println(v)
	}
}
