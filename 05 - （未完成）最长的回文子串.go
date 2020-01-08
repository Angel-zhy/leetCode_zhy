package main

import (
	"fmt"
	"math"
	"strings"
)

func main()  {
	//回文子串 ：指正着看和倒着看是一样的，例如 ：abcba
	//"babad"  => 输出："bab" 或者  "aba"
	//"cbbd"   => 输出： "bb"
	str := "babacdca"
	fmt.Println(longestPalindrome(str))
	fmt.Println(longestPalindrome2(str))
}

//https://blog.csdn.net/heart66_A/article/details/83957358
//方法一：
//时间复杂度O(n^3)
func longestPalindrome(s string) string {

	if strings.TrimSpace(s) == ""{
		return ""
	}

	sLen := len(s)//babad
	maxStr :=string(s[0])
	maxStrLen :=0
	for i:=0;i<sLen;i++{  //str := "babadada" //输出的结果应该为："adada"
		location := strings.IndexByte(s[i:],s[i])
		location +=i
		for location!=-1{
			//比较是否是回文数
			temp := string(s[i:location+1])
			if temp== reverceString(temp) && len(temp) >=maxStrLen{
				maxStrLen = len(temp)
				maxStr = temp
			}
			location++ //下次找的起点位置
			if i := strings.IndexByte(string(s[location:]),s[i]);i==-1{
				location =-1

			}else{
				location +=i
			}
		}
	}

	return maxStr
}

func reverceString(str string)string {
	bs := []byte(str)

	//分开写，看这更清楚
	from, to := 0, len(bs)-1
	for from < to {
		bs[from], bs[to] = bs[to], bs[from]
		from++
		to--
	}
	return string(bs)

}


//方法二：
//中心扩展算法，可以在n*n之间完成
func longestPalindrome2(s string) string {   //babad bab
	if strings.TrimSpace(s) =="" || len(s) <1 {
		return ""
	}
	start,end := 0,0
	for i:=0;i<len(s);i++{
		len1 :=expandAroundCenter(s,i,i) //a [b] a bd 可能回文数的中间数是一个 则从一个数开始两边扩散
		len2 := expandAroundCenter(s,i,i+1)//a [bb] c 可能回文数的中间数可以使两个。从两个开始两边是扩散
		len := int(math.Max(float64(len1),float64(len2))) //比较出两个的最长的长度
		if len > end-start{
			//len=3  i=1
			start =i-(len-1)/2 //i:在这里就是一个回文数的中间数，（len-1）:是处理len可能是偶数的情况
			end = i+len/2
		}
	}
	return string(s[start:end+1])
}
////babad bab
func expandAroundCenter(s string,left,right int)int{
	L,R := left,right
	for L>=0 && R <len(s) && s[L] ==s[R]{
		L--
		R++
	}
	return R-L-1 //计算出该串的长度
}