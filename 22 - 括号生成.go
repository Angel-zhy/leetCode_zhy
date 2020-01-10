package main

import "fmt"

//方法一：暴力法
//func generateParenthesis1(n int) []string {
//
//}

//方法二：递归法 （ 推荐！！！）（与 方法五 思路差不多）
func generateParenthesis2(n int) []string {
	res := make([]string, 0)
	gene2(&res,"",0,0,n)  //left表示字符str中左括号已经添加了多少个
	return  res
}

func gene2(res *[]string, str string, left, right ,max int)  {
     if len(str) == max*2 {
     	 *res = append(*res,str)
		 return
	 }
	 if left < max {                          //左括号还有剩余
		 gene2(res,str+"(",left+1,right,max)
	 }
	 if right < left {                        //右括号剩余 比 左括号剩余 多时 才能拼接右括号
		 gene2(res,str+")",left,right+1,max)
	 }
}

//方式三：闭合数


//方式四：递归（太复杂，不推荐）
//对于一个长度为2*n的有效的括号字符串，符合如下条件：
//1.其第一个字符必然是左括号“(”，其下标为0；
//2.该左括号对应的右括号“)”可以出现在下标为1，3，5，7 … … 2*n-1 等的位置；
//3.假设取右括号下标为k，那么在下标0到下标k之间的子串必然也是一个符合上述条件的有效括号字符串，并且下标在k+1到2n-1位置的字串亦是如此。
//这样一来，原问题（规模为n）就变成了两个子问题，规模分别为(k-1)/2和(2n-k-1)/2
//将两个子问题返回的字符串进行组合，就得到了原问题的解。
func generateParenthesis4(n int) []string {
	ans := make([]string, 0)

	for i := 1; i < n*2; i += 2 {
		str := "("
		// 求解子问题
		left := generateParenthesis4((i - 1) / 2)
		right := generateParenthesis4((2*n - i - 1) / 2)

		//如果子问题返回的切片是空的，就加一个空字符串方便下面的循环
		if len(left) == 0 {
			left = append(left, "")
		}

		if len(right) == 0 {
			right = append(right, "")
		}
		// 组合两个子问题返回的子字符串
		for _, vl := range left {
			for _, vr := range right {
				temp := str + vl + ")" + vr   // ")" 为分割字符（而且为：子字符串第1个左括号所对应的右括号）
				ans = append(ans, temp)
			}
		}
	}
	return ans
}


//方式五：（ 推荐！！！）
//思路：模拟将 n个'(' 和 n个')' 分别放入两个栈中left和right。每次从两个栈中取值,并且到str中
// 如果left==right，此时我们只能从left中出栈，
// 如果left<=right 那么就有两种选择，要么从left中出栈，要么充right中出栈。
//整个过程，要保证left<=right。
func generateParenthesis5(n int) []string {
	res := make([]string, 0)
	gene(&res,"",n,n)    //(结果切片ers，拼接字符str, 左括号剩余数，右括号剩余数)
	return res
}

func gene(res *[]string, str string, left, right int) {
	if left == 0 { // 此时表明left已经出栈完毕，需要把right中的全部出栈
		for i := 0; i < right; i++ {
			str += ")"
		}
		*res = append(*res, str)
		return
	}
	gene(res, str+"(", left-1,right)   //递归调用
	if left < right {
		gene(res, str+")", left,right-1)  //str是拷贝的，递归之间不会有影响
	}
}



func main()  {
	//给出 n 代表生成括号的对数，请你写出一个函数，使其能够生成所有可能的并且有效的括号组合
	fmt.Println(generateParenthesis2(3))
	fmt.Println(generateParenthesis4(3))
	fmt.Println(generateParenthesis5(3))
}
