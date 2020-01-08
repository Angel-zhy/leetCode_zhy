package main

import "fmt"

//把得到的结果放入切片中返回
func twoSum(s []int,target int)[]int{
	//定义存放结果的切片
	ret := make([]int,2)
	//定义Map存放已经遍历过的值和其坐标
    vkMap := make(map[int]int)
	for k,v := range s {
		 //差值(查看是否存在vkMap中)
		 addend := target - v
		 if _,ok := vkMap[addend]; ok {  //此时找到的v为第二个值
		 	//如果存在，证明找到两个数了
		 	ret[0] = vkMap[addend]
		 	ret[1] = k
		 }
		 //如果不符合，则把这一次的值存入已遍历的Map中
		 vkMap[v] = k
	}
	return ret
}

func main(){
	//给定 nums = [2, 7, 11, 15], target = 9
	//不能重复利用这个数组中同样的元素。
   nums := []int{2, 7, 11, 15}
   retSlice := twoSum(nums,9)
   fmt.Println(retSlice)
}
