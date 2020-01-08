package main

import "fmt"

//方法：哈希表
//解题思路：
//  将所有字母的组合存入哈希表，遍历时 先判断 是否是双字母的，不是再判断是否是单字母的。
func romanToInt(s string) int {
	sMap := map[string]int     {"M":1000,"CM":900,"D":500,"CD":400,"C":100,"XC":90,"L":50,"XL":40,"X":10,"IX":9,"V":5,"IV":4,"I":1}
	result :=0
    for i:= 0; i<len(s); i++ {
    	if i<len(s)-1{
    		if _,ok := sMap[s[i:i+2]]; ok {  //先检测两个字母
    			result = result + sMap[s[i:i+2]]
    			i++  //这里加1，for上面加1,就跳过了两个字母
    			continue
			}
		}
		if _,ok := sMap[s[i:i+1]]; ok {
			result = result + sMap[s[i:i+1]]
		}
	}
	return result
}

func main(){

    str := "MCMXCIII"
	fmt.Println(romanToInt(str))
}
