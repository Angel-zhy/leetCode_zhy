package main

import (
	"fmt"
	"math"
)

//方法一：暴力法
//时间复杂度：O(n^2)，计算所有 n(n-1)/2种高度组合的面积。
//空间复杂度：O(1)，使用恒定的额外空间。
func maxArea(height []int) int {
    var max float64= 0.0
    for i:=0; i<len(height); i++ {
    	for j:= i+1; j<len(height); j++ {
    		max = math.Max(max,math.Min(float64(height[i]),float64(height[j]))*float64((j-i)))
		}
	}
	return int(max)
}

//方法二：双指针法
//这种方法背后的思路在于，两线段之间形成的区域总是会受到其中较短那条长度的限制。
//此外，两线段距离越远，得到的面积就越大。
//我们在由线段长度构成的数组中使用两个指针，一个放在开始，一个置于末尾。

//时间复杂度：O(n)O(n)，一次扫描。
//空间复杂度：O(1)O(1)，使用恒定的空间。
func maxArea2(height []int) int {
	var max float64 = 0.0
	l, r := 0, len(height)-1
	for l<r {
		max = math.Max(max,math.Min(float64(height[l]),float64(height[r]))*float64((r-l)))
        if height[l] < height[r] {
        	l++
		}else{
			r--
		}
	}
    return int(max)
}


func main(){
    //垂直的两条线段之间的最大化矩形区域面积
    slice :=  []int{1,9,6,2,5,4,8,15,9}

    //1.暴力法
    fmt.Println(maxArea(slice))

    //2.双指针法
    fmt.Println(maxArea2(slice))
}