package main

import (
	"fmt"
	"strconv"
)

//方式一：（首先将整数转换成字符串）
func isPalindrome1(x int) bool {
    str := strconv.Itoa(x)
    j := len(str)-1
    for i:=0; i<len(str)/2; i++ {
    	if str[i] != str[j] {
    		return false
		}
		j--
	}
	return true
}

//方法二：
func isPalindrome2(x int) bool {
	if x >= 0 {
		if x>0 && x%10 == 0 {
			return false  //大于0，并以0结尾的数，一定不会是回文数
		}
		var y,m int
		for {
			m = x % 10
			y = y*10 + m
			if y >= x || y >= x/10 {
				if y == x || y == x/10 {
					return true
				}
				return false
			}
			x = x/10
		}
	}
	return false
}

func main()  {
	 i := 1
     fmt.Println(isPalindrome1(i))
     fmt.Println(isPalindrome2(i))
}