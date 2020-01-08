package main

//方法一：遍历法
//解题思路：将原数组转成单个字符的二维数组遍历比较
func longestCommonPrefix(strs []string) string {
	strsLen := len(strs);
	if strsLen==0 {
		return ""
	}
	if strsLen==1 {
		return strs[0]
	}
	result:= strs[0]
	for i:=1;i<strsLen; i++ {
		if len(result)>len(strs[i]) {
			result = result[0:len(strs[i])]
		}
		if len(result)<1 {
			return result
		}
		for j:=0; j<len(result); j++ {
			if result[j]!=strs[i][j] {
				result = result[0:j]
				break
			}
		}
	}
	return result
}

func main(){
	 //最长公共前缀
	 s := []string{"flower","flow","flight"}
	 longestCommonPrefix(s)
}

