package main

func main() {
	/*
		0 1 1 1 0
		1 1 1 1 0
		0 1 1 1 1
		0 1 1 1 1
		0 0 1 1 1
	*/
	//方法： 动态规划
	//注意：
	//    第一个循环dp(i,j) 的值：表示以坐标(i,j)是否为一个小方块，
	//    第二个循环dp(i,j) 的值：表示以坐标(i,j)为右下角的最大正方形的边长，

	/*时间复杂度：O(mn)O(mn)，其中 mm 和 nn 是矩阵的行数和列数。
	需要遍历原始矩阵中的每个元素计算 dp 的值。
	空间复杂度：O(mn)O(mn)，其中 mm 和 nn 是矩阵的行数和列数。
	创建了一个和原始矩阵大小相同的矩阵 dp。
	由于状态转移方程中的 dp(i, j) 由其上方、左方和左上方的三个相邻位置的 dp 值决定，
	因此可以使用两个一维数组进行状态转移，空间复杂度优化至 O(n)。
	*/
	var slice = make([][]byte, 5)
	slice[0] = []byte{'0', '1', '1', '1', '0'}
	slice[1] = []byte{'1', '1', '1', '1', '0'}
	slice[2] = []byte{'0', '1', '1', '1', '1'}
	slice[3] = []byte{'0', '1', '1', '1', '1'}
	slice[4] = []byte{'0', '0', '1', '1', '1'}
	maximalSquare(slice)
}

func maximalSquare(matrix [][]byte) int {
	dp := make([][]int, len(matrix))
	maxSide := 0
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[i]))
		for j := 0; j < len(matrix[i]); j++ {
			dp[i][j] = int(matrix[i][j] - '0')
			if dp[i][j] == 1 {
				maxSide = 1
			}
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if dp[i][j] == 1 {
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
				if dp[i][j] > maxSide {
					maxSide = dp[i][j]
				}
			}
		}
	}

	return maxSide * maxSide

}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
