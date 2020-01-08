package main

import (
	"fmt"
)

//方法一：遍历法
//解题思路：将原数组转成单个字符的二维数组遍历比较
func longestCommonPrefix(str []string) string {
	strLen := len(str)
	if strLen==0 {
		return ""
	}
	if strLen==1 {
		return str[0]
	}
	result:= str[0]
	for i:=1;i<strLen; i++ {
		if len(result)>len(str[i]) {
			result = result[0:len(str[i])]
		}
		if len(result)<1 {  //后面的元素为空，即“”
			return result
		}
		for j:=0; j<len(result); j++ {
			if result[j]!=str[i][j] {  //二维数组
				result = result[0:j]
				break
			}
		}
	}
	return result
}


//方法二：分治法
//解题思路：采用分治法，直到分解成两个字符串比较。
func longestCommonPrefix2(str []string) string {
     strLen := len(str)
     if strLen == 0 {
     	return ""
	 }
	 if strLen == 1 {
	 	return str[0]
	 }
	 return comonPrefixLong(str,0,len(str)-1)
}

func comonPrefixLong(str []string, left, right int)string{
	if left == right {
		return str[left]
	}
	middle := (left+right)/2
	leftWord := comonPrefixLong(str,left,middle)
	rightword := comonPrefixLong(str,middle+1,right)
	return commonPreFix(leftWord,rightword)
}

func commonPreFix(res,work string)string  {
	if len(res) > len(work) {
		res = res[:len(work)]
	}
	if len(res) < 1 {
		return res
	}
	for j:=0; j<len(res); j++ {
		if res[j] != work[j] {
			res = res[:j]  //左闭右开
			break
		}
	}
	return res
}



func main(){
	 //最长公共前缀
	 s := []string{"flower","flow","flight"}
	 fmt.Println(longestCommonPrefix(s))
	 fmt.Println(longestCommonPrefix2(s))
}

