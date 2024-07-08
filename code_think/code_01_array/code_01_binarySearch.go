package main

// 时间复杂度 O(logn)
func search1(nums []int, target int) int {
	// 初始化左右边界
	left := 0
	right := len(nums) - 1

	// 循环逐步缩小区间范围
	for left <= right {
		// 求区间中点
		mid := left + (right-left)>>1

		// 根据 nums[mid] 和 target 的大小关系
		// 调整区间范围
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	// 在输入数组内没有找到值等于 target 的元素
	return -1
}

// 时间复杂度 O(logn)
func search2(nums []int, target int) int {
	// 初始化左右边界
	left := 0
	right := len(nums)

	// 循环逐步缩小区间范围
	for left < right {
		// 求区间中点
		mid := left + (right-left)>>1

		// 根据 nums[mid] 和 target 的大小关系
		// 调整区间范围
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	// 在输入数组内没有找到值等于 target 的元素
	return -1
}

func main() {
	println(search1([]int{1, 3, 5, 6}, 5))
	println(search2([]int{1, 3, 5, 6}, 2))

}
