package main
	var a = b + c // a initialized third, to 3
	var b = f()	  // b initialized second, to 2, by calling f 
	var c = 1     // c initialized first, to 1

func main() {}

func f() int {return c+1 }