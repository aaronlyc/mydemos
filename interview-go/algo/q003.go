package algo

// 最长递增子序列

// 1. 使用dp方法
func lengthOfLIS(nums []int) int {
	l := len(nums)

	if l == 0 {
		return 0
	}

	dp := make([]int, l)

	res := 1
	for i := 0; i < l; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
