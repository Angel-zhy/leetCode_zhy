package main

import "fmt"

//方法：双指针法
//时间复杂度：O(n)，假设数组的长度是 n，那么 i 和 j 分别最多遍历 n 步。
//空间复杂度：O(1)。
func removeDuplicates(nums []int) int {
     if len(nums) == 0 {
     	return 0
	 }
	//数组完成排序后，我们可以放置两个指针 i 和 j，其中 i 是慢指针，而 j 是快指针。
	//只要 nums[i] = nums[j]，我们就增加 j 以跳过重复项;

	//比较 i 和 j 位置的元素是否相等。
	//如果相等，j 后移 1 位
	//如果不相等，将 j 位置的元素复制到 i+1 位置上，j 后移一位，i 后移 1 位
	//重复上述过程，直到 j 等于数组长度。
	//返回 i + 1，即为新数组长度。

	 var i int
     for j:=1; j<len(nums); j++ {
     	if nums[j] != nums[i] {
     		i++
		   nums[i] = nums[j]
		}
	 }
	 //fmt.Println(nums)  //[1 2 3 6 8 9 10 9 10 10]
	 return i+1
}

func main()  {
	//要求：
	//给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。
	//不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成

	 slice := []int{1,2,2,3,6,8,8,9,10,10}
	 fmt.Println(removeDuplicates(slice))

}