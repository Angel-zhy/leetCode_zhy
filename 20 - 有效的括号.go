package main

import "fmt"


//方法：栈
//1.
func isValid1(s string) bool {
	//测试空字符串
	if s == "" {
		return true
	}

	//定义一个栈
	stack := make([]int32, len(s))
	stackLength := 0

	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			//左括号,入栈
			stack[stackLength] = v
			stackLength++
		} else {
			//右括号,比较栈顶,匹配则移除,不匹配return false
			if stackLength == 0 {
				return false
			}
			if (v == ')' && stack[stackLength-1] == '(') || (v == ']' && stack[stackLength-1] == '[') || (v == '}' && stack[stackLength-1] == '{') {
				stackLength--
			} else {
				return false
			}
		}
	}

	return stackLength == 0
}

//2. ( 推荐 )
func isValid(s string) bool {
	//测试空字符串
	if s == "" {
		return true
	}
	// 奇数个括号肯定不能全匹配
	if len(s)%2 == 1 {
		return false
	}
	// 建立映射关系
	keyMap := map[string]string{
		"}": "{",
		"]": "[",
		")": "(",
	}
    var stack  []string
    for _,k := range s {
          //如果栈里有元素
          if len(stack)>0 {
			  // 如果和map里面的有匹配，也就是，目前元素是右半边括号
			  temp,ok := keyMap[string(k)]
			  if ok {
			  	   top := stack[len(stack)-1]
			  	   if top == temp {
			  	   	  //栈顶元素去掉pop
			  	   	  stack = stack[:len(stack)-1]
			  	   	  //跳过这次循环
			  	   	  continue
				   }
			  }
		  }
		  //空栈或者没匹配上，就加入栈
		  stack = append(stack,string(k))
	}

	// 如果最后栈是空的，说明都匹配上了，返回true
	return len(stack) == 0
}

func main()  {
	//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效.

	s := "({})"
   fmt.Println(isValid(s))
}