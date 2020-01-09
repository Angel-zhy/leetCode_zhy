package main

import "fmt"

func letterCombinations(digits string) []string {
	ret := make([]string, 0)

	if digits == ""{
		return ret
	}
	digit2LetterList := map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}
	ret = append(ret, "")
	//fmt.Println(ret,len(ret))
	for i := 0; i < len(digits); i++ {
		retLen := len(ret)        //结果数组长度，前一次遍历得到的结果数 3
		//fmt.Println(retLen)
		digit := digits[i : i+1]   //遍历到的数字
		letterList := digit2LetterList[digit]   //此数字对应的字母列表
		for j := 0; j < retLen; j++ {
			for _, letter := range letterList {
				ret = append(ret, ret[0]+letter)
				//fmt.Println(ret)
			}
			ret = ret[1:]
			//fmt.Println(ret)
		   //没循环一次去掉第1个结果，letLen有多长，就去多好次，最后把上一个数字得到的结果全部去除。
		}

	}

	return ret
}


func letterCombinations1(digits string) []string {
	if len(digits)==0{
		return nil
	}
	m := map[byte][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}
	s := m[digits[0]]   //s为字符片数组
	for i := 1; i < len(digits); i++ {
		temp := make([]string, 0)
		for j := 0; j < len(s); j++ {
			for k := 0; k < len(m[digits[i]]); k++ {  //m[digits[i]]表示第i个数字包含多少个字符
				temp = append(temp, s[j] + m[digits[i]][k])
			}
		}
		s = temp
	}
	return s
}


func main(){
	//给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
	//给出数字到字母的映射如下（与电话按键相同）：
	//输入："23"
	//输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
	str := "234"
	fmt.Println(letterCombinations(str))
	fmt.Println(letterCombinations1(str))   //方法原理一样，但是这里函数逻辑好理解一点
}





