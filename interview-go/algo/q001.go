package algo

// 实现斐波那契数列

// 使用暴力递归方法
func fib1(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib1(n-1) + fib1(n-2)
}

// 使用备忘录的方法记住重复问题
func fib2(n int) int {
	memo := make([]int, n+1)
	return dptable(memo, n)
}

func dptable(memo []int, n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	if memo[n] != 0 {
		return memo[n]
	}

	memo[n] = dptable(memo, n-1) + dptable(memo, n-2)
	return memo[n]
}

// 使用dp数组的迭代方法解
func fib3(n int) int {
	if n == 0 {
		return 0
	}

	dp := make([]int, n+1)
	// base case
	dp[0], dp[1] = 0, 1

	// 状态转移
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}
