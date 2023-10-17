package algo

import "fmt"

// Input:
// A = [1,2,3,0,0,0], m = 3
// B = [2,5,6],       n = 3

// Output: [1,2,2,3,5,6]

func merge(A []int, m int, B []int, n int) {
	i := m - 1
	j := n - 1
	// cur := m+n-1
	for cur := m + n - 1; cur >= 0; cur-- {
		if i < 0 && j < 0 {
			break
		} else if i < 0 {
			A[cur] = B[j]
			j--
		} else if j < 0 {
			A[cur] = A[i]
			i--
		} else {
			if A[i] > B[j] {
				A[cur] = A[i]
				i--
			} else {
				A[cur] = B[j]
				j--
			}
		}
	}
	fmt.Println(A)
}
