package algo

func canJump(nums []int) bool {
	dp := make([]bool, len(nums))
	dp[len(nums)-1] = true

	for i := len(nums) - 2; i >= 0; i-- {
		for j := nums[i]; j > 0; j-- {
			if i+j >= len(nums)-1 || dp[i+j] {
				dp[i] = true
				break
			}
		}
	}

	return dp[0]
}
