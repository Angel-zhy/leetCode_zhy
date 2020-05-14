package main

import "fmt"

/*
   factor:表示改位数上1
*/

func main() {
	// 给定一个整数 n，计算所有小于等于 n 的非负整数中数字 1 出现的个数。
	fmt.Println(countDigitOne(10001))
}

//  使用数学方法 Time: O(log10(n)), Space: O(1)
func countDigitOne(n int) int {
	// 边界情况
	if n < 1 {
		return 0
	}
	count, factor := 0, 1 //初始化计数器为0,和位数初始化为1
	for n/factor != 0 {
		//先求出当前位置上的数字(digit)
		digit := (n / factor) % 10 //n除以10倍的factor
		// 然后计算更高位的数字high
		high := n / (10 * factor) // n除以10倍的factor
		fmt.Println(digit, high)
		if digit == 0 {
			count += high * factor
		} else if digit == 1 {
			count += high * factor
			count += (n % factor) + 1
		} else {
			count += (high + 1) * factor
		}
		//计算完当前factor位上1出现的次数
		fmt.Println(count, factor)
		factor = factor * 10
	}
	return count
}
