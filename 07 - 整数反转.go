package main

import (
	"fmt"
)


func reverse(x int) int  {
	y := 0
	for x!=0 {
		y = y*10 + x%10
		if !( -(1<<31) <= y && y <= (1<<31)-1) { //y值范围
			return 0
		}
		x /= 10
		fmt.Println(x)
		fmt.Println(y)
	}
	return y
}


func main(){
	//1. 123 -> 321
	//2. -123 -> -321
	//注意：  假设我们的环境只能存储得下 32 位的有符号整数，
	//则其数值范围为 [−2^31,  2^31 − 1]。请根据这个假设，
	//如果反转后整数溢出那么就返回 0。
	var x int = -55683461
	fmt.Println(reverse(x))
}
