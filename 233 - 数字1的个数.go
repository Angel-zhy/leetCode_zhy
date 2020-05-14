package main

import "fmt"

/*
   （从个位，十位，百位，这些十进制位上切入，来计算1出现的次数。）
   factor:表示计算位数为factor上为1时的所有情况，直到最高位；
   例如： 计算12345
   第1次循环：factor为1；即个位上的值为1的情况：
	  分两部分， 高位：1234  1
	  结果： （1234+1）种可能 （注意：0也是，高位为0，即数字为1）
   第2次循环：factor为10；即十位上的值为1的情况：
      分两部分,   高位：123 1 低位：5
	  结果： （123+1）*10
	  解析：高位上的值有124中可能， 低位上的值有10中可能
	  （注意：低位上的值为10的位数次方，因为低位上的数每一位都有10种选择【0~9】，
			  不同于高位上的数）
    第3次循环：factor为100；即百位上的值为1的情况：
	  分两部分,   高位：12 1 低位：45
	  结果： （12+1）*10^2
	第4次循环...
	......
	（直到：n/factor == 0）
*/

func main() {
	// 给定一个整数 n，计算所有小于等于 n 的非负整数中数字 1 出现的个数。
	fmt.Println(countDigitOne(12345))
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
		//高位：然后计算更高位的数字high
		high := n / (10 * factor) // n除以10倍的factor
		fmt.Println(digit, high)
		/*
		  高位取0值时，digit决定高低位的取值范围。
		  digit == 0 时，高位一定不能为0，低位取值范围是位数的次方；
		  digit == 1 时，高位可以为0；低位取值范围为低位的值+1
		  digit > 1 时，高位可以为0；低位取值范围是位数的次方；
		*/
		if digit == 0 {
			count += high * factor
		} else if digit == 1 {
			count += high * factor
			count += (n % factor) + 1
		} else {
			count += (high + 1) * factor
		}
		//计算完当前factor位上1出现的次数
		// fmt.Println(count, factor)
		factor = factor * 10
	}
	return count
}
