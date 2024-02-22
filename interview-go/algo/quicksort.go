package algo

func QuickSort(nums []int) {
	quicksort(nums, 0, len(nums)-1)
}

func quicksort(nums []int, left, right int) {
	if left >= right {
		return
	}

	pivot := partition(nums, left, right)
	quicksort(nums, left, pivot-1)
	quicksort(nums, pivot+1, right)
}

func partition(nums []int, left, right int) int {
	//	以nums[left]作为基准数
	i, j := left, right
	for i < j {
		for i < j && nums[j] >= nums[left] {
			j--
		}
		for i < j && nums[i] <= nums[left] {
			i++
		}
		//	元素交换
		nums[i], nums[j] = nums[j], nums[i]
	}

	//	将基准数交换至两子数组的分界线
	nums[i], nums[left] = nums[left], nums[i]
	return i
}
