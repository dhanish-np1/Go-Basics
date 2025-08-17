package main

import "fmt"

func main() {
	fmt.Println(feb(25)) // 25th Fibonacci number is 75,025
}

func feb(n int) int {
	x, y := 0,1
	for i:=0; i<n; i++ {
		x, y = y, x+y
	}
	return x
}