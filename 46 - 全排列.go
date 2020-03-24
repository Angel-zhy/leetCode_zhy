package main

import (
   "fmt"
   "sort"
)

func permute(nums []int) [][]int {
    size := len(nums)
    if size == 0 {
       return nil
    }
    //排序
    sort.Ints(nums)

    ans := make([]int,0)  
    res := make([][]int,0)
    used := make([]bool,size)
    helper(nums,used,ans,&res)

    return res
}

func helper(nums []int,used []bool,ans []int, res *[][]int){
   if len(ans) == len(nums) {
      *res = append(*res,append([]int{}, ans...))
      return
   }

   for i:= 0; i<len(nums); i++ {
      if !used[i] {
          used[i] = true
          ans = append(ans,nums[i])

          helper(nums,used,ans,res)

          ans = ans[:len(ans)-1]
          used[i] = false
      }
   }
}

func main()  {  
   // 给定一个没有重复数字的序列，返回其所有可能的全排列	
   s := []int{1,2,3}
   result := permute(s)
   fmt.Println(result)
}