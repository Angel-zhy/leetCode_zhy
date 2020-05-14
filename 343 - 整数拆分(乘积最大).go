package main

import (
	"fmt"
	"math"
)

func main() {
	/*
			推论一： 将数字 n 尽可能以因子 3 等分时，乘积最大。
		    拆分规则：
			最优： 3 。把数字 n 可能拆为多个因子 3 ，余数可能为 0,1,2三种情况。
		    次优： 2 。若余数为 2 ；则保留，不再拆为 1+1 。
		    最差： 1 。若余数为 1 ；则应把一份 3 + 1 替换为 2 + 2，因为 2 x 2 > 3 x 1
	*/
	/*
	   时间复杂度 O(1) ： 仅有求整、求余、次方运算。
	   求整和求余运算：查阅资料，提到不超过机器数的整数可以看作是 O(1) ；
	   幂运算：查阅资料，提到浮点取幂为 O(1)。
	   空间复杂度 O(1) ： a 和 b 使用常数大小额外空间。
	*/
	//使用数学方法
	/*
			例子：（找规律）
		     2 => 2
			 3 => 2
			 4 => 2x2
		   	 5 => 2x3
			 6 => 3x3
			 7 => 2x2x3
			 8 => 2x3x3
			 9 => 3x3x3
			 10=> 2x2x3x3
		     11=> 2x3x3x3
		     12=> 3x3x3x3
	*/
	fmt.Println(integerBreak(11))
}

func integerBreak(n int) int {
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	x := n % 3
	if x == 0 {
		return int(math.Pow(3, float64(n/3)))
	} else if x == 1 {
		return 4 * int(math.Pow(3, float64(n-4)/3))
	} else {
		return 2 * int(math.Pow(3, float64(n-2)/3))
	}
}
