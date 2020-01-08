package main

import (
	"fmt"
	"reflect"
	"sort"
)

//给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？
//找出所有满足条件且不重复的三元组。

//方法一：复杂度到，超级繁琐（不建议使用）
func threeSum(nums []int) [][]int {
	var res [][]int
	var tem []int
    sort.Ints(nums)  //排序
	for  i := 0; i < len(nums) - 2; i++ { // 每个人
		for j := i + 1; j < len(nums) - 1; j++ { // 依次拉上其他每个人
			for k := j + 1; k < len(nums); k++ { // 去问剩下的每个人
				if nums[i] + nums[j] + nums[k] == 0 { // 我们是不是可以一起组队
				       tem = append(tem,nums[i],nums[j],nums[k])
				       if len(res) == 0 {
						   res = append(res,tem)
					   }
				       //检测是否重复
				       flag := 0
				       for f:=0; f<len(res);f++{
				       	   if reflect.DeepEqual(res[f],tem){
							   flag = 1
							   break
						   }
					   }
					   if flag == 0 {
						   res = append(res,tem)
					   }
					   tem = nil
				 }
	       }
	    }
	}
	return res
}

//方法二：推荐
func threeSum2(nums []int) [][]int {
	sort.Ints(nums)  //排序
	var result [][]int
	for i := 0; i < len(nums)-1 && nums[i] <= 0; i++ {  //nums[i]如果大于零，那么三数之和一定大于零
		if i > 0 && nums[i] == nums[i-1] { //相等必会重复，所以忽略一个，去重
			continue
		}
		j := i + 1
		k := len(nums) - 1
		target := 0 - nums[i]  //变为求两数之和target
		for ; j < k; {
			sum := nums[j] + nums[k]
			if sum == target {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				j++
				for j < k && nums[j-1] == nums[j] {
					j++
				}
				k--
				for j < k && nums[k+1] == nums[k] {
					k--
				}
			} else if sum > target {
				k--
				for j < k && nums[k+1] == nums[k] {
					k--
				}
			} else {
				j++
				for j < k && nums[j-1] == nums[j] {
					j++
				}
			}
		}
	}
	return result
}

func main(){
	nums := []int{-1, 0, 1, 2, -1, -4, 5}
	fmt.Println(threeSum(nums))
	fmt.Println(threeSum2(nums))
}