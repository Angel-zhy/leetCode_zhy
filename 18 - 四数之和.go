package main

import (
	"fmt"
	"sort"
)

func fourSum3(nums []int, target int) [][]int {
	l := len(nums)
	res := make([][]int, 0)
	if l<4 {return res}
	// a,b,c,d四指针
	sort.Ints(nums)
	var sum int
	for a:=0; a<l-3; a++ {	// 最左指针
		if a>0 && nums[a]==nums[a-1] {  // 剔除重复项
			continue
		}
		for b:=a+1; b<l-2; b++ {	// 次左指针
			if b>a+1 && nums[b]==nums[b-1] {  // 剔除重复项
				continue
			}
			// 内层双指针
			for c, d := b+1, l-1; c<d; {
				sum = nums[a] + nums[b] + nums[c] + nums[d]
				switch {
				case sum < target:
					c++
					for c<d && nums[c]==nums[c-1] {c++} 	// 剔除重复项
				case sum > target:
					d--
					for c<d && nums[d]==nums[d+1] {d--}		// 剔除重复项
				default:	// sum = target
					res = append(res, []int{nums[a], nums[b], nums[c], nums[d]})	// 找到可行解且不重复
					d--
					for c<d && nums[d]==nums[d+1] {d--}		// 剔除重复项
					c++
					for c<d && nums[c]==nums[c-1] {c++} 	// 剔除重复项
				}
			}
		}
	}

	return res
}

func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return nil
	}
	sort.Ints(nums) //排序
	var res [][]int
	var temp []int
	sLen := len(nums)
	for i:=0; i<sLen-3; i++ {
		// 不存在 (很重要，有可能减少很多步操作)
		if  target<=0 && nums[i]>0 {
			break
		}
		// 此时条件不满足，后面的也不用再遍历了
		if nums[i]+nums[i+1]+nums[i+2]+nums[i+3]>target{
			break
		}
		// 此时i对应的固定值 + 最大的三个值都小于目标值，无需遍历；跳过num[i](增大固定值)，再来判断
		if nums[i]+nums[sLen-3]+nums[sLen-2]+nums[sLen-1]<target {
			continue   //只能跳过，不能使用break,因为num[i]的值在下一次循环中可以变大（重点！！！)
		}
		//  重复项 (有结果的组合，在前一个i是已经得出结果)
		if i>0 && nums[i] == nums[i-1] {
			continue
		}

		for j:= i+1; j< sLen-2; j++ {   //第2个固定值循环倒数第3个数集客，后面还有2个指针
			//不存在 (此时j所对应的、两个指针指向最小值时，相加已经大于目标值)
			if nums[i] + nums[j] + nums[j+1] + nums[j+2] > target {
				break
			}
			//   不存在 ，两指针指向最大值
			if nums[i] + nums[j] + nums[sLen-2] + nums[sLen-1] < target {
				continue   //只能跳过，不能使用break（重点！！！)
			}
			//重复项
			if j>i+1 && nums[j] == nums[j-1]{
				continue
			}

			start := j+1 //左指针
			end := sLen-1 //右指针
			for start < end {
				sum := nums[i] + nums[j] + nums[start] + nums[end] //四数之和
				//fmt.Println(sum)
				if sum < target {
					start++
				}else if sum > target {
					end--
				}else{
					temp = append(temp, nums[i],nums[j],nums[start],nums[end])
					res = append(res,temp)
					temp = nil
					//去重
					for start<end && nums[start]==nums[start+1] {start++}		// 剔除重复项
					for start<end && nums[end]==nums[end-1] {end--}
					start++
				}
			}
		}
	}
	return res
}

func main(){
	//给定一个包含 n 个整数的数组 nums 和一个目标值 target，
	//判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？
	//找出所有满足条件且不重复的四元组。
	//解题思路:
	//基本思想就和 三数之和 相同

	//先后固定其中两个数值 ii、jj，再使用双指针寻找与目标合适的差值
	//算法流程：
	//1.判断数组大小是否符合要求
	//2.对数组进行排序 --> 此处是为了方便双指针根据数值大小移动
	//3.固定一个数 nums[i]，并判断条件是否成立 或 是否有重复项
	//4.在固定 nums[i] 的基础上，再次固定一个数值 nums[j]，同样判断条件是否成立 或 是否有重复项
	//5.在 nums[j] 后面两端设定两个指针 start=j+1, end= len(nums)−1, 计算 nums[i]+nums[j]+nums[start]+nums[end]，判断四数之和 SumSum 与目标值的大小关系:

	//如果 Sum < target, end指针左移
	//如果 Sum > target, start指针右移
	//如果 Sum == target, 将四数添加进结果数组中
	//nums[start] == nums[start - 1]--> 结果重复，跳过这一结果
	//nums[end] == nums[end - 1]--> 结果重复，跳过这一结果

	//复杂度分析:
	//时间复杂度：O(n^3)
	//数组排序： O(nlog(n))
	//数组遍历两次： O(n^2)
	//双指针：O(n)
	//总体：O(nlog(n)) + O(n^2)*O(n))
	//空间复杂度：O(1)


     //slice := []int{1, 0, -1, 0, -2, 1, 2}
     //slice := []int{-5,-2,1,1,3,5,5,5}  //4
     //slice := []int{-1,0,1,2,-1,-4}  //-1
     slice := []int{0,4,-5,2,-2,4,2,-1,4}  //-1
     target := 12
     fmt.Println(fourSum(slice,target))

}
