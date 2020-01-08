package main

import "fmt"

func lenSubstring(s string)int  {
	var left, res int //定义str的坐标left,如果重复left加1; 否则更新res
	usedChar := make(map[rune]int)  //str中包含的字符和及其最新坐标
	for i,v := range s{
		if index,ok := usedChar[v]; ok && left<=index {   //index为坐标
			left = index + 1
		}else{
			res = max(res,i-left+1)
		}
		usedChar[v] = i
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main()  {
	str := "abebfbcg"
	fmt.Println(lenSubstring(str))
}

