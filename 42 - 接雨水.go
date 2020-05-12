package main

import (
	"fmt"
)

func main()  {
	//给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水.
    //输入: [1,0,2,1,0,1,2]  图见：https://leetcode-cn.com/problems/trapping-rain-water/
    slice := []int{2,0,2}

    //方法1：暴力法
    count1 := trap1(slice)
    fmt.Println(count1)

    //方法2：动态规划
	count2 := trap2(slice)
	fmt.Println(count2)

    //方法3： 使用栈来处理（推荐）
    count3 := trap3(slice)
	fmt.Println(count3)

	//方法4.双指针法（推荐）
	slice1 := []int{1,0,2,1,3,0,1,2}
	count4 := trap4(slice1)
	fmt.Println(count4)
}

//1.暴力法
/*
时间复杂度：O(n^2)
*/
func trap1(height []int)int{
	var ans int
	cols := len(height)

	for i:=1; i<cols-1; i++ {
		var maxLeft int
		var maxRight int
		var lower int

		for j:=0; j<i; j++ {
			maxLeft = Max(maxLeft, height[j])
		}

		for j:= i+1; j<cols; j++{
			maxRight = Max(maxRight,height[j])
		}

		lower = Min(maxLeft,maxRight)
		//fmt.Println(i,maxLeft,maxRight)
		if lower > height[i]{
			ans = ans + lower - height[i]
		}
	}
	return ans
}

//2.动态规划
/*
时间复杂度：O(n)
空间复杂度：O(n)，用来保存每一列左边最高的墙和右边最高的墙。
*/
func trap2(height []int)int{
	var sum int
	maxLeft := make([]int,len(height))  //maxLeft [i] 代表第 i 列左边最高的墙的高度
	maxRight := make([]int,len(height))  //maxRight [i] 代表第 i 列右边最高的墙的高度

	for i:= 1; i < len(height) - 1; i++{
		maxLeft[i] = Max(maxLeft[i - 1], height[i - 1])
	}
	for i:= len(height) - 2; i >= 0; i-- {
		maxRight[i] = Max(maxRight[i + 1], height[i + 1])
	}
	for  i:= 1; i < len(height) - 1; i++ {
		min := Min(maxLeft[i], maxRight[i])
		if min > height[i] {
			sum = sum + (min - height[i])
		}
	}
	return sum
}


//3.栈方法
/*
解法思想： 两墙夹一水，在找到右墙后，左墙数组逐步回退，与左边的每一个墙做结算
时间复杂度: O(n)
空间复杂度：O(n) 栈的空间
横着看：（一行一行计算）
    = 0 0 0 =         = 代表柱子
= 0 = = 0 = =         0 代表雨水
- - - - - - -         - 代表地面
0 1 2 3 4 5 6         代表坐标
*/
/*
1.获取 0 位置的柱子高度,栈为空, 横坐标加入栈中,
  stack: 0 (此处是横坐标,不是高度)
2.获取 1 位置的柱子高度, 不大于栈顶元素的高度, 横坐标加入栈中
  stack: 0 , 1
3.获取 2 位置的柱子高度, 比栈顶元素的高度大, 说明中间可能有水, 开始计算(当前元素未入栈)
  将栈顶元素 1 出栈. h=0 ( h 可以理解成池塘底的高度 ), 出栈后, 栈中剩余 stack: 0 , 此时的栈顶元素就是 0 , 计算和此时栈顶元素(也就是墙)的积水量 (2-0-1)*(min(2,1)-0)=1, sum+=1 sum=1 (这里有点绕, 但是很精妙).
  此时栈已空, 循环结束, 将 2 入栈
  stack: 2 ,
3.获取 3 位置的柱子高度, 不大于栈顶元素的高度, 横坐标加入栈中
  stack: 2 , 3
4.获取 4 位置的柱子高度, 不大于栈顶元素的高度, 横坐标加入栈中
  stack: 2 , 3 , 4
5.获取 5 位置的柱子高度, 比栈顶元素的高度大, 说明中间可能有水, 开始计算(当前元素未入栈)
  将栈顶元素 4 出栈.h=0 出栈后, 栈中剩余 stack: 2 , 3 ,此时的栈顶元素就是 3 , 计算和此时栈顶元素(也就是墙)的积水量 (5-3-1)*(min(1,1)-0)=1, sum+=1 sum=2 ,
  stack: 2 , 3 , 5
6.获取 6 位置的柱子高度, 比栈顶元素的高度大, 说明中间可能有水, 开始计算(当前元素未入栈)
  将栈顶元素 5 出栈.h=1, 出栈后, 栈中剩余 stack: 2 , 3 ,此时的栈顶元素就是 3 , 计算和此时栈顶元素(也就是墙)的积水量 (6-3-1)*(min(1,2)-1)=0, sum+=0 sum=2 ,
  栈顶元素 3 的高度比当前柱子的高度低,故继续循环. 将栈顶元素 3 出栈.h=1, 出栈后, 栈中剩余 stack: 2 ,此时的栈顶元素就是 2 , 计算和此时栈顶元素(也就是墙)的积水量 (6-2-1)*(min(2,2)-1)=3, sum+=3 sum=5 ,
  当前位置 5 的柱子高度 不比栈顶元素的高度高, 循环结束, 将 6 入栈
  stack: 2 , 6
7.height 数组遍历完成, 循环结束*/
func trap3(height []int) int {
	sum := 0
	stack := make([]int, 0)
	current := 0

	for current < len(height) {
		for len(stack) != 0 && height[current] > height[stack[len(stack)-1]] {
			h := height[stack[len(stack)-1]]
			stack = stack[0 : len(stack)-1]  //取出栈顶元素（去除最后一个元素）（切片截取：不包含右边）
			if len(stack) == 0 {
				break
			}
			distance := current - stack[len(stack)-1] - 1
			min := Min(height[stack[len(stack)-1]], height[current])
			sum += distance * (min - h)
		}

		stack = append(stack, current)
		current++

	}
	return sum
}


//4.双指针法
/*
时间复杂度： O(n)
空间复杂度： O(1)
解法思想： 两墙夹一水，在找到右墙后，左墙数组逐步回退，与左边的每一个墙做结算(一列一列计算)
        =
    = 0 = 0 0 =         = 代表柱子
= 0 = = = 0 = =         0 代表雨水
- - - - - - - -       - 代表地面
0 1 2 3 4 5 6 7       代表坐标
*/
func trap4(height []int) int {
	var left, right, leftMax, rightMax, res int
	right = len(height) - 1
	for left < right {
		if height[left] < height[right] {  //左边柱子低于右边柱子
				if height[left] >= leftMax {
					//设置左边最高柱子
					leftMax = height[left]
				} else {
					//右边必定有柱子挡水，所以，遇到所有值小于等于leftMax的，全部加入水池
					res += leftMax - height[left]
				}
				left++
		} else {
				if height[right] > rightMax {
					//设置右边最高柱子
					rightMax = height[right]
				} else {
					//左边必定有柱子挡水，所以，遇到所有值小于等于rightMax的，全部加入水池
					res += rightMax - height[right]
				}
				right--
		}
	}
	return res
}


func Max(a,b int) int {
	if a>b{
		return a
	}else{
		return b
	}
}

func Min(x,y int) int {
	if x < y {
		return x
	}
	return y
}

