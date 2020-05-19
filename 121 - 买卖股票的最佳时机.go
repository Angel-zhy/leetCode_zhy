package main

import "fmt"

func main() {
	s := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(s))
}

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	ans := 0
	min := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i]-min > ans {
			ans = prices[i] - min
		}
		if min > prices[i] {
			min = prices[i]
		}
	}

	return ans
}
