package main

import (
	"fmt"
	"math"
	"strings"
)


//方法一：
func myAtoi1(str string) int {
	//str = strings.Replace(str," ","",-1)  //去所有空格
	str = strings.TrimSpace(str)  //去两头的空格
	if len(str) == 0 {
		return 0
	}
	//fmt.Printf("%q\n",str)
	//fmt.Printf("%c\n",str[0])
	if !(str[0] >= '0' && str[0]<='9') && str[0] !='-' && str[0] !='+' {
		return 0   //不为o~9开头 ； 不以-或者+开头
	}
    if (len(str)==1 && str[0]=='+') || (len(str)==1 && str[0]=='-') {
    	return 0   //已-或者+开头，但是长度为+
	}

    pInteger := 1  //正整数 Positive integer
	start :=0  //数字开启位置
    if str[0] == '-' {
		pInteger = -1
		start = 1
	}
	if str[0] =='+'{
		start = 1
	}



	//找到开始的的位置(数字为1~9位置)
	if str[start] =='0'{
		i:=0
		for i=start;i<len(str);i++{
			if str[i] !='0'{
				start = i
				break
			}
		}
		//全部为零的后者不为零但也不是数字 例如： 00000000paf123
		if i == len(str) || (str[i] >'9' || str[i] <'0'){
			return 0
		}
	}

	//找到end的位置
	end := len(str)
	for i:=start;i<len(str);i++{
		if !(str[i] >='0' && str[i]<='9'){
			end = i
			break
		}
		if str[i] =='.'{  //作用？？？？？？？？？？？？？？
			end = i-1
			break
		}
	}


	result :=0
	//判断最后的结果不能超过指定范围
	for ;start<end;start++{
		n  := int(str[start]-'0')
		powN := end-start-1
		if powN >=10 && pInteger==1{
			return math.MaxInt32
		}

		if powN >=10 && pInteger==-1{
			return math.MinInt32
		}

		temp := int32(math.Pow(float64(10),float64(powN)))

		n *= int(temp)
		result += n
	}
	result = result *pInteger
	if result >= math.MaxInt32 {
		return math.MaxInt32
	}

	if result <= math.MinInt32 {
		return math.MinInt32
	}
	return result

}


//方法二：（推荐！！！）
const INT_MAX = int(^uint32(0) >> 1)
const INT_MIN = ^INT_MAX
func myAtoi2(str string) int{
	n := 0   //返回变量
	init := false  //一旦遇到不会空格时，就改为true
	m := 1  //正负数
	for _, char := range []byte(str) {
		//fmt.Printf("%c\n",char)
		switch  {
			case char == ' ' :
				if init {
					return n
				}
		    case char == '+' :
		    	if !init {
		    		init = true
		    		m = 1
				}else {
					return n
				}
			case char == '-' :
				if !init {
					init = true
					m = -1
				}else {
					return n
				}
		    case char >= '0' && char <= '9':
				init = true
				n = n * 10 + int(char-'0') * m   //int(char-'0') 重点！！！
				if n > INT_MAX {
					return INT_MAX
				}
				if n < INT_MIN {
					return INT_MIN
				}
		    default:
			    if init{
			    	return n
				}else{
					return 0
				}
		}


	}
    return n
}



func main(){
    //1.如果数值超过其数值范围为 [−231,  231 − 1]，返回INT_MAX (231 − 1) 或 INT_MIN (−231)
    //2.该函数会根据需要丢弃无用的开头空格字符，直到寻找到第一个非空格的字符为止。
    str := "        -127345.3   asda   df   "
    fmt.Println(myAtoi1(str))

	//推荐
	fmt.Println(myAtoi2(str))

}