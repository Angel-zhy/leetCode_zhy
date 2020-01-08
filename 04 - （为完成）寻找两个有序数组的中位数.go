package main

import "fmt"

//时间复制度不符合要求
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var midIndex int
	pos1, pos2 := 0, 0
	len1, len2 := len(nums1), len(nums2)
	mergeSortedList := make([]int, 0)
	for ; pos1 < len1 && pos2 < len2; {
		if nums1[pos1] <= nums2[pos2] {
			mergeSortedList = append(mergeSortedList, nums1[pos1])
			pos1 += 1
		} else {
			mergeSortedList = append(mergeSortedList, nums2[pos2])
			pos2 += 1
		}
	}
	if pos1 == len1 {
		mergeSortedList = append(mergeSortedList, nums2[pos2:]...)
	} else {
		mergeSortedList = append(mergeSortedList, nums1[pos1:]...)
	}
	if (len1 + len2) % 2 == 0 {
		midIndex = (len1 + len2) / 2
		return float64(mergeSortedList[midIndex-1] + mergeSortedList[midIndex]) / 2
	} else {
		midIndex = (len1 + len2) / 2
		return float64(mergeSortedList[midIndex])
	}
}

func main()  {
	//请你找出这两个有序数组的中位数，
	//要求: 算法的时间复杂度为 O(log(m + n))
	nums1 := []int{1, 3}
	nums2 := []int{2, 5, 6}
	fmt.Println(findMedianSortedArrays(nums1,nums2))
}