package main

import "fmt"

//时间复杂度：O(n)
func convert(s string, numRows int) string {
	// 不满足，提前返回
	if len(s) <= 2 || numRows == 1 {
		return s
	}
	// 保存最终结果
	var res string
	// 保存每次的结果
	var tmp = make([]string, numRows)
	// 初始位置
	curPos := 0
	// 用来标记是否该转向了
	shouldTurn := -1
	// 遍历每个元素
	for _, val := range s {
		// 添加进tmp里面，位置为curPos
		tmp[curPos] += string(val)
		// 如果走到头或者尾，就该转向了
		// 因为就在numRows的长度里面左右震荡
		if curPos == 0 || curPos == numRows-1 {
			// 转向
			shouldTurn = -shouldTurn
		}
		// curPos 用来标记tmp里面我们该去哪个位置
		curPos += shouldTurn
		fmt.Println(curPos)
	}
	fmt.Println(tmp)
	// 这时tmp里面已经保存了数据了，我们只需要转换一下输出格式
	for _, val := range tmp {
		res += val
	}
	// 最后的输出
	return res
}

func main()  {
	//以从上往下、从左到右进行 Z 字形排列
	//比如输入字符串为 "LEETCODEISHIRING" 

	//1.行数为 3 时，排列如下：
	//L   C   I   R
	//E T O E S I I G
	//E   D   H   N

    //2.行数为 4 时，排列如下：
	//L     D     R
	//E   O E   I I
	//E C   I H   N
	//T     S     G

	 str := "LEETCODEISHIRING"
	 fmt.Println(convert(str,4))
}
