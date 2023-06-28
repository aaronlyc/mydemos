package algo

import (
	"math"
)

// 凑零钱问题

// 1. 使用暴力递归
func coinChange1(coins []int, amount int) int {
	// base case
	if amount == 0 {
		return 0
	}

	if amount < 0 {
		return -1
	}

	res := math.MaxInt32
	for _, coin := range coins {
		// 计算子问题
		subProblem := coinChange1(coins, amount-coin)
		// 子问题无解则跳过
		if subProblem == -1 {
			continue
		}
		// 子问题中选择最优的解
		res = min(res, subProblem+1)
	}
	// 如果没有找到最小值，则返回失败，-1代替
	if res == math.MaxInt32 {
		return -1
	}
	return res
}

// 2. 使用带备忘的方法
func coinChange2(coins []int, amount int) int {
	memo := make([]int, amount+1)
	for i := range memo {
		memo[i] = -666
	}

	var dp func([]int, int) int
	dp = func(coins []int, amount int) int {
		if amount == 0 {
			return 0
		}
		if amount < 0 {
			return -1
		}

		// 如果子问题已经计算过了，则直接返回子问题的答案
		if memo[amount] != -666 {
			return memo[amount]
		}

		// res 保存本次子问题的答案
		res := math.MaxInt32
		for _, coin := range coins {
			subProblem := dp(coins, amount-coin)
			if subProblem == -1 {
				continue
			}
			res = min(res, subProblem+1)
		}
		if res == math.MaxInt32 {
			memo[amount] = -1
		} else {
			memo[amount] = res
		}
		return memo[amount]
	}

	return dp(coins, amount)
}

// 3. 使用迭代的方式实现
func coinChange3(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = amount + 1
	}

	// base case
	dp[0] = 0
	// 外层 for 循环在遍历所有状态的所有取值
	for i := 0; i < len(dp); i++ {
		// 内层 for 循环在求所有选择的最小值
		for _, coin := range coins {
			// 子问题无解，跳过
			if i-coin < 0 {
				continue
			}
			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}

	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
