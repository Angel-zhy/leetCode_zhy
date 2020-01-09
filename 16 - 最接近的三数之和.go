package main

import (
	"fmt"
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) (sum,a,b,c int) {
	sort.Ints(nums)
    result := nums[0] + nums[1] + nums[2]
    for i:=0; i<len(nums)-2; i++ {
		l, r := i+1, len(nums)-1
		for l < r {
			sums := nums[i] + nums[l] + nums[r]
			if math.Abs(float64(sums-target))  <  math.Abs(float64(result-target)) {   //差的绝对值比较
				result = sums
				a,b,c = nums[i] , nums[l] , nums[r]  //比例每次之和是否比上一次接近，如果是，更新记录这三个值
			}
			if sums > target {
				r--
			}else if sums < target {
				l++
			}else {
				return sums,nums[i],nums[l],nums[r]
			}
		}
	}
	return result,a,b,c
}

func main(){
	//双指针法
	//数组排序，然后从首位下标 i:=0开始遍历，左指针i+1 右指正len(nums)-1;
	//当 三个数和== target时，返回target，如果大于和，则右指针-- 否则左指针++
	 nums := []int{-4,-2,-1,1,5}
	 target := 1
     fmt.Println(threeSumClosest(nums,target))
}




