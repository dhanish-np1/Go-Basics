package main

import "fmt"

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}

}

func popCount(x uint64) int {
	return int(pc[byte(x>>(0*8))]+
		pc[byte(x>>(1*8))]+
		pc[byte(x>>(2*8))]+
		pc[byte(x>>(3*8))]+
		pc[byte(x>>(4*8))]+
		pc[byte(x>>(5*8))]+
		pc[byte(x>>(6*8))]+
		pc[byte(x>>(7*8))])	
}

func main() {
	fmt.Println(popCount(0))     // 0 → no ones
	fmt.Println(popCount(7))     // 7 (111) → 3 ones
	fmt.Println(popCount(255))   // 255 (11111111) → 8 ones
	fmt.Println(popCount(1023))  // 1023 (1111111111) → 10 ones
	fmt.Println(popCount(1000023))
	fmt.Println(popCount(18446744073709551615)) // → 64 ones
}